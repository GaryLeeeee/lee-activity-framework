package service

import (
	"lee-activity-framework/lee-activity-api/activity"
	"lee-activity-framework/lee-activity-api/code"
	"lee-activity-framework/lee-activity-api/rank"
	"lee-activity-framework/lee-activity-rank/dao"
	"lee-activity-framework/lee-activity-rank/model"
	"lee-activity-framework/lee-activity-rank/strategy"
)

func ListRankByActivity(req rank.ListRankByActivityReq) (*rank.ListRankByActivityResp, error) {
	// 1.判断活动是否存在
	isRunning, err := activity.IsActivityRunning(req.ActivityId)
	if err != nil {
		return nil, err
	}
	if !isRunning {
		return nil, code.ActivityNotExists
	}

	// 2.查询活动对应的榜单配置
	rankList, err := dao.ListRankByActivityId(req.ActivityId)
	if err != nil {
		return nil, code.ListRankErr
	}
	if len(rankList) <= 0 {
		return &rank.ListRankByActivityResp{}, nil
	}

	return &rank.ListRankByActivityResp{
		RankList: listRankByRankType(rankList),
	}, nil
}

func ListRankData(req rank.ListRankDataReq) (*rank.ListRankDataResp, error) {
	// 1.判断榜单是否存在
	rankListByActivity, err := ListRankByActivity(rank.ListRankByActivityReq{
		Req: req.Req,
	})
	if err != nil {
		return nil, err
	}
	var subRank *rank.SubRank
	for _, subRankItem := range rankListByActivity.RankList {
		if subRankItem.Id == req.RankId && subRankItem.Period == req.Period {
			subRank = &subRankItem
			break
		}
	}
	if subRank == nil {
		return nil, code.RankNotExists
	}

	// 2.查询榜单数据
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	rankDataList, err := dao.ListRankData(subRank.Id, subRank.Period, req.Page, req.PageSize)
	if err != nil {
		return nil, code.ListRankErr
	}

	// 3.查询我的榜单数据
	myRankData, err := dao.GetMyRankData(req.Uid, subRank.Id, subRank.Period)
	if err != nil {
		return nil, code.ListRankErr
	}

	// 4.填充数据
	resp := &rank.ListRankDataResp{
		RankDataList: make([]rank.RankData, 0),
	}
	for i, item := range rankDataList {
		rankData := rank.RankData{
			Rank:    (req.Page-1)*req.PageSize + i + 1,
			Uid:     item.Uid,
			PeerUid: item.PeerUid,
			Score:   item.Score,
		}
		resp.RankDataList = append(resp.RankDataList, rankData)
		if item.Uid == req.Uid {
			resp.MyRankData = &rankData
		}
	}

	// 5.兼容用户未上榜
	if resp.MyRankData == nil && myRankData != nil {
		resp.MyRankData = &rank.RankData{
			Rank:    -1,
			Uid:     myRankData.Uid,
			PeerUid: myRankData.PeerUid,
			Score:   myRankData.Score,
		}
	}
	return resp, nil
}

func listRankByRankType(rankList []*model.LeeActivityRankConfig) []rank.SubRank {
	subRankList := make([]rank.SubRank, 0)
	if len(rankList) == 0 {
		return subRankList
	}
	for _, rankConfig := range rankList {
		subRankList = append(subRankList, strategy.GetRankStrategy(rankConfig.Type).Convert2SubRank(rankConfig)...)
	}
	return subRankList
}
