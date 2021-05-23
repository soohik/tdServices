package main

import (
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	v2 := r.Group("/v1")
	GetHandlersV2(v2)
}

func GetHandlersV2(r *gin.RouterGroup) {
	r.Use(version())
	public := r.Group("/")

	/**
	 * Point urls
	 */
	public.POST("/getall/", createPointHandler)
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
