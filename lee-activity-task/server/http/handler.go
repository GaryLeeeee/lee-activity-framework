package http

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getTaskDetail(c *gin.Context) {

}

func getTaskConfig(c *gin.Context) {

}

func getTaskRewardConfig(c *gin.Context) {

}

func listTaskByActivity(c *gin.Context) {

}

func listTaskHistory(c *gin.Context) {

}

func getTaskFinishCount(c *gin.Context) {

}
