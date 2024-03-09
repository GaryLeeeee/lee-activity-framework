package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/task"
	"lee-activity-framework/lee-activity-task/service"
	"net/http"
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

func listTaskRewardConfig(c *gin.Context) {

}

func listTaskByActivity(c *gin.Context) {
	var req task.ListTaskByActivityReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.ListTaskByActivity(req)))
}

func listTaskHistory(c *gin.Context) {

}

func getTaskFinishCount(c *gin.Context) {

}
