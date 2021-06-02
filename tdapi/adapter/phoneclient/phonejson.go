// Package userclient is client library if you need to call the user Micro-service as a client.
// It provides client library and the data transformation service.
package phoneclient

import (
	"net/http"
	"tdapi/model"

	"github.com/gin-gonic/gin"
)

// GrpcToUser converts from grpc User type to domain Model user type
func JsonToPhone(c *gin.Context) (*model.Phone, error) {

	//声明接收的数据结构
	var jsonData model.Phone
	// 将request的body中数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	return &jsonData, nil
}
