package views

import (
	"github.com/gin-gonic/gin"
	"webgin/controller"
	"webgin/global"
)

var Engine *gin.Engine

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.GLog.Debug("hello")
		c.Set("request", "clinet_request")
		c.Next()
		global.GLog.Warning("hello world warinng")
	}
}

func init() {
	//gin.DisableConsoleColor()
	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	Engine = gin.Default()
	Engine.LoadHTMLGlob("templates/*")
	Engine.Use(MiddleWare())
	v1 := Engine.Group("/v1")
	v2 := Engine.Group("/v2")
	{
		v1.GET("/ping", controller.V1Index)
		v1.POST("/ping", controller.V1POST)
		v1.DELETE("/ping", controller.V1Delete)
	}
	{
		v2.GET("/ping/:uuid", controller.V2Index)
	}

}
