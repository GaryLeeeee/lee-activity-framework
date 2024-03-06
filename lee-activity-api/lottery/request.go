package lottery

import (
	"lee-activity-framework/lee-activity-api/base"
	"time"
)

// 查询抽奖券数量
type GetLotteryTicketAmountReq struct {
	base.Req
}
type GetLotteryTicketAmountResp struct {
	LotteryTicketAmount int `json:"lottery_ticket_amount"` // 抽奖券数量
}

// 使用抽奖券(抽奖)
type Reward struct {
	Id     int    `json:"id"`
	Type   int    `json:"type"`
	Amount int    `json:"amount"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}
type UseLotteryTicketReq struct {
	base.Req
}
type UseLotteryTicketResp struct {
	LotteryReward []Reward `json:"lottery_reward"` // 抽奖奖励
}

// 发放抽奖券
type GrantLotteryTicketReq struct {
	base.Req
	GrantAmount int `json:"grant_amount"` // 发放数量
}
type GrantLotteryTicketResp struct {
}

// 查询抽奖流水
type LotteryHistory struct {
	RewardType   int       `json:"reward_type"`   // 奖励类型
	RewardAmount int       `json:"reward_amount"` // 奖励数量
	CreateTime   time.Time `json:"create_time"`
}
type ListLotteryHistoryReq struct {
	base.Req
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
type ListLotteryHistoryResp struct {
	LotteryHistoryList []LotteryHistory `json:"lottery_history_list"` // 抽奖流水
}

// 查询抽奖奖励配置
type RewardConfig struct {
	RewardType   int     `json:"reward_type"`   // 奖励类型
	RewardId     int     `json:"reward_id"`     // 奖励id
	RewardAmount int     `json:"reward_amount"` // 奖励数量
	Probability  float64 `json:"probability"`   // 奖励概率
}
type ListLotteryRewardConfigReq struct {
	base.Req
}
type ListLotteryRewardConfigResp struct {
	LotteryRewardConfigList []RewardConfig `json:"lottery_reward_config_list"`
}
