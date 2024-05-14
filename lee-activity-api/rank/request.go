package rank

import (
	"lee-activity-framework/lee-activity-api/base"
	"time"
)

// 查询活动对应的榜单列表(不含数据)
type SubRank struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Type      int       `json:"type"`   // 榜单类型 1:总榜 2:日榜
	Period    string    `json:"period"` // 榜单周期(总榜为空，日榜为20060102)
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	ActivityId int `json:"-"`
}
type ListRankByActivityReq struct {
	base.Req
}
type ListRankByActivityResp struct {
	RankList []SubRank `json:"rank_list"`
}

// 查询榜单数据
type RankData struct {
	Rank    int     `json:"rank"`     // 排名(-1为未上榜)
	Uid     int     `json:"uid"`      // 用户uid
	PeerUid int     `json:"peer_uid"` // 对方uid(如有)
	Score   float64 `json:"score"`    // 分数
}
type ListRankDataReq struct {
	base.Req
	RankId   int    `json:"rank_id"` // 榜单id
	Period   string `json:"period"`  // 榜单周期(总榜为空，日榜为20060102)
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
type ListRankDataResp struct {
	RankDataList []RankData `json:"rank_data_list"` // 榜单数据
	MyRankData   *RankData  `json:"my_rank_data"`   // 我的榜单
}

// 增加榜单数据
type AddRankDataReq struct {
	base.Req
	Uid     int     `json:"uid"`      // 用户uid
	PeerUid int     `json:"peer_uid"` // 对方uid(如有)
	Score   float64 `json:"score"`    // 分数
}
type AddRankDataResp struct {
}

// 查询榜单奖励配置
type RewardConfig struct {
	FromRank     int `json:"from_rank"`     // 开始排名
	ToRank       int `json:"to_rank"`       // 结束排名
	RewardType   int `json:"reward_type"`   // 奖励类型
	RewardAmount int `json:"reward_amount"` // 奖励数量
}
type ListRankRewardConfigReq struct {
	base.Req
}
type ListRankRewardConfigResp struct {
	RankRewardConfigList []RewardConfig `json:"rank_reward_config_list"` // 榜单奖励配置
}
