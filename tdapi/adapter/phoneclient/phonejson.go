// Package userclient is client library if you need to call the user Micro-service as a client.
// It provides client library and the data transformation service.
package phoneclient

import (
	"net/http"
	"tdapi/model"

	"github.com/gin-gonic/gin"
)

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToPhone(c *gin.Context) (*model.RegPhone, error) {

	//声明接收的数据结构
	var jsonData model.RegPhone
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToInvated(c *gin.Context) (*model.Invated, error) {

	//声明接收的数据结构
	var jsonData model.Invated
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToLink(c *gin.Context) (*model.Linkurl, error) {

	//声明接收的数据结构
	var jsonData model.Linkurl
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToAgent(c *gin.Context) (*model.Agent, error) {

	//声明接收的数据结构
	var jsonData model.Agent
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToMe(c *gin.Context) (*model.Me, error) {

	//声明接收的数据结构
	var jsonData model.Me
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToMessage(c *gin.Context) (*model.Text, error) {

	//声明接收的数据结构
	var jsonData model.Text
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToCreateGroup(c *gin.Context) (*model.Friends, error) {

	//声明接收的数据结构
	var jsonData model.Friends
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		err = jsonData.Validate()
		if err != nil {
			return nil, err
		}
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToContents(c *gin.Context) (*model.AddContacts, error) {

	//声明接收的数据结构
	var jsonData model.AddContacts
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToGroup(c *gin.Context) (*model.SaveGroup, error) {

	//声明接收的数据结构
	var jsonData model.SaveGroup
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToChat(c *gin.Context) (*model.SaveChat, error) {

	//声明接收的数据结构
	var jsonData model.SaveChat
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToTask(c *gin.Context) (*model.Me, error) {

	//声明接收的数据结构
	var jsonData model.Me
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		var msg model.Message
		msg.Code = 400
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusOK, msg)
		return nil, err
	}
	return &jsonData, nil
}
