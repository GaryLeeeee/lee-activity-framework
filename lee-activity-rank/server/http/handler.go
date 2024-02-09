package http

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func listRankData(c *gin.Context) {

}

func addRankData(c *gin.Context) {

}

func getRankRewardConfig(c *gin.Context) {

}
