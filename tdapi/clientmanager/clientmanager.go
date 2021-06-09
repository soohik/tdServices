// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"errors"
	"fmt"
	"tdapi/dataservice"
	"tdapi/log"
	"tdapi/model"
	"time"

	"github.com/Arman92/go-tdlib"
)

const (
	tddata = "../../tddata/"
	tdfile = "../../tdfile/"
)

type TdInstance struct {
	Phone  string
	ReReg  int //重试次数
	Client *TDClient
}

// ListUserUseCase implements ListUseCaseInterface.
type ClientManagerUseCase struct {
	TdInstances []TdInstance
}

var ClientManager ClientManagerUseCase

func (c *ClientManagerUseCase) LoadTdInstance() error {

	clients, err := dataservice.GetAllPhone()
	if err != nil {
		return nil
	}
	c.AddTdlibClient(clients)

	return nil
}

//
func (c *ClientManagerUseCase) RemoveClient(phone string) bool {

	var index = -1
	for index = range c.TdInstances {
		if c.TdInstances[index].Phone == phone {
			c.TdInstances[index].Client.tdlibClient.Close()
			c.TdInstances[index].Client.tdlibClient = nil
			c.TdInstances[index].Client = nil

			break
		}
	}
	if index >= 0 {
		c.TdInstances = append(c.TdInstances[:index], c.TdInstances[index+1:]...)
		return true
	}

	return false
}

//
func (c *ClientManagerUseCase) GetClient(phone string) (model.Phone, bool) {

	return dataservice.GetPhone(phone)

}

func RegisterPhone(phonenumber, logincode string) (bool, model.Client) {

	_, find := ClientManager.GetClient(phonenumber)
	if !find {
		return false, model.Client{}
	}
	client, ret := ClientManager.AddInstance(phonenumber, logincode)
	if ret == model.SOK {
		return true, client
	}
	return false, model.Client{}

}

func PreRegisterPhone(phonenumber string) (bool, int) {

	var phone model.Phone
	phone.Account = phonenumber
	phone.Phone = phonenumber
	phone.Registered = 0
	phone.Tddata = tddata + phonenumber + "-tdlib-db"

	phone.Tdfile = tdfile + phonenumber + "-tdlib-files"

	ok := dataservice.Preregister(phone)
	if ok {
		ClientManager.AddInstance(phonenumber, "")

	}

	return false, model.SOK

}

//增加单个客户端
func (c *ClientManagerUseCase) ReAddClient(phonenumber string) {

	find := c.RemoveClient(phonenumber)

	if find {

		var phone model.Phone
		phone.Account = phonenumber
		phone.Phone = phonenumber
		phone.Registered = 0
		phone.Tddata = tddata + phonenumber + "-tdlib-db"

		phone.Tdfile = tdfile + phonenumber + "-tdlib-files"

		if find { //插入成功
			ClientManager.AddClient(phone)
		}
	}

}

//增加单个客户端
func (c *ClientManagerUseCase) AddClient(m model.Phone) {

	client := TdInstance{}
	client.Client = new(TDClient)
	client.Phone = m.Phone

	client.Client.AcountName = m.Account
	client.Client.TdlibDbDirectory = m.Tddata
	client.Client.TdlibFilesDirectory = m.Tdfile

	go client.Client.AddInstance()
}
func (c *ClientManagerUseCase) AddTdlibClient(m []model.Phone) {

	for _, value := range m {
		client := TdInstance{}
		client.Client = new(TDClient)
		client.Phone = value.Phone

		client.Client.AcountName = value.Account
		client.Client.TdlibDbDirectory = value.Tddata
		client.Client.TdlibFilesDirectory = value.Tdfile
		go client.Client.AddInstance()
	}

}

func BuildClientManager() {
	// ClientManager.LoadTdInstance()

}

func Joinlink(account, linkurl, groupname string) (int, error) {

	dataservice.InsertGroup(groupname, linkurl)
	dataservice.InsertGroupsInfo(account, groupname)

	tdlib.SetLogVerbosityLevel(1)
	// Create new instance of client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "228834",
		APIHash:             "e4d4a67594f3ddadacab55ab48a6187a",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   tddata + account + "-tdlib-db",
		FileDirectory:       tdfile + account + "-tdlib-files",
		IgnoreFileNames:     false,
	})
	defer client.Close() //关闭

	currentState, _ := client.Authorize()
	for ; currentState.GetAuthorizationStateEnum() != tdlib.AuthorizationStateReadyType; currentState, _ = client.Authorize() {
		time.Sleep(300 * time.Millisecond)
	}

	char, err := client.SearchPublicChat(groupname)
	if err != nil {
		return model.BadRequest, errors.New("无效的组")
	}
	ok, _ := client.JoinChat(char.ID)
	fmt.Println(ok, err)

	if err != nil {
		return model.AuthSenCodeErr, errors.New("加入失败")
	}

	return model.SOK, nil
}

func (c *ClientManagerUseCase) AddInstance(account, code string) (model.Client, int) {
	var ret int
	var repclient model.Client

	tdlib.SetLogVerbosityLevel(1)
	// Create new instance of client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "228834",
		APIHash:             "e4d4a67594f3ddadacab55ab48a6187a",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   tddata + account + "-tdlib-db",
		FileDirectory:       tdfile + account + "-tdlib-files",
		IgnoreFileNames:     false,
	})
	defer client.Close() //关闭

	for {
		currentState, err := client.Authorize()
		if err != nil {
			log.Infof("Error getting current state: %s %v", account, err)
			break
		}
		fmt.Println(currentState)
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {

			_, err := client.SendPhoneNumber(account)
			if err != nil {
				log.Infof("phone %s  err: %v", account, err)
				ret = model.AuthSendTimeout
				break
			}
			ret = model.AuthWaitCode

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			if code != "" {
				_, err := client.SendAuthCode(code)
				if err != nil {
					fmt.Printf("Error sending auth code : %v", err)
				}
			} else {
				return repclient, model.AuthWaitCode
			}

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			log.Info("Authorization Ready! Let's rock", account)
			user, _ := client.GetMe()
			repclient.Id = user.ID
			repclient.Name = user.FirstName
			repclient.Username = user.Username
			repclient.PhoneNumber = user.PhoneNumber
			ret = model.SOK
			break

		}
	}

	return repclient, ret
}

func Getallgroups(agent string) ([]model.Groups, error) {
	return dataservice.GetAlGroups(agent)
}

func GetMegroups(agent string) ([]model.Groups, error) {
	return dataservice.GetAlGroups(agent)
}
