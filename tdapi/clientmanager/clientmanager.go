// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"fmt"
	"tdapi/dataservice"
	"tdapi/model"

	"github.com/Arman92/go-tdlib"
)

const (
	SOK             = 200
	AuthWaitCode    = 201
	AuthSendTimeout = -1
	AuthSenCodeErr  = -2
	AuthorizationStateClosed
	RegisterFailed = 409
)

const (
	tddata = "./tddata/"
	tdfile = "./tdfile/"
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
func (c *ClientManagerUseCase) GetClient(phone string) *TdInstance {

	for index := range c.TdInstances {
		if c.TdInstances[index].Phone == phone {

			return &c.TdInstances[index]
		}
	}

	return nil

}

func RegisterPhone(phonenumber, logincode string) bool {

	m := ClientManager.GetClient(phonenumber)
	if m != nil {
		m.Client.Code <- logincode
	}

	return true

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

	return false, SOK

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

func (c *ClientManagerUseCase) AddInstance(account, code string) int {
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
		DatabaseDirectory:   tddata,
		FileDirectory:       tddata,
		IgnoreFileNames:     false,
	})

	for {
		currentState, err := client.Authorize()
		if err != nil {
			fmt.Printf("Error getting current state: %v", err)
			continue
		}
		fmt.Println(currentState)
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {

			_, err := client.SendPhoneNumber(account)
			if err != nil {
				return AuthSendTimeout
			}
			return AuthWaitCode

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {

			_, err := client.SendAuthCode(code)
			if err != nil {
				fmt.Printf("Error sending auth code : %v", err)
			}

		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("Authorization Ready! Let's rock")
			return AuthWaitCode

		}
	}
	return 0
}
