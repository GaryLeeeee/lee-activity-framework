package task

import (
	"lee-activity-framework/lee-activity-api/base"
	"time"
)

// 查询任务详情

// 查询任务配置
type TaskConfig struct {
	ActivityId int    `json:"activity_id"` // 活动id
	Name       string `json:"name"`        // 任务名称
	Type       int    `json:"type"`        // 任务类型
}
type GetTaskConfigReq struct {
	base.Req
	TaskId int `json:"task_id"` // 任务id
}
type GetTaskConfigResp struct {
	TaskConfig TaskConfig `json:"task_config"` // 任务配置
}

// 查询任务奖励配置
type RewardConfig struct {
	RewardType   int `json:"reward_type"`   // 奖励类型
	RewardId     int `json:"reward_id"`     // 奖励id
	RewardAmount int `json:"reward_amount"` // 奖励数量
}
type GetTaskRewardConfigReq struct {
	base.Req
	TaskId int `json:"task_id"`
}
type GetTaskRewardConfigResp struct {
	TaskRewardConfigList []RewardConfig `json:"reward_config_list"` // 奖励配置
}

// 查询活动对应的任务列表
type ListTaskByActivityReq struct {
	base.Req
}
type ListTaskByActivityResp struct {
	TaskList []TaskConfig `json:"task_list"` // 任务列表
}

// 查询任务流水
type TaskHistory struct {
	Uid     int `json:"uid"`      // 用户uid
	PeerUid int `json:"peer_uid"` // 对方uid(如有)
	// taskId
	TaskType    int       `json:"task_type"`    // 任务类型
	FinishCount int       `json:"finish_count"` // 任务完成次数
	CreateTime  time.Time `json:"create_time"`  // 创建时间
}
type ListTaskHistoryReq struct {
	base.Req
	TaskId int `json:"task_id"` // 任务id
}
type ListTaskHistoryResp struct {
	TaskHistoryList []TaskHistory `json:"task_history_list"` // 任务流水
}
