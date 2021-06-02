package main

import (
	"tdapi/adapter/phoneclient"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	phoneclient.JsonToPhone(c)
}

func preregister(c *gin.Context) {
	phoneclient.JsonToPhone(c)
}
