package prop

import "lee-activity-framework/lee-activity-api/base"

// 查询道具数量
type GetPropAmountReq struct {
	base.Req
	PropId int `json:"prop_id"` // 道具id
}
type GetPropAmountResp struct {
	PropAmount int `json:"prop_amount"` // 道具数量
}

// 使用道具
type UsePropReq struct {
	base.Req
	PropId int `json:"prop_id"` // 道具id
}
type UsePropResp struct {
}

// 发放道具
type GrantPropReq struct {
	base.Req
	PropId     int `json:"prop_id"`     // 道具id
	PropAmount int `json:"prop_amount"` // 道具数量
}
type GrantPropResp struct {
}

// 兑换道具(多换一)
type ExchangePropReq struct {
	base.Req
	SourcePropIdList []int `json:"source_prop_id_list"` // 原道具id
	TargetPropIdList []int `json:"target_prop_id_list"` // 目标道具id
	ExchangeCount    int   `json:"exchange_count"`      // 兑换次数
}
type ExchangePropResp struct {
}
