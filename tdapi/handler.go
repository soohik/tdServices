package main

import (
	"tdapi/adapter/phoneclient"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	phoneclient.JsonToPhone(c)

}

func preregister(c *gin.Context) {
	// phone, err := phoneclient.JsonToPhone(c)
	// if err != nil {
	// 	return
	// }

	//查找数据库
	// find := Preregister(phone.Phone)

	// fmt.Println(find, err)

}
