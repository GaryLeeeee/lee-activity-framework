package model

import "time"

type LeeActivityRankReward struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // id
	ActivityId   int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	RankId       int       `gorm:"column:rank_id" json:"rank_id"`                                   // 榜单id
	Uid          int       `gorm:"column:uid" json:"uid"`                                           // 用户uid
	RewardType   int       `gorm:"column:reward_type" json:"reward_type"`                           // 奖励类型 1:道具 2:积分
	RewardAmount int       `gorm:"column:reward_amount" json:"reward_amount"`                       // 奖励数量
	RankPeriod   string    `gorm:"column:rank_period" json:"rank_period"`                           // 榜单周期
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
}

func (d *LeeActivityRankReward) TableName() string {
	return "lee_activity_rank_reward"
}
