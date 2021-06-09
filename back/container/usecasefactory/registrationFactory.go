package usecasefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/usecase/clientmanager"
	"tdapi/usecase/registration"

	"github.com/pkg/errors"
)

type RegistrationFactory struct {
}

// Build creates concrete type for RegistrationUseCaseInterface
func (rf *RegistrationFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.Registration
	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	udi.InitUseCase(appConfig)

	tdi, err := buildTxData(c, &uc.TxDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi, TxDataInterface: tdi}

	return &ruc, nil
}

type ClientManagerFactory struct {
}

// Build creates concrete type for RegistrationUseCaseInterface
func (rf *ClientManagerFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.ClientManager
	udi, err := buildUserData(c, &uc.UserDataConfig)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tdi, err := buildTxData(c, &uc.TxDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := clientmanager.ClientManagerUseCase{UserDataInterface: udi, TxDataInterface: tdi}
	return &ruc, nil
}
