// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"tdapi/dataservice"
	"tdapi/model"
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

	ClientManager.GetClient(phonenumber)
	return true

}

func PreRegisterPhone(phonenumber string) bool {

	var phone model.Phone
	phone.Account = phonenumber
	phone.Phone = phonenumber
	phone.Registered = 0
	phone.Tddata = tddata + phonenumber + "-tdlib-db"

	phone.Tdfile = tdfile + phonenumber + "-tdlib-files"

	find := dataservice.Preregister(phone)
	if find { //插入成功
		ClientManager.AddClient(phone)
	}
	return true

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
	ClientManager.LoadTdInstance()

}
