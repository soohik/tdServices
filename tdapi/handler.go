package main

import (
	"encoding/json"
	"net/http"
	"tdapi/adapter/phoneclient"
	"tdapi/clientmanager"
	"tdapi/model"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	phone, err := phoneclient.JsonToPhone(c)
	if err != nil {
		return
	}

	//查找数据库
	find, client := clientmanager.RegisterPhone(phone.Phone, phone.Code)

	var msg model.Message
	if find {
		b, _ := json.Marshal(&client)
		_ = json.Unmarshal(b, &msg.Data)

		msg.Code = model.SOK

	} else {
		msg.Code = model.RegisterFailed
	}

	c.JSON(http.StatusOK, msg)
}

func preregister(c *gin.Context) {
	phone, err := phoneclient.JsonToPhone(c)
	if err != nil {
		return
	}

	//查找数据库

	find, regerr := clientmanager.PreRegisterPhone(phone.Phone)

	var msg model.Message
	if find {
		msg.Code = model.SOK

	} else {
		msg.Code = regerr
	}

	c.JSON(http.StatusOK, msg)

}

func JoinChatByInviteLink(c *gin.Context) {
	phone, err := phoneclient.JsonToLink(c)
	if err != nil {
		return
	}
	ret := clientmanager.Joinlink(phone.Phone, phone.Url)

	var msg model.Message
	if ret == model.SOK {
		msg.Code = model.SOK

	}

	c.JSON(http.StatusOK, msg)
}
