package http

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func isActivityRunning(c *gin.Context) {

}

func getActivity(c *gin.Context) {

}