package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/rank"
	"lee-activity-framework/lee-activity-rank/service"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func listRankByActivity(c *gin.Context) {
	var req rank.ListRankByActivityReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}
	c.JSON(http.StatusOK, base.BuildResp(service.ListRankByActivity(req)))
}

func listRankData(c *gin.Context) {
	var req rank.ListRankDataReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}
	c.JSON(http.StatusOK, base.BuildResp(service.ListRankData(req)))
}

func addRankData(c *gin.Context) {

}

func listRankRewardConfig(c *gin.Context) {

}
