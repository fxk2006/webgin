package controller

import (
	"webgin/tasks"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"fmt"
)

func StartLongTask(c *gin.Context) {
	id := uuid.NewV4()
	go tasks.LongTask(id)
	c.JSON(200, gin.H{
		"uuid": id,
	})
}

func QueryLongTask(c *gin.Context) {
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
