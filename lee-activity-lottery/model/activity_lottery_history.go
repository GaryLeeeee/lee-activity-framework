package model

import "time"

type LeeActivityLotteryHistory struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 抽奖流水id
	ActivityId   int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	LotteryId    int       `gorm:"column:lottery_id" json:"lottery_id"`                             // 抽奖id
	Uid          int       `gorm:"column:uid" json:"uid"`                                           // 用户uid
	RewardType   int       `gorm:"column:reward_type" json:"reward_type"`                           // 奖励类型
	RewardAmount int       `gorm:"column:reward_amount" json:"reward_amount"`                       // 奖励数量
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	App          string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityLotteryHistory) TableName() string {
	return "lee_activity_lottery_history"
}
