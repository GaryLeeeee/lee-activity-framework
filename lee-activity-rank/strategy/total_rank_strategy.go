package strategy

import (
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"lee-activity-framework/lee-activity-api/rank"
	"lee-activity-framework/lee-activity-rank/model"
	"time"
)

type TotalRankStrategy struct {
}

func (s TotalRankStrategy) GetRankType() int {
	return rank.RankTypeTotal
}

func (s TotalRankStrategy) GetRankPeriod(date time.Time) string {
	return ""
}

func (s TotalRankStrategy) ConsumeSendGiftEvent(rankId int, msg kafka.SendGiftMsg) error {
	// TODO
	return nil
}

func (s TotalRankStrategy) Convert2SubRank(rankConfig *model.LeeActivityRankConfig) []rank.SubRank {
	subRankList := make([]rank.SubRank, 0)
	if rankConfig == nil || rankConfig.Type != s.GetRankType() {
		return subRankList
	}
	subRankList = append(subRankList, rank.SubRank{
		Id:        rankConfig.Id,
		Name:      rankConfig.Name,
		Type:      rankConfig.Type,
		Period:    "",
		StartTime: rankConfig.StartTime,
	})
	return subRankList
}
