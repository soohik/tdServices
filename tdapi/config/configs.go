package config

var Config *AppConfig

func LoadConfigs(filename string) error {
	var err error

	Config, err = ReadConfig(filename)
	if err != nil {
		return err
	}
	return err

}
