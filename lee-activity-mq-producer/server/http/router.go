package http

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/mq")
	{
		apiGroup.Any("/send_gift", sendGift)
		apiGroup.Any("/send_message", sendMessage)
	}
	_ = r.Run(":8888")
}
