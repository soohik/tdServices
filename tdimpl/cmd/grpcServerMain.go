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

func catchPanic() {
	if p := recover(); p != nil {
		logger.Println("%+v\n", p)
	}
}

func (c *PhoneService) RegPhone(con context.Context, req *phone.PhoneRegRequest) (*phone.PhoneRegResponse, error) {

	defer catchPanic()

	_, err := getRegistrationUseCase(c.container)
	if err != nil {
		logger.Println("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	mu, err := phone.GrpcToUser(req.Phone)

	if err != nil {
		logger.Println("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.Println("mu:", mu)
	// resultUser, err := ruci.UnregisterUser("1")
	// if err != nil {
	// 	logger.Println("%+v\n", err)
	// 	return nil, errors.Wrap(err, "")
	// }
	// logger.Println("resultUser:", resultUser)
	// gu, err := userclient.UserToGrpc(resultUser)
	// if err != nil {
	// 	logger.Println("%+v\n", err)
	// 	return nil, errors.Wrap(err, "")
	// }

	// logger.Println("user registered: ",)

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
