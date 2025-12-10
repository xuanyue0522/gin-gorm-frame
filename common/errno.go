package common

type Errno struct {
	Code   int    // 错误码
	Msg    string // 错误提示（用于在前端展示）
	ErrMsg string // 详细错误信息或原因（用于调试）
}

func (e Errno) Error() string {
	return e.Msg
}

// WithMsg 对简短错误信息进行补充说明
func (e Errno) WithMsg(msg string) Errno {
	e.Msg = e.Msg + " " + msg
	return e
}

// WithErr 错误的详细错误信息
func (e Errno) WithErr(rawErr error) Errno {
	var msg string
	if rawErr != nil {
		msg = rawErr.Error()
	}
	e.ErrMsg = msg
	return e
}

func (e Errno) IsOk() bool {
	return e.Code == 200
}

var (

	/*
		服务器级别错误
	*/
	OK              = Errno{Code: 200, Msg: "OK"}
	ServerError     = Errno{Code: 500, Msg: "Server Error"}
	ParamError      = Errno{Code: 400, Msg: "Param Error"}
	AuthError       = Errno{Code: 401, Msg: "Auth Error"}
	PermissionError = Errno{Code: 403, Msg: "Permission Error"}

	/*
		数据库级别错误
		错误码从10000开始
	*/
	DatabaseError = Errno{Code: 10000, Msg: "Database Error"}
	RedisError    = Errno{Code: 10001, Msg: "Redis Error"}

	/*
		工具级别错误
		错误码从20000开始
	*/
	InvalidCaptchaError = Errno{Code: 20000, Msg: "滑块校验失败，请刷新重试"}

	/*
		用户级别错误
		错误码从30000开始
	*/
	UserNotFound      = Errno{Code: 30000, Msg: "User Not Found"}
	UserAlreadyExists = Errno{Code: 30001, Msg: "User Already Exists"}
	UserPasswordError = Errno{Code: 30002, Msg: "User Password Error"}
	UserNotFoundError = Errno{Code: 30003, Msg: "User Not Found"}
)
