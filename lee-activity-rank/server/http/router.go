package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/rank")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/list_rank_by_activity", listRankByActivity)     // 查询活动对应的榜单列表(不含数据)
		apiGroup.Any("/list_rank_data", listRankData)                  // 查询榜单数据
		apiGroup.Any("/add_rank_data", addRankData)                    // 增加榜单数据
		apiGroup.Any("/list_rank_reward_config", listRankRewardConfig) // 查询榜单奖励配置
	}
	_ = r.Run()
}
