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
	InvalidRequestParam = NewResultCode(4001, "请求参数不合法")
	InvalidRequestBody  = NewResultCode(4002, "请求数据不合法")
	InvalidCurrency     = NewResultCode(4003, "转账货币不一致")
	InsufficientBalance = NewResultCode(4004, "账户余额不足")

	IncorrectPassword = NewResultCode(4011, "密码不正确")

	AccountNotFound = NewResultCode(4041, "账户不存在")
	UserNotFound    = NewResultCode(4042, "用户不存在")

	NotUserCreateAccount = NewResultCode(4031, "只允许已注册用户创建账户")
	CreateMultiAccount   = NewResultCode(4032, "用户只允许创建不同货币的账户各一个")
	InvalidUsername      = NewResultCode(4033, "用户名已注册")
	InvalidEmail         = NewResultCode(4034, "邮箱已注册")
)
