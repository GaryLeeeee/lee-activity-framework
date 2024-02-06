package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-config/dao"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func selectAllActivityConfig(c *gin.Context) {
	configs, err := dao.SelectAllActivityConfig()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    configs,
		})
	}
}
