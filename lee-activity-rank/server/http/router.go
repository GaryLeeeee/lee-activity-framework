package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/rank")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/list_rank_data", listRankData)                // 查询榜单数据
		apiGroup.Any("/add_rank_data", addRankData)                  // 增加榜单数据
		apiGroup.Any("/get_rank_reward_config", getRankRewardConfig) // 查询榜单奖励配置
	}
	_ = r.Run()
}
