package strategy

type TaskStrategy interface {
	// 任务类型
	GetTaskType() int
}

var taskStrategyMap map[int]TaskStrategy

func init() {
	taskStrategyMap = make(map[int]TaskStrategy)
	// 注册策略
	sendGiftTaskStrategy := new(SendGiftTaskStrategy)
	taskStrategyMap[sendGiftTaskStrategy.GetTaskType()] = sendGiftTaskStrategy
	receiveGiftTaskStrategy := new(ReceiveGiftTaskStrategy)
	taskStrategyMap[receiveGiftTaskStrategy.GetTaskType()] = receiveGiftTaskStrategy
	sendMessageTaskStrategy := new(SendMessageTaskStrategy)
	taskStrategyMap[sendMessageTaskStrategy.GetTaskType()] = sendMessageTaskStrategy
}
