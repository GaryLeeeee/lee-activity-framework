package http

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	apiGroup := r.Group("/api/activity/lottery")
	{
		apiGroup.Any("/ping", ping)
		apiGroup.Any("/get_lottery_ticket_amount", getLotteryTicketAmount)   // 查询抽奖券数量
		apiGroup.Any("/use_lottery_ticket", useLotteryTicket)                // 使用抽奖券(抽奖)
		apiGroup.Any("/grant_lottery_ticket", grantLotteryTicket)            // 发放抽奖券
		apiGroup.Any("/list_lottery_history", getLotteryHistory)             // 查询抽奖流水
		apiGroup.Any("/list_lottery_reward_config", listLotteryRewardConfig) // 查询抽奖奖励配置
	}
	_ = r.Run(":8081")
}
