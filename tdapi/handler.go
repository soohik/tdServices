package main

import (
	"fmt"
	"tdapi/adapter/phoneclient"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	phoneclient.JsonToPhone(c)

}

func preregister(c *gin.Context) {
	phone, err := phoneclient.JsonToPhone(c)
	if err != nil {
		return
	}

	fmt.Println(phone)

}
