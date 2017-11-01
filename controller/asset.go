package controller

import (
	"github.com/gin-gonic/gin"
	"webgin/model"
	"github.com/jinzhu/gorm"
	"webgin/util"
	"fmt"
	"strconv"
	"strings"
)

type PublicUtil struct {
	DB  *gorm.DB
	Log *util.Log
}

func (this PublicUtil) GetJson(c *gin.Context, T interface{}) {
	err := c.BindJSON(&T)
	if err != nil {
		this.Log.Debug(err)
		c.AbortWithError(400, err)
	}
}

type AssetController struct {
	PublicUtil
}

func (this AssetController) Get(c *gin.Context) {
	//cmdb/asset/*id
	id := c.Param("id")
	if id != "/" {
		if intId, err := strconv.Atoi(strings.TrimLeft(id,"/")); err == nil {
			var asset model.Asset
			this.DB.First(&asset, intId)
			c.IndentedJSON(200, asset.ToMapJson())
			return
		}else {
			c.IndentedJSON(400, gin.H{
				"reason":err,
			})
			return
		}

	}
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
func (this AssetController) Post(c *gin.Context) {
	var asset model.Asset
	c.BindJSON(&asset)
	fmt.Println(asset)
	c.JSON(200, asset)
}
func (this AssetController) Put(c *gin.Context) {
	c.String(200, "hello world")
}
func (this AssetController) Delete(c *gin.Context) {
	id := c.Param("id")
	this.DB.Where("id = ?", id).Delete(&model.Asset{})
	c.JSON(200, gin.H{
		"Result": "",
	})
}
