package service

import (
	"lee-activity-framework/lee-activity-api/prop"
	"lee-activity-framework/lee-activity-prop/dao"
)

func GetPropAmount(req prop.GetPropAmountReq) (*prop.GetPropAmountResp, error) {
	propAmount, err := dao.GetPropAmountByUidAndPropId(req.Uid, req.ActivityId, req.PropId)
	if err != nil {
		return nil, err
	}
	return &prop.GetPropAmountResp{
		PropAmount: propAmount,
	}, nil
}
