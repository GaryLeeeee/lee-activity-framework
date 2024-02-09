package model

import "time"

type LeeActivityPropConfig struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 道具id
	Name       string    `gorm:"column:name" json:"name"`                                         // 道具名称
	Type       int       `gorm:"column:type" json:"type"`                                         // 道具类型
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	App        string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityPropConfig) TableName() string {
	return "lee_activity_prop_config"
}
