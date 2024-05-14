package strategy

import (
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"lee-activity-framework/lee-activity-api/rank"
	"lee-activity-framework/lee-activity-rank/model"
	"time"
)

type DayRankStrategy struct {
}

func (s DayRankStrategy) GetRankType() int {
	return rank.RankTypeDay
}

func (s DayRankStrategy) GetRankPeriod(date time.Time) string {
	return date.Format("20060102")
}

func (s DayRankStrategy) ConsumeSendGiftEvent(rankId int, msg kafka.SendGiftMsg) error {
	// TODO
	return nil
}

func (s DayRankStrategy) Convert2SubRank(rankConfig *model.LeeActivityRankConfig) []rank.SubRank {
	subRankList := make([]rank.SubRank, 0)
	if rankConfig == nil || rankConfig.Type != s.GetRankType() {
		return subRankList
	}
	current := rankConfig.StartTime
	for current.Before(rankConfig.EndTime) {
		endOfDay := time.Date(current.Year(), current.Month(), current.Day(), 23, 59, 59, 0, current.Location())
		if endOfDay.After(rankConfig.EndTime) {
			endOfDay = rankConfig.EndTime
		}
		subRankList = append(subRankList, rank.SubRank{
			Id:         rankConfig.Id,
			Name:       rankConfig.Name,
			Type:       rankConfig.Type,
			Period:     current.Format("20060102"),
			StartTime:  current,
			EndTime:    endOfDay,
			ActivityId: rankConfig.ActivityId,
		})
		current = endOfDay.Add(time.Second)
	}
	return subRankList
}
