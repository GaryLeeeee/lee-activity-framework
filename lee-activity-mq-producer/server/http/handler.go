package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/event"
	"lee-activity-framework/lee-activity-mq-producer/service"
	"net/http"
)

func sendGift(c *gin.Context) {
	var req event.SendGiftReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.SendGift(req)))
}

func sendMessage(c *gin.Context) {
	var req event.SendMessageReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.SendMessage(req)))
}
