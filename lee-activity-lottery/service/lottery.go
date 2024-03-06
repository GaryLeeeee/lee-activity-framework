package service

import (
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/code"
	"lee-activity-framework/lee-activity-api/lottery"
	"lee-activity-framework/lee-activity-api/prop"
	"lee-activity-framework/lee-activity-lottery/dao"
)

func GetLotteryTicketAmount(req lottery.GetLotteryTicketAmountReq) (*lottery.GetLotteryTicketAmountResp, error) {
	// 1.判断活动是否存在
	isRunning, err := activity.IsActivityRunning(req.ActivityId)
	if err != nil {
		return nil, err
	}
	if !isRunning {
		return nil, code.ActivityNotExists
	}

	// 2.查询活动对应的抽奖配置
	lotteryConfig, err := dao.GetLotteryConfig(req.ActivityId)
	if err != nil {
		return nil, code.ActivityNotExists
	}

	// 3.查询抽奖券数量
	lotteryAmount, _ := prop.GetPropAmount(req.Uid, req.ActivityId, lotteryConfig.LotteryPropId)
	return &lottery.GetLotteryTicketAmountResp{
		LotteryTicketAmount: lotteryAmount,
	}, nil
}
