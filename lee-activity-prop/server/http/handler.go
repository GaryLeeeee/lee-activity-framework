package http

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getPropAmount(c *gin.Context) {

}

func useProp(c *gin.Context) {

}

func grantProp(c *gin.Context) {

}

func exchangeProp(c *gin.Context) {

}
