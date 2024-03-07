package service

import (
	"encoding/json"
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/code"
	"lee-activity-framework/lee-activity-api/logger"
	"lee-activity-framework/lee-activity-api/lottery"
	"lee-activity-framework/lee-activity-api/prop"
	"lee-activity-framework/lee-activity-lottery/dao"
	"lee-activity-framework/lee-activity-lottery/model"
	"math/rand"
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

func UseLotteryTicket(req lottery.UseLotteryTicketReq) (*lottery.UseLotteryTicketResp, error) {
	// 1.分布式锁
	lockUuid, err := dao.UidLock(req.Uid)
	if err != nil {
		return nil, code.FrequentOperation
	}
	defer func() {
		err = dao.UidUnlock(req.Uid, lockUuid)
		if err != nil {
			logger.Errorf("UseLotteryTicket | uid:%d unlock error:%v", req.Uid, err)
		}
	}()

	// 2.判断活动是否存在
	isRunning, err := activity.IsActivityRunning(req.ActivityId)
	if err != nil {
		return nil, err
	}
	if !isRunning {
		return nil, code.ActivityNotExists
	}

	// 3.查询活动对应的抽奖配置
	lotteryConfig, err := dao.GetLotteryConfig(req.ActivityId)
	if err != nil {
		return nil, code.ActivityNotExists
	}

	// 4.查询抽奖券数量
	lotteryAmount, _ := prop.GetPropAmount(req.Uid, req.ActivityId, lotteryConfig.LotteryPropId)
	if lotteryAmount <= 0 {
		return nil, code.LotteryTicketNotEnough
	}
	logger.Infof("UseLotteryTicket | uid:%d", req.Uid)

	// 5.获取一个随机奖励
	var rewardConfigList []model.LotteryRewardConfig
	err = json.Unmarshal([]byte(lotteryConfig.RewardConfig), &rewardConfigList)
	if err != nil {
		return nil, code.SysError
	}
	randomRewardConfig, err := getRandomRewardConfig(rewardConfigList)
	if err != nil {
		return nil, code.UseLotteryTicketErr
	}

	// TODO：6.发放奖励
	logger.Info("UseLotteryTicket | uid:%d getRandomRewardConfig result:%v", req.Uid, randomRewardConfig)
	return &lottery.UseLotteryTicketResp{
		LotteryReward: nil,
	}, nil
}

func getRandomRewardConfig(rewardConfigList []model.LotteryRewardConfig) (*model.LotteryRewardConfig, error) {
	// 先计算总概率
	totalProbability := 0
	for _, value := range rewardConfigList {
		totalProbability += (int)((value.Probability) * 10000)
	}

	// 再取一个随机概率
	randomProbability := rand.Intn(totalProbability)
	tempTotalProbability := 0
	for _, value := range rewardConfigList {
		before := tempTotalProbability
		after := tempTotalProbability + int(value.Probability*10000)
		if randomProbability >= before && randomProbability <= after {
			return &value, nil
		}

		// 当前配置未被本次随机命中
		tempTotalProbability += int(value.Probability * 10000)
	}
	return nil, nil
}
