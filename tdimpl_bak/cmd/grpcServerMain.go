package main

import (
	"fmt"
	"net"
	"tdimpl/config"
	"tdimpl/container/servicecontainer"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	DEV_CONFIG string = "../../config/appConfigDev.yaml"
)

type server struct{}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
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
