package dao

import (
	"github.com/jinzhu/gorm"
	"lee-activity-framework/lee-activity-lottery/model"
)

func GetLotteryConfig(activityId int) (*model.LeeActivityLotteryConfig, error) {
	var data model.LeeActivityLotteryConfig
	err := _db.Table("lee_activity_lottery_config").Where("activity_id = ?", activityId).First(&data).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return &data, err
}
