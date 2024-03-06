package dao

import (
	"github.com/jinzhu/gorm"
	"lee-activity-framework/lee-activity-prop/model"
)

func GetPropAmountByUidAndPropId(uid int, activityId int, propId int) (int, error) {
	var data model.LeeActivityPropDetail
	err := _db.Table("lee_activity_prop_detail").Where("uid = ? AND activity_id = ? AND prop_id = ?", uid, activityId, propId).First(&data).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, nil
		}
		return 0, err
	}

	return data.PropAmount, nil
}
