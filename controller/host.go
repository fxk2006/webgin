package controller

import (
	"github.com/gin-gonic/gin"
	"webgin/model"
)

type HostController struct {
	PublicUtil
}

func (this HostController) Get(c *gin.Context) {
	hosts := []model.Host{}
	this.DB.Find(&hosts)
	this.Log.Debug(hosts)
	c.IndentedJSON(200, hosts)
}
func (this HostController) Post(c *gin.Context) {
	c.String(200, "hello world")
}

func (this HostController) Put(c *gin.Context) {
	c.String(200, "hello world")
}
func (this HostController) Delete(c *gin.Context) {
	c.String(200, "hello world")
}
