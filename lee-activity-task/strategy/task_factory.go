package strategy

import "lee-activity-framework/lee-activity-api/mq/kafka"

type TaskFactory interface {
	// 任务类型
	GetTaskType() int
	// 消费礼物流水
	ConsumeSendGiftMsg(taskId int, msg *kafka.SendGiftMsg) error
	// 消费消息流水
	ConsumeSendMessageMsg(taskId int, msg *kafka.SendMessageMsg) error
}

var taskStrategyMap map[int]TaskFactory

func GetTaskStrategy(rankType int) TaskFactory {
	return taskStrategyMap[rankType]
}

func init() {
	taskStrategyMap = make(map[int]TaskFactory)
	// 注册策略
	sendGiftTaskStrategy := new(SendGiftTaskStrategy)
	taskStrategyMap[sendGiftTaskStrategy.GetTaskType()] = sendGiftTaskStrategy
	receiveGiftTaskStrategy := new(ReceiveGiftTaskStrategy)
	taskStrategyMap[receiveGiftTaskStrategy.GetTaskType()] = receiveGiftTaskStrategy
	sendMessageTaskStrategy := new(SendMessageTaskStrategy)
	taskStrategyMap[sendMessageTaskStrategy.GetTaskType()] = sendMessageTaskStrategy
}
