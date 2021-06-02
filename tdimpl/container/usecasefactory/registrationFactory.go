package usecasefactory

import (
	"tdimpl/config"
	"tdimpl/container"
	"tdimpl/container/dataservicefactory"
	"tdimpl/dataservice"
	"tdimpl/usecase/registration"

	"github.com/pkg/errors"
)

type RegistrationFactory struct {
}

func buildTxData(c container.Container, dc *config.DataConfig) (dataservice.TxDataInterface, error) {
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tdi := dsi.(dataservice.TxDataInterface)
	return tdi, nil
}

// Build creates concrete type for RegistrationUseCaseInterface
func (rf *RegistrationFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.Registration
	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tdi, err := buildTxData(c, &uc.TxDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi, TxDataInterface: tdi}

	return &ruc, nil

}
