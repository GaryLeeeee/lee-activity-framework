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
	InvalidParam = ResErr{Code: 10001, Message: "参数错误"}
	SysError     = ResErr{Code: 10002, Message: "系统异常"}

	// activity
	ActivityNotExists = ResErr{Code: 20001, Message: "活动不存在"}

	// lottery

	// prop

	// rank

	// task
)

var resErrMap = map[int]ResErr{}

func init() {
	register(InvalidParam)
	register(SysError)
	register(ActivityNotExists)
}

func register(e ResErr) {
	resErrMap[e.Code] = e
}

func GetResErrMap() map[int]ResErr {
	return resErrMap
}
