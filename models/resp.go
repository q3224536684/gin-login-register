package models

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	RequestId string      `json:"requestId"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

func NewResp(requestId string, code int, msg string, data interface{}) *Resp {
	return &Resp{RequestId: requestId, Code: code, Msg: msg, Data: data}
}

// Result 响应
func Result(code int, msg string, data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Resp{
		RequestId: requestid.Get(ctx),
		Code:      code,
		Msg:       msg,
		Data:      data,
	})
}

// Success 成功返回
func Success(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Resp{
		RequestId: requestid.Get(ctx),
		Code:      200,
		Msg:       "SUCCESS",
		Data:      data,
	})
}

// Fail 失败返回
func Fail(code int, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Resp{
		RequestId: requestid.Get(ctx),
		Code:      code,
		Msg:       msg,
		Data:      nil,
	})
}
