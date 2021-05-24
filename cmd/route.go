package main

import (
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {

	v1 := r.Group("/v1")
	GetHandlersV1(v1)
}

func GetHandlersV1(r *gin.RouterGroup) {
	r.Use(version())
	public := r.Group("/")

	/**
	 * Point urls
	 */
	public.POST("/getall/", createPointHandler)
	public.GET("/addphone/", createPhoneHandler)
}

/**
 * Middlewares functions
 */

func version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("version", "2")
		c.Next()
	}
}

func createPointHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func createPhoneHandler(c *gin.Context) {
	request := fetchPointsForEventsRequest{}

	c.JSON(200, gin.H{
		"message": "test",
	})
}
