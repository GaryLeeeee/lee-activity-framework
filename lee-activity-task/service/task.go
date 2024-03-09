package service

import (
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/code"
	"lee-activity-framework/lee-activity-api/task"
	"lee-activity-framework/lee-activity-task/dao"
)

func ListTaskByActivity(req task.ListTaskByActivityReq) (*task.ListTaskByActivityResp, error) {
	// 1.判断活动是否存在
	isRunning, err := activity.IsActivityRunning(req.ActivityId)
	if err != nil {
		return nil, err
	}
	if !isRunning {
		return nil, code.ActivityNotExists
	}

	// 2.查询活动任务配置
	taskList, err := dao.ListTaskByActivityId(req.ActivityId)
	if err != nil {
		return nil, code.GetTaskErr
	}
	if len(taskList) <= 0 {
		return &task.ListTaskByActivityResp{}, nil
	}

	// 3.查询活动任务进度&完成次数
	resp := &task.ListTaskByActivityResp{
		TaskList: make([]task.TaskConfig, 0),
	}
	for _, _ = range taskList {
		// ...
	}

	return resp, nil
}
