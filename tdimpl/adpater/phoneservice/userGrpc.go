// Package userclient is client library if you need to call the user Micro-service as a client.
// It provides client library and the data transformation service.
package phoneservice

import (
	"tdimpl/model"
)

// GrpcToUser converts from grpc User type to domain Model user type
func GrpcToUser(user *Phone) (*model.Phone, error) {
	if user == nil {
		return nil, nil
	}
	resultphone := model.Phone{}

	return &resultphone, nil
}
