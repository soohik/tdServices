package usecasefactory

import (
	"tdapi/config"
	"tdapi/container"
	"tdapi/usecase/listuser"

	"github.com/pkg/errors"
)

type ListUserFactory struct{}

func (luf *ListUserFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.ListUser

	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi, err := buildCacheData(c, &uc.CacheDataConfig)
	luuc := listuser.ListUserUseCase{UserDataInterface: udi, CacheDataInterface: cdi}
	return &luuc, nil
}
