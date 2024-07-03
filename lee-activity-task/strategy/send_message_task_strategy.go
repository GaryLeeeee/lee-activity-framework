package strategy

import (
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"lee-activity-framework/lee-activity-api/task"
)

type SendMessageTaskStrategy struct {
}

func (s *SendMessageTaskStrategy) GetTaskType() int {
	return task.TaskTypeSendMessage
}

func (s *SendMessageTaskStrategy) ConsumeSendGiftMsg(rankId int, msg *kafka.SendGiftMsg) error {
	return nil
}

func (s *SendMessageTaskStrategy) ConsumeSendMessageMsg(rankId int, msg *kafka.SendMessageMsg) error {
	return nil
}
