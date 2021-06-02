package main

import (
	"fmt"
	"tdapi/config"
	"tdapi/container"
	"tdapi/container/logger"
	"tdapi/container/servicecontainer"
	"tdapi/usecase"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var router *gin.Engine

var db *gorm.DB

const (
	DEV_CONFIG  string = "../../config/appConfigDev.yaml"
	PROD_CONFIG string = "../../config/appConfigProd.yaml"
)

type UserService struct {
	container container.Container
}

func catchPanic() {
	if p := recover(); p != nil {
		logger.Log.Errorf("%+v\n", p)
	}
}

func runServer(sc *servicecontainer.ServiceContainer) error {
	logger.Log.Debug("start runserver")

	gin.SetMode(gin.DebugMode)

	router = gin.Default()

	initializeRoutes()

	//l, err:=net.Listen(GRPC_NETWORK, GRPC_ADDRESS)
	ugc := sc.AppConfig.UserGrpcConfig
	logger.Log.Debugf("userGrpcConfig: %+v\n", ugc)

	LoadTdList(sc)

	return router.Run()

}

func main() {
	filename := DEV_CONFIG
	//filename := PROD_CONFIG
	container, err := buildContainer(filename)
	if err != nil {
		fmt.Printf("%+v\n", err)
		//logger.Log.Errorf("%+v\n", err)
		panic(err)
	}
	if err := runServer(container); err != nil {
		logger.Log.Errorf("Failed to run user server: %+v\n", err)
		panic(err)
	} else {
		logger.Log.Info("server started")
	}
}

func buildContainer(filename string) (*servicecontainer.ServiceContainer, error) {

	factoryMap := make(map[string]interface{})
	config := config.AppConfig{}
	container := servicecontainer.ServiceContainer{factoryMap, &config}

	err := container.InitApp(filename)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &container, nil
}

func getClientManagerUseCase(c container.Container) (usecase.ClientManagerUseCaseInterface, error) {
	key := config.CLIENTMANAGER
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.ClientManagerUseCaseInterface), nil

}

//加载所有列表
func LoadTdList(container container.Container) {
	ruci, err := getClientManagerUseCase(container)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}

	ruci.LoadUserList()

	//user := model.User{Name: "Brian", Department: "Marketing", Created: created}

	// resultUser, err := ruci.RegisterUser(&user)
	// if err != nil {
	// 	logger.Log.Errorf("user registration failed:%+v\n", err)
	// } else {
	// 	logger.Log.Info("new user registered:", resultUser)
	// }
}
