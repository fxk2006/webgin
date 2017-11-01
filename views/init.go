package views

import (
	"github.com/gin-gonic/gin"
	"webgin/controller"
	"webgin/global"
	"webgin/model"
)

var Engine *gin.Engine

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.GLog.Debug("hello")
		c.Set("request", "clinet_request")
		c.Next()
		global.GLog.Info("zhongjian jian")
	}
}

func Init() {

	Engine = gin.Default() //gin 路由初始化
	Engine.LoadHTMLGlob("templates/*")
	Engine.Use(MiddleWare()) //拦截所有url请求的中间件

	cmdb := Engine.Group("/cmdb")
	task := Engine.Group("/task")
	public := controller.PublicUtil{
		DB:  model.MasterDB,
		Log: global.GLog,
	}
	{
		//cmdb前缀的请求
		{
			//cmdb/asset
			asset := &controller.AssetController{public}
			cmdb.GET("/asset/*id", asset.Get)
			cmdb.PUT("/asset", asset.Put)
			cmdb.POST("/asset", asset.Post)
			cmdb.DELETE("/asset/:id", asset.Delete)
		}
		{
			//cmdb/host
			host := &controller.HostController{public}
			cmdb.GET("/host", host.Get)
			cmdb.PUT("/host", host.Put)
			cmdb.POST("/host", host.Post)
			cmdb.DELETE("/host/:id", host.Delete)
		}
	}

	{
		//task前缀的请求
		task.GET("/longtask", controller.StartLongTask)
		task.GET("/longtask/:uuid", controller.QueryLongTask)
	}

}

func Start() {
	Init()
	ip, _ := global.Config.String("server", "listen")
	port, _ := global.Config.String("server", "port")
	addr := ip + ":" + port
	Engine.Run(addr)
}
