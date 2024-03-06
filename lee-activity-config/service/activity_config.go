package service

import (
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/code"
	"lee-activity-framework/lee-activity-config/dao"
	"time"
)

func IsActivityRunning(req activity.IsActivityRunningReq) (*activity.IsActivityRunningResp, error) {
	activityById, err := dao.GetActivity(req.ActivityId)
	if err != nil {
		return nil, err
	}
	if activityById == nil {
		return nil, code.ActivityNotExists
	}

	now := time.Now()
	isActivityRunning := false
	if activityById.StartTime.Before(now) && activityById.EndTime.After(now) {
		isActivityRunning = true
	}

	return &activity.IsActivityRunningResp{
		IsRunning: isActivityRunning,
	}, nil
}

func GetActivity(req activity.GetActivityReq) (*activity.GetActivityResp, error) {
	activityById, err := dao.GetActivity(req.ActivityId)
	if err != nil {
		return nil, err
	}
	if activityById == nil {
		return nil, nil
	}

	return &activity.GetActivityResp{
		Activity: &activity.Activity{
			Id:        activityById.Id,
			Name:      activityById.Name,
			StartTime: activityById.StartTime,
			EndTime:   activityById.EndTime,
		},
	}, nil
}

func ListAllRunningActivity(req activity.ListAllRunningActivityReq) (*activity.ListAllRunningActivityResp, error) {
	resp := &activity.ListAllRunningActivityResp{
		AllRunningActivityList: make([]activity.Activity, 0),
	}

	allRunningActivity, err := dao.ListAllRunningActivity()
	if err != nil {
		return nil, err
	}
	if len(allRunningActivity) == 0 {
		return resp, nil
	}
	for _, runningActivity := range allRunningActivity {
		resp.AllRunningActivityList = append(resp.AllRunningActivityList, activity.Activity{
			Id:        runningActivity.Id,
			Name:      runningActivity.Name,
			StartTime: runningActivity.StartTime,
			EndTime:   runningActivity.EndTime,
		})
	}

	return resp, nil
}
