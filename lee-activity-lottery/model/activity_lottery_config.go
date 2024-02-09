package model

import "time"

type LeeActivityLotteryConfig struct {
	Id            int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 抽奖id
	Name          string    `gorm:"column:name" json:"name"`                                         // 抽奖名称
	LotteryPropId int       `gorm:"column:lottery_prop_id" json:"lottery_prop_id"`                   // 抽奖券id
	RewardConfig  string    `gorm:"column:reward_config" json:"reward_config"`                       // 奖励配置
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	App           string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityLotteryConfig) TableName() string {
	return "lee_activity_lottery_config"
}
