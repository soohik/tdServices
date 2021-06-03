// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package clientmanager

import (
	"tdapi/dataservice"
	"tdapi/model"
)

var ClientManager ClientManagerUseCase

type TdInstance struct {
	Phone  string
	Client *TDClient
}

// ListUserUseCase implements ListUseCaseInterface.
type ClientManagerUseCase struct {
	TdInstances []TdInstance
}

func (c *ClientManagerUseCase) LoadTdInstance() error {

	clients, err := dataservice.GetAllPhone()
	if err != nil {
		return nil
	}
	c.AddTdlibClient(clients)

	return nil
}

func PreRegisterPhone(phone string) error {

	return nil
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

func (c *ClientManagerUseCase) PreRegisterPhone(phone string) error {

	return nil
}

func BuildClientManager() {
	ClientManager.LoadTdInstance()
}
