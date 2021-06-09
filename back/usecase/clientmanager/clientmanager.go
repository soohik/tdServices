// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"tdapi/dataservice"
	"tdapi/model"
)

type TdInstance struct {
	Phone  string
	Client *TDClient
}

// ListUserUseCase implements ListUseCaseInterface.
type ClientManagerUseCase struct {
	TdInstances       []TdInstance
	UserDataInterface dataservice.UserDataInterface
	TxDataInterface   dataservice.TxDataInterface
}

func (c *ClientManagerUseCase) LoadTdInstance() error {

	return nil
}

func (c *ClientManagerUseCase) LoadUserList() error {
	phones, err := GetAllClient(c.UserDataInterface)
	if err != nil {
		return nil
	}
	c.AddTdlibClient(phones)

	return nil
}

func GetAllClient(udi dataservice.UserDataInterface) ([]model.Phone, error) {
	return udi.GetAllClient()

}

func PreRegisterPhone(phone string) error {

	return nil
}

func (c *ClientManagerUseCase) AddTdlibClient(m []model.Phone) {
	client := TdInstance{}
	client.Client = new(TDClient)

	for _, value := range m {
		client.Phone = value.Phone

		client.Client.AcountName = value.Account
		client.Client.TdlibDbDirectory = value.Tddata
		client.Client.TdlibFilesDirectory = value.Tdfile

		go client.Client.AddInstance()
	}

}

func (c *ClientManagerUseCase) PreRegisterPhone(phone string) error {
	phones, err := GetAllClient(c.UserDataInterface)
	if err != nil {
		return nil
	}
	c.AddTdlibClient(phones)

	return nil
}
