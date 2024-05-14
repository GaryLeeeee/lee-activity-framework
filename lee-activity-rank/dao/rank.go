package dao

import (
	"github.com/jinzhu/gorm"
	"lee-activity-framework/lee-activity-rank/model"
)

func ListRankByActivityId(activityId int) ([]*model.LeeActivityRankConfig, error) {
	var data []*model.LeeActivityRankConfig
	err := _db.Table("lee_activity_rank_config").Where("activity_id = ?", activityId).Find(&data).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return data, err
}

func ListRankData(rankId int, period string, page int, pageSize int) ([]*model.LeeActivityRankDetail, error) {
	var data []*model.LeeActivityRankDetail
	err := _db.Table("lee_activity_rank_detail").Where("rank_id = ? AND period = ?", rankId, period).Limit(pageSize).Offset((page - 1) * pageSize).Order("score desc, update_time asc").Find(&data).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return data, nil
}

func GetMyRankData(uid int, rankId int, period string) (*model.LeeActivityRankDetail, error) {
	var data model.LeeActivityRankDetail
	err := _db.Table("lee_activity_rank_detail").Where("uid = ? AND rank_id = ? AND period = ?", uid, rankId, period).First(&data).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &model.LeeActivityRankDetail{
				RankId: rankId,
				Uid:    int(uid),
				Score:  0,
				Period: period,
			}, nil
		}
		return nil, err
	}
	return &data, nil
}
