package controller

import "github.com/gin-gonic/gin"
import (
	"webgin/model"
	"strconv"
)

func V1Index(c *gin.Context) {
	var product []model.Products
	model.DB.Find(&product)
	c.JSON(200, product)
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