package dao

import (
	"github.com/jinzhu/gorm"
	"lee-activity-framework/lee-activity-config/model"
)

func ListAllRunningActivity() ([]model.LeeActivityConfig, error) {
	var data []model.LeeActivityConfig
	err := _db.Table("lee_activity_config").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetActivity(activityId int) (*model.LeeActivityConfig, error) {
	var data model.LeeActivityConfig
	err := _db.Table("lee_activity_config").Where("id = ?", activityId).First(&data).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return &data, nil
}
