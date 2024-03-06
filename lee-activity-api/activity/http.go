package activity

import (
	"encoding/json"
	"fmt"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/code"
)

func IsActivityRunning(activityId int) (bool, error) {
	// TODO：环境判断
	uri := "http://127.0.0.1:8080/api/activity/config/is_activity_running"
	req := &IsActivityRunningReq{
		Req: base.Req{
			ActivityId: activityId,
		},
	}

	resp := new(base.Resp)
	err := base.HttpPost(uri, req, &resp)
	if err != nil {
		fmt.Printf("IsActivityRunning error:%v", err)
		return false, err
	}
	if !resp.IsSuccess() {
		return false, code.SysError
	}

	// TODO：优化解析
	var data IsActivityRunningResp
	b, _ := json.Marshal(resp.Data)
	_ = json.Unmarshal(b, &data)
	return data.IsRunning, nil
}
