package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"` //code==0成功，非0为失败
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

const (
	Success = 0
	Error   = 7
)

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "Success", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c) //map[string]any{}空的对象
}
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c) //map[string]any{}空的对象
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
