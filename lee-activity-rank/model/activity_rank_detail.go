package model

import "time"

type LeeActivityRankDetail struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 榜单详情id
	ActivityId int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	RankId     int       `gorm:"column:rank_id" json:"rank_id"`                                   // 榜单id
	Uid        int       `gorm:"column:uid" json:"uid"`                                           // 用户uid
	PeerUid    int       `gorm:"column:peer_uid" json:"peer_uid"`                                 // 对方uid（如有）
	Score      float64   `gorm:"column:score" json:"score"`                                       // 分数
	Period     string    `gorm:"column:period" json:"period"`                                     // 榜单周期（默认总榜）
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
}

func (d *LeeActivityRankDetail) TableName() string {
	return "lee_activity_rank_detail"
}
