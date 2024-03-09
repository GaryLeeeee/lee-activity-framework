package strategy

import "lee-activity-framework/lee-activity-api/task"

type SendGiftTaskStrategy struct {
}

func (s *SendGiftTaskStrategy) GetTaskType() int {
	return task.TaskTypeSendGift
}
