package model

import "time"

type LeeActivityTaskConfig struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 任务id
	ActivityId   int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	Type         int       `gorm:"column:type" json:"type"`                                         // 任务类型
	Name         string    `gorm:"column:name" json:"name"`                                         // 任务名称
	TaskConfig   string    `gorm:"column:task_config" json:"task_config"`                           // 任务配置
	RewardConfig string    `gorm:"column:reward_config" json:"reward_config"`                       // 奖励配置
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	App          string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityTaskConfig) TableName() string {
	return "lee_activity_task_config"
}
