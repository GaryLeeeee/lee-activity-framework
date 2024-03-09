package strategy

import "lee-activity-framework/lee-activity-api/task"

type ReceiveGiftTaskStrategy struct {
}

func (s *ReceiveGiftTaskStrategy) GetTaskType() int {
	return task.TaskTypeReceiveGift
}
