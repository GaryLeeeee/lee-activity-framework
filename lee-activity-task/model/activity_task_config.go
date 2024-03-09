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
}

func (d *LeeActivityTaskConfig) TableName() string {
	return "lee_activity_task_config"
}

// ...
// 送礼任务：累计送出${gift_id}礼物${gift_amount}个即可完成任务
type SendGiftTask struct {
	GiftId     int `json:"gift_id"`
	GiftAmount int `json:"gift_amount"`
}

// 收礼任务：累计收到${gift_id}礼物${gift_amount}个即可完成任务
type ReceiveGiftTask struct {
	GiftId     int `json:"gift_id"`
	GiftAmount int `json:"gift_amount"`
}

// 发消息任务：累计发送${message_count}条消息即可完成任务（is_related为1时则需要发给同一个人）
type SendMessageTask struct {
	MessageCount int `json:"message_count"`
	IsRelated    int `json:"is_related"`
}
