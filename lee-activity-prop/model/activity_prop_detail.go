package model

import "time"

type LeeActivityPropDetail struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 道具详情id
	PropId     int       `gorm:"column:prop_id" json:"prop_id"`                                   // 道具id
	Uid        int       `gorm:"column:uid" json:"uid"`                                           // 用户uid
	PropAmount int       `gorm:"column:prop_amount" json:"prop_amount"`                           // 道具数量
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	App        string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityPropDetail) TableName() string {
	return "lee_activity_prop_detail"
}
