package api

import (
	"gin-gorm-frame/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	ErrMsg string `json:"err_msg"`
	Data   any    `json:"data"`
}

var resp Resp

// WriteResp 响应
func WriteResp(ctx *gin.Context, data any, errno common.Errno) {

	if errno.Code == 0 {
		errno.Code = http.StatusOK
	}

	if errno.Msg == "" {
		errno.Msg = "ok"
	}

	resp = Resp{
		Code:   errno.Code,
		Msg:    errno.Msg,
		ErrMsg: errno.ErrMsg,
		Data:   data,
	}
	ctx.JSON(http.StatusOK, resp)
}

// RespError 响应错误（一般用在中间件中）
func RespError(ctx *gin.Context, httpCode int, errno common.Errno) {

	if errno.Code == 0 {
		errno.Code = http.StatusInternalServerError
	}

	if errno.Msg == "" {
		errno.Msg = "Server Error"
	}

	resp = Resp{
		Code:   errno.Code,
		Msg:    errno.Msg,
		ErrMsg: errno.ErrMsg,
	}
	ctx.JSON(httpCode, resp)
}
