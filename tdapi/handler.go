package main

import (
	"fmt"
	"tdapi/adapter/phoneclient"
	"tdapi/clientmanager"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	phone, err := phoneclient.JsonToPhone(c)
	if err != nil {
		return
	}

	//查找数据库
	find := clientmanager.RegisterPhone(phone.Phone, phone.Code)

	fmt.Println(find, err)
}

func preregister(c *gin.Context) {
	phone, err := phoneclient.JsonToPhone(c)
	if err != nil {
		return
	}

	//查找数据库
	find := clientmanager.PreRegisterPhone(phone.Phone)

	fmt.Println(find, err)

}
