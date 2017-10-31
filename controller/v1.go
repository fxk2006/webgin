package controller

import (
	"webgin/tasks"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"fmt"
)

func V1Index(c *gin.Context) {
	id := uuid.NewV4()
	go tasks.LongTask(id)
	c.JSON(200, gin.H{
		"uuid": id,
	})
}

func V2Index(c *gin.Context) {
	paramId := c.Param("uuid")
	uuId, err := uuid.FromString(paramId)
	if err != nil {
		c.JSON(200, gin.H{
			"result": err,
		})
	}
	progress, ok := tasks.DBTasks[uuId]
	if ok {
		c.JSON(200, progress)
	} else {
		c.JSON(200, gin.H{
			"result": fmt.Sprintf("%s not found", paramId),
		})
	}
}

func V1POST(c *gin.Context) {
	c.JSON(200, gin.H{
		"ok": "ok",
	})
}

func V1Delete(c *gin.Context) {
	c.JSON(201, gin.H{
		"ok": "ppp",
	})
}
