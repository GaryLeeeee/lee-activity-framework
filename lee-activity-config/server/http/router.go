package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/config")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/select_all", selectAllActivityConfig)
	}
	_ = r.Run()
}
