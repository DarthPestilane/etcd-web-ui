package bizio

import (
	"github.com/gin-gonic/gin"
)

var _ error = &HttpError{}

// HttpError http 错误; 实现了 error 接口
type HttpError struct {
	HttpCode int    // http 响应码
	Code     int    // 内部错误码
	Msg      string // 错误信息
}

func (be *HttpError) Error() string {
	return be.Msg
}

// NewHttpErr 创建一个 http 错误对象
func NewHttpErr(httpCode int, err error, errCode ...int) *HttpError {
	var code = -httpCode
	if len(errCode) != 0 {
		code = errCode[0]
	}
	return &HttpError{
		HttpCode: httpCode,
		Code:     code,
		Msg:      err.Error(),
	}
}

// HandleHttpErr 处理 http 错误，发出对应的 response.
func HandleHttpErr(c *gin.Context, err error) {
	if e, ok := err.(*HttpError); ok {
		HttpFail(c, e.HttpCode, e.Msg, e.Code)
		return
	}
	HttpFail(c, 500, "Server Error")
}

// HttpResponse http 响应结构体
type HttpResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	ErrMsg string      `json:"err_msg"`
}

// HttpOK 发送 http 成功响应
func HttpOK(c *gin.Context, data interface{}) {
	c.JSON(200, &HttpResponse{
		Code:   0,
		Data:   data,
		ErrMsg: "",
	})
}

// HttpFail 发送 http 失败响应
func HttpFail(c *gin.Context, httpCode int, msg string, errCode ...int) {
	var code = httpCode * 10
	if len(errCode) != 0 {
		code = errCode[0]
	}
	c.JSON(httpCode, &HttpResponse{
		Code:   code,
		Data:   nil,
		ErrMsg: msg,
	})
}
