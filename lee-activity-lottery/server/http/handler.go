package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/lottery"
	"lee-activity-framework/lee-activity-lottery/service"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getLotteryTicketAmount(c *gin.Context) {
	var req lottery.GetLotteryTicketAmountReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.GetLotteryTicketAmount(req)))
}

func useLotteryTicket(c *gin.Context) {

}

func grantLotteryTicket(c *gin.Context) {

}

func getLotteryHistory(c *gin.Context) {

}

func listLotteryRewardConfig(c *gin.Context) {

}
