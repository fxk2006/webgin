package views

import (
	"github.com/gin-gonic/gin"
	"webgin/controller"
)


func init() {
	//gin.DisableConsoleColor()
    // Logging to a file.
    //f, _ := os.Create("gin.log")
    //gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	v1 :=r.Group("/v1")
	v2 :=r.Group("/v2")
	v1.GET("/ping",controller.V1Index)
	v1.POST("/ping",controller.V1POST)
	v1.DELETE("/ping",controller.V1Delete)
	v2.GET("/ping",controller.V1Index)
	r.Run() // listen and serve on 0.0.0.0:8080
}