package userdataservicefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/dataservice"
)

var udsFbMap = map[string]userDataServiceFbInterface{
	config.SQLDB: &sqlUserDataServiceFactory{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type userDataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (dataservice.UserDataInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetUserDataServiceFb(key string) userDataServiceFbInterface {
	return udsFbMap[key]
}
