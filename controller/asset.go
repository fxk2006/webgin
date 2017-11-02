package controller

import (
	"github.com/gin-gonic/gin"
	"webgin/model"
	"github.com/jinzhu/gorm"
	"webgin/util"
	"io/ioutil"
	"encoding/json"
)

type PublicUtil struct {
	DB  *gorm.DB
	Log *util.Log
}

func (this PublicUtil)Head(c *gin.Context)  {
	c.Status(200)
}
func (this PublicUtil)Options(c *gin.Context)  {
	c.Status(200)
}
func (this PublicUtil)Patch(c *gin.Context)  {
	c.Status(200)
}
func (this PublicUtil)Connect(c *gin.Context)  {
	c.Status(200)
}
func (this PublicUtil)Trace(c *gin.Context)  {
	c.Status(200)
}
//func (this PublicUtil)Any(c *gin.Context)  {
//	switch c.Request.Method{
//	case "GET":
//		this.Get(c)
//
//
//	}
//}


type AssetController struct {
	PublicUtil
}

func (this *AssetController) Get(c *gin.Context) {

	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	var assets model.ListAsset
	this.DB.Limit(limit).Offset(offset).Find(&assets)
	this.Log.Debug(assets)

	c.IndentedJSON(200, gin.H{
		"Total":  len(assets),
		"Data":   assets.ToListJson(),
		"Limit":  limit,
		"Offset": offset,
	})
}

func (this *AssetController) Post(c *gin.Context) {

	var assetJson model.AssetJson
	c.BindJSON(&assetJson)
	if this.DB.Create(&assetJson).Error != nil {
		c.JSON(200, assetJson)
		this.Log.Debug("chucuola")
		return
	}
	c.JSON(200, assetJson)
}
func (this *AssetController) Put(c *gin.Context) {
	id := c.Param("id")
	var asset model.Asset
	bytearray, _ := ioutil.ReadAll(c.Request.Body)
	var container gin.H
	json.Unmarshal(bytearray, &container)
	db := this.DB.First(&asset, "id = ?", id)
	for k, v := range container {
		db.Update(k, v)
	}
	c.JSON(200, container)
}
func (this *AssetController) Delete(c *gin.Context) {
	id := c.Param("id")
	this.DB.Where("id = ?", id).Delete(&model.Asset{})
	c.JSON(200, gin.H{
		"Result": "",
	})
}

