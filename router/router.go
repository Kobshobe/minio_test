package router

import "github.com/gin-gonic/gin"

var engine = gin.Default()

func Init() *gin.Engine {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	engine.POST("/upload", func(c *gin.Context) {
		
	})
	return engine
}
