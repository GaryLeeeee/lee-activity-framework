package http

import (
	"github.com/gin-gonic/gin"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/prop"
	"lee-activity-framework/lee-activity-prop/service"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getPropAmount(c *gin.Context) {
	var req prop.GetPropAmountReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, base.BuildResp(nil, err))
	}

	c.JSON(http.StatusOK, base.BuildResp(service.GetPropAmount(req)))

}

func useProp(c *gin.Context) {

}

func grantProp(c *gin.Context) {

}

func exchangeProp(c *gin.Context) {

}
