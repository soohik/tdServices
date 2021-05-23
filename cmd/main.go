package main

import (
	"fmt"
	"os"
	"tdServices/config"
	"tdServices/container"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DEV_CONFIG string = "../config/appConfigDev.yaml"
)

func main() {
	filename := DEV_CONFIG
	container, err := buildContainer(filename)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(container)

	if err := runServer(container); err != nil {
		fmt.Errorf("Failed to run user server: %+v\n", err)
		panic(err)
	} else {
		fmt.Errorf("server started")
	}

}

func buildContainer(filename string) (*container.ServiceContainer, error) {
	factoryMap := make(map[string]interface{})
	appConfig := config.AppConfig{}
	container := container.ServiceContainer{factoryMap, &appConfig}

	err := container.InitApp(filename)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return &container, nil
}

func runServer(sc *container.ServiceContainer) error {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(gin.Recovery())
	app.Use(gin.ErrorLogger())
	addr := httpListenString(sc)

	app.Run(addr)

	return nil
}
func httpListenString(sc *container.ServiceContainer) string {
	listen := fmt.Sprintf("%s:%d", "127.0.0.1", sc.AppConfig.Service.Port)
	return listen
}
