package prop

import (
	"encoding/json"
	"fmt"
	"lee-activity-framework/lee-activity-api/base"
	"lee-activity-framework/lee-activity-api/code"
)

func GetPropAmount(uid int, activityId int, propId int) (int, error) {
	uri := "http://127.0.0.1:8082/api/activity/prop/get_prop_amount"
	req := &GetPropAmountReq{
		Req: base.Req{
			Uid:        uid,
			ActivityId: activityId,
		},
		PropId: propId,
	}

	resp := new(base.Resp)
	err := base.HttpPost(uri, req, &resp)
	if err != nil {
		fmt.Printf("GetPropAmount error:%v", err)
		return 0, err
	}
	if !resp.IsSuccess() {
		return 0, code.SysError
	}

	var data GetPropAmountResp
	b, _ := json.Marshal(resp.Data)
	_ = json.Unmarshal(b, &data)
	return data.PropAmount, nil
}
