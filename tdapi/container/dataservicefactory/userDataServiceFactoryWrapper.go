package dataservicefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/container/dataservicefactory/userdataservicefactory"
	"tdapi/container/logger"

	"github.com/pkg/errors"
)

// userDataServiceFactory is a empty receiver for Build method
type userDataServiceFactoryWrapper struct{}

func (udsfw *userDataServiceFactoryWrapper) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("UserDataServiceFactory")
	key := dataConfig.DataStoreConfig.Code
	udsi, err := userdataservicefactory.GetUserDataServiceFb(key).Build(c, dataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return udsi, nil
}

