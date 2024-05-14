package model

import "time"

type LeeActivityRankConfig struct {
	Id               int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 榜单id
	ActivityId       int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	Type             int       `gorm:"column:type" json:"type"`                                         // 榜单类型 1:总榜 2:日榜
	Name             string    `gorm:"column:name" json:"name"`                                         // 榜单名称
	DataSourceConfig string    `gorm:"column:data_source_config" json:"data_source_config"`             // 数据源配置
	RewardConfig     string    `gorm:"column:reward_config" json:"reward_config"`                       // 榜单奖励配置
	BizConfig        string    `gorm:"column:biz_config" json:"biz_config"`                             // 业务配置
	StartTime        time.Time `gorm:"column:start_time" json:"start_time"`                             // 开始时间
	EndTime          time.Time `gorm:"column:end_time" json:"end_time"`                                 // 结束时间
	CreateTime       time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
}

func (d *LeeActivityRankConfig) TableName() string {
	return "lee_activity_rank_config"
}
