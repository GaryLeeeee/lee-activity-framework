package activity

import (
	"lee-activity-framework/lee-activity-api/base"
	"time"
)

// 查询活动是否进行中
type IsActivityRunningReq struct {
	base.Req
}
type IsActivityRunningResp struct {
	IsRunning bool `json:"is_running"` // 是否运行中
}

// 查询某个活动
type Activity struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
type GetActivityReq struct {
	base.Req
}
type GetActivityResp struct {
	Activity *Activity `json:"activity"`
}

// 查询所有运行中活动
type ListAllRunningActivityReq struct {
	base.Req
}
type ListAllRunningActivityResp struct {
	AllRunningActivityList []Activity `json:"all_running_activity_list"` // 所有运行中活动
}
