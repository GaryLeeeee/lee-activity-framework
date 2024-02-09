package model

import "time"

type LeeActivityTaskHistory struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                  // 任务流水id
	ActivityId  int       `gorm:"column:activity_id" json:"activity_id"`                           // 活动id
	Uid         int       `gorm:"column:uid" json:"uid"`                                           // 用户uid
	PeerUid     int       `gorm:"column:peer_uid" json:"peer_uid"`                                 // 对方uid(如有)
	TaskId      int       `gorm:"column:task_id" json:"task_id"`                                   // 任务id
	TaskType    int       `gorm:"column:task_type" json:"task_type"`                               // 任务类型
	FinishCount int       `gorm:"column:finish_count" json:"finish_count"`                         // 任务完成次数
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	App         string    `gorm:"column:app" json:"app"`                                           // app
}

func (d *LeeActivityTaskHistory) TableName() string {
	return "lee_activity_task_history"
}
