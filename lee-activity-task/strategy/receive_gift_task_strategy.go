package strategy

import (
	"encoding/json"
	"errors"
	"lee-activity-framework/lee-activity-api/logger"
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"lee-activity-framework/lee-activity-api/task"
	"lee-activity-framework/lee-activity-task/dao"
	"lee-activity-framework/lee-activity-task/model"
)

type ReceiveGiftTaskStrategy struct {
}

func (s *ReceiveGiftTaskStrategy) GetTaskType() int {
	return task.TaskTypeReceiveGift
}

func (s *ReceiveGiftTaskStrategy) ConsumeSendGiftMsg(taskId int, msg *kafka.SendGiftMsg) error {
	// 1.查询任务配置
	taskConfig, err := dao.GetTask(taskId)
	if err != nil {
		logger.Errorf("ConsumeSendGiftMsg | taskId:%d GetTask error:%v", taskId, err)
		return err
	}
	if taskConfig == nil || taskConfig.Type != s.GetTaskType() {
		return errors.New("not supported task or type")
	}

	// 2.判断是否需要统计该礼物
	var receiveGiftTaskList []model.ReceiveGiftTask
	err = json.Unmarshal([]byte(taskConfig.TaskConfig), &receiveGiftTaskList)
	if err != nil {
		logger.Errorf("ConsumeSendGiftMsg | taskId:%d Unmarshal error:%v", taskId, err)
		return err
	}
	if !hasReceiveGiftId(msg.GiftId, receiveGiftTaskList) {
		return nil
	}

	// TODO 3.累计礼物个数

	// TODO 4.判断礼物是否达标，是则完成任务
	return nil
}

func hasReceiveGiftId(giftId int, receiveGiftTaskList []model.ReceiveGiftTask) bool {
	for _, receiveGiftTask := range receiveGiftTaskList {
		if receiveGiftTask.GiftId == giftId {
			return true
		}
	}
	return false
}

func (s *ReceiveGiftTaskStrategy) ConsumeSendMessageMsg(taskId int, msg *kafka.SendMessageMsg) error {
	return nil
}
