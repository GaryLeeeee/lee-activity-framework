package http

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getLotteryTicketAmount(c *gin.Context) {

}

func useLotteryTicket(c *gin.Context) {

}

func grantLotteryTicket(c *gin.Context) {

}

func getLotteryHistory(c *gin.Context) {

}

func getLotteryRewardConfig(c *gin.Context) {

}
