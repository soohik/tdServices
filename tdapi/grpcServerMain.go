package main

import (
	"tdapi/clientmanager"
	"tdapi/config"
	"tdapi/dataservice"
	"tdapi/log"
	"tdapi/tasks"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

const (
	DEV_CONFIG  string = "./config/appConfigDev.yaml"
	PROD_CONFIG string = "../../config/appConfigProd.yaml"
)

func catchPanic() {
	if p := recover(); p != nil {

	} 
}

func runServer() error {

	gin.SetMode(gin.DebugMode)

	router = gin.Default()
	router.Use(log.LoggerToFile())

	initializeRoutes()

	clientmanager.BuildClientManager()
	go tasks.InitTasks()

	//l, err:=net.Listen(GRPC_NETWORK, GRPC_ADDRESS)

	// LoadTdList(sc)
	log.Info("run server!")

	return router.Run()

}

func main() {
	filename := DEV_CONFIG

	config.LoadConfigs(filename)
	dataservice.SqlBuild()
	if err := runServer(); err != nil {

		panic(err)
	} else {

	}
}
