package usecasefactory

import (
	"tdimpl/config"
	"tdimpl/container"
)

type RegistrationFactory struct {
}

// Build creates concrete type for RegistrationUseCaseInterface
func (rf *RegistrationFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	// uc := appConfig.UseCase.Registration
	// udi, err := buildUserData(c, &uc.UserDataConfig)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }
	// tdi, err := buildTxData(c, &uc.TxDataConfig)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "")
	// }
	// ruc := registration.RegistrationUseCase{UserDataInterface: udi, TxDataInterface: tdi}

	// return &ruc, nil
	return nil, nil
}
