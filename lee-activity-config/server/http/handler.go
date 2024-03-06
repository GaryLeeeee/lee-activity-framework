package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-config/service"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func isActivityRunning(c *gin.Context) {
	var req activity.IsActivityRunningReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.IsActivityRunning(req)))
}

func getActivity(c *gin.Context) {
	var req activity.GetActivityReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.GetActivity(req)))
}

func listAllRunningActivity(c *gin.Context) {
	var req activity.ListAllRunningActivityReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.ListAllRunningActivity(req)))
}
