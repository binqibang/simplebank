package common

import "net/http"

type Response struct {
	Code int         `json:"code"`           // 错误码
	Msg  string      `json:"msg,omitempty"`  // 错误描述
	Data interface{} `json:"data,omitempty"` // 返回数据
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// Success common success result.
func Success(data interface{}) *Response {
	return &Response{
		Code: http.StatusOK,
		Msg:  "",
		Data: data,
	}
}

// Error common error result.
func Error(err error) *Response {
	return &Response{
		Code: http.StatusInternalServerError,
		Msg:  err.Error(),
		Data: nil,
	}
}

// ErrorWithCode error with specified result code.
func ErrorWithCode(code *ResultCode) *Response {
	return &Response{
		Code: code.Code,
		Msg:  code.Msg,
		Data: nil,
	}
}
