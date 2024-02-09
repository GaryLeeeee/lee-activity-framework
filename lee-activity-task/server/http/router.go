package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/task")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/get_task_detail", getTaskDetail)              // 查询任务详情
		apiGroup.Any("/get_task_config", getTaskConfig)              // 查询任务配置
		apiGroup.Any("/get_task_reward_config", getTaskRewardConfig) // 查询任务奖励配置
		apiGroup.Any("/list_task_by_activity", listTaskByActivity)   // 查询活动对应的任务列表
		apiGroup.Any("/list_task_history", listTaskHistory)          // 查询任务流水
		apiGroup.Any("/get_task_finish_count", getTaskFinishCount)   // 查询任务完成次数
	}
	_ = r.Run()
}
