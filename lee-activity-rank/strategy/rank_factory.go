package strategy

import (
	"lee-activity-framework/lee-activity-api/mq/kafka"
	"lee-activity-framework/lee-activity-api/rank"
	"lee-activity-framework/lee-activity-rank/model"
	"time"
)

type RankFactory interface {
	// 榜单类型
	GetRankType() int
	// 榜单周期
	GetRankPeriod(date time.Time) string
	// 消费送礼流水
	ConsumeSendGiftEvent(rankId int, msg kafka.SendGiftMsg) error
	// 子榜单转换
	Convert2SubRank(rankConfig *model.LeeActivityRankConfig) []rank.SubRank
}

var rankStrategyMap map[int]RankFactory

func GetRankStrategy(rankType int) RankFactory {
	return rankStrategyMap[rankType]
}

func init() {
	rankStrategyMap = make(map[int]RankFactory)
	// 注册策略
	totalRankStrategy := new(TotalRankStrategy)
	rankStrategyMap[totalRankStrategy.GetRankType()] = totalRankStrategy
	dayRankStrategy := new(DayRankStrategy)
	rankStrategyMap[dayRankStrategy.GetRankType()] = dayRankStrategy
}
