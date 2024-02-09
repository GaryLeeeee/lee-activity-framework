package model

import "time"

type LeeActivityRankConfig struct {
	Id           int       `gorm:"column:id;primary_key" json:"id"`                                 // 榜单id
	ActivityId   int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	Name         string    `gorm:"column:name" json:"name"`                                         // 榜单名称
	Type         int       `gorm:"column:type" json:"type"`                                         // 榜单类型
	RewardConfig string    `gorm:"column:reward_config" json:"reward_config"`                       // 奖励配置
	StartTime    time.Time `gorm:"column:start_time" json:"start_time"`                             // 开始时间
	EndTime      time.Time `gorm:"column:end_time" json:"end_time"`                                 // 结束时间
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	App          string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityRankConfig) TableName() string {
	return "lee_activity_rank_config"
}
