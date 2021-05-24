package container

import (
	"tdServices/config"

	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	AppConfig  *config.AppConfig
}

func (sc *ServiceContainer) InitApp(filename string) error {
	var err error
	config, err := loadConfig(filename)
	if err != nil {
		return errors.Wrap(err, "loadConfig")
	}
	sc.AppConfig = config

	if err != nil {
		return errors.Wrap(err, "loadLogger")
	}

	return nil
}

// loads the application configurations
func loadConfig(filename string) (*config.AppConfig, error) {

	ac, err := config.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read container")
	}
	return ac, nil
}

func (sc *ServiceContainer) BuildUseCase(code string) (interface{}, error) {
	// return usecasefactory.GetUseCaseFb(code).Build(sc, sc.AppConfig, code)
	return nil, nil
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
}