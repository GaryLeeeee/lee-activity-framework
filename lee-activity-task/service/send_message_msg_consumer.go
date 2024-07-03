package service

import (
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/logger"
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"lee-activity-framework/lee-activity-api/task"
	"lee-activity-framework/lee-activity-task/dao"
	"lee-activity-framework/lee-activity-task/strategy"
)

func ConsumeSendMessageMsg(msg *kafka.SendMessageMsg) {
	// 1.Kafka幂等
	if dao.CheckKafkaUUIDExists(msg.Uuid) {
		logger.Warnf("ConsumeSendMessageMsg | kafka uuid exist, msg:%v", msg)
		return
	}
	// 2.查询所有运行中的活动
	allRunningActivity, err := activity.ListAllRunningActivity()
	if err != nil {
		logger.Warnf("ConsumeSendMessageMsg | ListAllRunningActivity error:%v", msg)
		return
	}
	if len(allRunningActivity) == 0 {
		return
	}
	// TODO 3.分布式锁
	// 4.查询需要统计消息的任务
	for _, runningActivity := range allRunningActivity {
		taskList, err := dao.ListTaskByActivityId(runningActivity.Id)
		if err != nil {
			logger.Errorf("ConsumeSendMessageMsg | activityId:%d ListTaskByActivity error:%v", runningActivity.Id, err)
			continue
		}
		for _, taskItem := range taskList {
			if taskItem.Type == task.TaskTypeSendMessage {
				_ = strategy.GetTaskStrategy(taskItem.Type).ConsumeSendMessageMsg(taskItem.Id, msg)
			}
		}
	}
}
