package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/config")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/is_activity_running", isActivityRunning)            // 查询活动是否进行中
		apiGroup.Any("/get_activity", getActivity)                         // 查询某个活动
		apiGroup.Any("/list_all_running_activity", listAllRunningActivity) // 查询所有运行中活动
	}
	_ = r.Run()
}
