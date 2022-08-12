package common

type ResultCode struct {
	Code int
	Msg  string
}

func NewResultCode(code int, msg string) *ResultCode {
	return &ResultCode{
		Code: code,
		Msg:  msg,
	}
}

// 前三位为对应HttpStatus状态码，在每个状态码下携带对应的code

var (
	BadRequest          = NewResultCode(4000, "错误请求")
	InvalidRequestParam = NewResultCode(4001, "请求参数不合法")
	InvalidRequestType  = NewResultCode(4002, "请求格式不合法")
	InvalidRequestBody  = NewResultCode(4003, "请求数据不合法")
	InvalidUrlDomain    = NewResultCode(4004, "无效URL")
	InvalidCurrency     = NewResultCode(4005, "转账货币不一致")
	InsufficientBalance = NewResultCode(4006, "账户余额不足")

	NotFound        = NewResultCode(4040, "无对应资源")
	AccountNotFound = NewResultCode(4041, "无此账户")
)
