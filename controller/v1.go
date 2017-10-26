package controller

import "github.com/gin-gonic/gin"
import (
	"webgin/model"
	"strconv"
	"net/http"
	"webgin/global"
)


func V1Index(c *gin.Context) {
	var product []model.Products
	model.DB.Find(&product)
	c.JSON(200, product)
}

func V2Index(c *gin.Context) {
	global.GLog.Debug("hello Products1")
	global.GLog.Info("hello Products2")
	global.GLog.Warning("hello Products2")
	var product []model.Products
	model.DB.Find(&product)
	c.HTML(http.StatusOK,"index.html",gin.H{
		"title":"yangyang",
		"name":[]string{"1","2"},
	})
}

func V1POST(c *gin.Context) {
	code := c.DefaultQuery("code", "deft")
	price := c.Query("price")
	tmp,_ := strconv.Atoi(price)
	p := uint(tmp)
	model.DB.Create(&model.Products{
		Code:code,
		Price:p,
	})
	c.JSON(200, gin.H{
		"ok":"ok",
	})
}

func V1Delete(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		var product model.Products
		model.DB.First(&product, id)
		model.DB.Delete(&product)
		c.JSON(201, gin.H{
			"ok":"ppp",
		})
	}
}
func V1Any(c *gin.Context) {
	name := c.Param("name")
	c.String(200,name)
}