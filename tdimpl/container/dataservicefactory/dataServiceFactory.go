// Package dataservicefactory using factory method pattern to create concrete type to provide persistence service
// The source of data can come from database ( for domain model "user") or from other service ( for domain model cache,
// which comes from a gRPC service).
// There is only one method Build() for the factory and all different types of data service following the same interface
// to build the data service.

package dataservicefactory

import (
	"tdimpl/config"

	"tdimpl/container"
)

// To map "data service code" to "data service interface builder"
// Each data service need a separate builder
// Concrete builder is in corresponding factory file. For example, "courseDataServiceFactory" is in
// "courseDataServiceFactory.go"
var dsFbMap = map[string]dataServiceFbInterface{
	// config.USER_DATA: &userDataServiceFactoryWrapper{},
}

// DataServiceInterface serves as a marker to indicate the return type for Build method
type DataServiceInterface interface{}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (DataServiceInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetDataServiceFb(key string) dataServiceFbInterface {
	return dsFbMap[key]
}
