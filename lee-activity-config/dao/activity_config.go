package dao

import "lee-activity-framework/lee-activity-config/model"

func SelectAllActivityConfig() ([]model.LeeActivityConfig, error) {
	var data []model.LeeActivityConfig
	err := _db.Table("lee_activity_config").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
