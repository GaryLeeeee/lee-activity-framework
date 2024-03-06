package base

import (
	"errors"
	"lee-activity-framework/lee-activity-api/code"
)

type Req struct {
	ActivityId int    `form:"activity_id"` // 活动id
	Uid        int    `form:"uid"`         // 用户uid
	App        string `form:"app"`         // app
}

type Resp struct {
	Code    int    `json:"code"`    // 返回码 0-成功
	Message string `json:"message"` // 返回文案

	Data interface{} `json:"data"` // 数据
}

// func
func BuildSuccessResp(data interface{}) Resp {
	return Resp{
		Code:    code.Success.Code,
		Message: code.Success.Message,
		Data:    data,
	}
}

func BuildFailResp() Resp {
	return Resp{
		Code:    code.SysError.Code,
		Message: code.SysError.Message,
	}
}

func BuildResCodeResp(e code.ResErr) Resp {
	return Resp{
		Code:    e.Code,
		Message: e.Message,
	}
}

func BuildResp(data interface{}, err error) Resp {
	if err == nil {
		return BuildSuccessResp(data)
	}

	for _, resErr := range code.GetResErrMap() {
		if errors.Is(err, resErr) {
			return BuildResCodeResp(resErr)
		}
	}

	// 不存在则默认系统异常
	return BuildFailResp()
}
