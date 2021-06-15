package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tdapi/adapter/phoneclient"
	"tdapi/clientmanager"
	"tdapi/model"

	"github.com/gin-gonic/gin"
)

const (
	TDURL = "https://t.me/"
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
	var msg model.Message

	phone, err := phoneclient.JsonToLink(c)
	if err != nil {
		return
	}

	strArr := strings.Split(phone.Url, TDURL)
	if len(strArr) <= 0 {
		msg.Code = model.BadRequest
		msg.Err = "不是有效的url"
		c.JSON(http.StatusOK, msg)
		return
	}

	ret, _ := clientmanager.Joinlink(phone.Phone, phone.Url, strArr[1])

	if ret == model.SOK {
		msg.Code = model.SOK

	}
	c.JSON(http.StatusOK, msg)

}

func Getallgroups(c *gin.Context) {
	var msg model.Message
	msg.Code = model.SOK

	agent, err := phoneclient.JsonToAgent(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	groups, err := clientmanager.Getallgroups(agent.Agent)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	b, _ := json.Marshal(&groups)
	err = json.Unmarshal(b, &msg.Data)
	fmt.Println(err)

	c.JSON(http.StatusOK, msg)

}

func Getmegroups(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToMe(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	groups, err := clientmanager.GetMegroups(agent.Name)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	b, _ := json.Marshal(&groups)
	_ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}

//邀请
func Invategroup(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToCreateGroup(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	err = clientmanager.CreateBasicGroup(agent.Account, *agent)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	// b, _ := json.Marshal(&groups)
	// _ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}

//发送
func Sendmessage(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToMe(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	groups, err := clientmanager.GetMegroups(agent.Name)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	b, _ := json.Marshal(&groups)
	_ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}

//发送
func Addtask(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToMe(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	groups, err := clientmanager.GetMegroups(agent.Name)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	b, _ := json.Marshal(&groups)
	_ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}

//发送
func AddContacts(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToContents(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	err = clientmanager.AddContacts(agent)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	// b, _ := json.Marshal(&groups)
	// _ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}

//发送
func GetmeContents(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToMe(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	err = clientmanager.GetmeContents(agent)

	if err != nil {
		msg.Code = model.BadRequest
		c.JSON(http.StatusOK, msg)
		return
	}

	// b, _ := json.Marshal(&groups)
	// _ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}

//发送
func GetgroupContents(c *gin.Context) {
	var msg model.Message

	agent, err := phoneclient.JsonToGroup(c)
	if err != nil {
		return
	}
	fmt.Println(agent)

	// err = clientmanager.GetmeContents(agent)

	// if err != nil {
	// 	msg.Code = model.BadRequest
	// 	c.JSON(http.StatusOK, msg)
	// 	return
	// }

	// b, _ := json.Marshal(&groups)
	// _ = json.Unmarshal(b, &msg.Data)

	c.JSON(http.StatusOK, msg)

}
