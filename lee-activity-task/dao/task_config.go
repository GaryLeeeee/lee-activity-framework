package dao

import (
	"github.com/jinzhu/gorm"
	"lee-activity-framework/lee-activity-task/model"
)

func ListTaskByActivityId(activityId int) ([]*model.LeeActivityTaskConfig, error) {
	var data []*model.LeeActivityTaskConfig
	err := _db.Table("lee_activity_task_config").Where("activity_id = ?", activityId).Find(&data).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return data, err
}
