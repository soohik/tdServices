package main

import (
	"fmt"
	"net"
	"tdimpl/config"
	"tdimpl/container/servicecontainer"

	"github.com/micro/go-micro/v2/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	DEV_CONFIG string = "../../config/appConfigDev.yaml"
)

type server struct{}

func runServer(sc *servicecontainer.ServiceContainer) error {
	logger.Log.Debug("start runserver")

	srv := grpc.NewServer()

	cs := &UserService{sc}
	uspb.RegisterUserServiceServer(srv, cs)
	//l, err:=net.Listen(GRPC_NETWORK, GRPC_ADDRESS)
	ugc := sc.AppConfig.UserGrpcConfig
	logger.Log.Debugf("userGrpcConfig: %+v\n", ugc)
	l, err := net.Listen(ugc.DriverName, ugc.UrlAddress)
	if err != nil {
		return errors.Wrap(err, "")
	} else {
		logger.Log.Debug("server listening")
	}
	return srv.Serve(l)
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
