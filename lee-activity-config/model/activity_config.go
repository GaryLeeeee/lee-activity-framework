package model

import "time"

type LeeActivityConfig struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 活动id
	Name       string    `gorm:"column:name" json:"name"`                                         // 活动名称
	StartTime  time.Time `gorm:"column:start_time" json:"start_time"`                             // 开始时间
	EndTime    time.Time `gorm:"column:end_time" json:"end_time"`                                 // 结束时间
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	Operator   string    `gorm:"column:operator" json:"operator"`                                 // 操作人
	App        string    `gorm:"column:app_name" json:"app"`                                      // app
}
