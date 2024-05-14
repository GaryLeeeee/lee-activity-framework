package code

// 自定义错误码
type ResErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ResErr) Error() string {
	return e.Message
}

var (
	Success = ResErr{Code: 0, Message: "success"}

	// 通用错误码
	InvalidParam      = ResErr{Code: 10001, Message: "参数错误"}
	SysError          = ResErr{Code: 10002, Message: "系统异常"}
	FrequentOperation = ResErr{Code: 10003, Message: "操作频繁"}

	// activity
	ActivityNotExists = ResErr{Code: 20001, Message: "活动不存在"}

	// lottery
	LotteryNotExists       = ResErr{Code: 30001, Message: "抽奖不存在"}
	LotteryTicketNotEnough = ResErr{Code: 30002, Message: "抽奖券不足"}
	UseLotteryTicketErr    = ResErr{Code: 30003, Message: "抽奖异常"}

	// prop

	// rank
	ListRankErr   = ResErr{Code: 50001, Message: "获取榜单异常"}
	RankNotExists = ResErr{Code: 50002, Message: "榜单不存在"}

	// task
	GetTaskErr = ResErr{Code: 60001, Message: "获取任务异常"}
)

var resErrMap = map[int]ResErr{}

func init() {
	register(InvalidParam)
	register(SysError)
	register(FrequentOperation)
	register(ActivityNotExists)
	register(LotteryNotExists)
	register(LotteryTicketNotEnough)
	register(UseLotteryTicketErr)
	register(ListRankErr)
	register(RankNotExists)
	register(GetTaskErr)
}

func register(e ResErr) {
	resErrMap[e.Code] = e
}

func GetResErrMap() map[int]ResErr {
	return resErrMap
}
