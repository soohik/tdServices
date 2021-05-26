package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	phone "tdimpl/adpater/phoneservice"
	"tdimpl/config"
	"tdimpl/container"
	"tdimpl/container/servicecontainer"
	"tdimpl/usecase"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

const (
	DEV_CONFIG string = "../config/appConfigDev.yaml"
)

// UnimplementedPhoneServiceServer must be embedded to have forward compatible implementations.
var uni phone.UnimplementedPhoneServiceServer

type PhoneService struct {
	container container.Container
	phone.UnimplementedPhoneServiceServer
}

func (c *PhoneService) RegPhone(context.Context, *phone.PhoneRegRequest) (*phone.PhoneRegResponse, error) {

	a := &phone.PhoneRegResponse{Err: 12345, Msg: "hello"}
	return a, nil

}

func getRegistrationUseCase(c container.Container) (usecase.RegistrationUseCaseInterface, error) {
	key := config.REGISTRATION
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.RegistrationUseCaseInterface), nil

}

func runServer(sc *servicecontainer.ServiceContainer) error {
	logger.Println("start runserver")

	srv := grpc.NewServer()

	cs := &PhoneService{sc, uni}

	phone.RegisterPhoneServiceServer(srv, cs)

	ugc := sc.AppConfig.UserGrpcConfig
	logger.Println("userGrpcConfig: %+v\n", ugc.UrlAddress)
	l, err := net.Listen(ugc.DriverName, ugc.UrlAddress)
	if err != nil {
		return errors.Wrap(err, "")
	} else {
		logger.Println("server listening")
	}
	return srv.Serve(l)
}

func buildContainer(filename string) (*servicecontainer.ServiceContainer, error) {

	factoryMap := make(map[string]interface{})
	config := config.AppConfig{}
	container := servicecontainer.ServiceContainer{factoryMap, &config}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	err = container.InitApp(filename)
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
		logger.Println("Failed to run user server: %+v\n", err)
		panic(err)
	} else {
		logger.Println("server started")
	}
}
