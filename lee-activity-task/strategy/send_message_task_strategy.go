package strategy

import "lee-activity-framework/lee-activity-api/task"

type SendMessageTaskStrategy struct {
}

func (s *SendMessageTaskStrategy) GetTaskType() int {
	return task.TaskTypeSendMessage
}
