package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
统一响应结构体
*/
type Response struct {
	ResCode int         `json:"res_code"` // 业务码
	Message string      `json:"msg"`      // 信息
	Success bool        `json:"success"`  // 业务是否成功
	Data    interface{} `json:"data"`     // 数据
}

/*
业务码
*/
const (
	SuccessCode = 0
	ERRORCode   = 7
)

/*
统一返回
*/
func Result(code int, data interface{}, msg string, success bool, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		ResCode: code,
		Message: msg,
		Success: success,
		Data:    data,
	})
}

/*
成功响应
*/
func Ok(c *gin.Context) {
	Result(SuccessCode, map[string]interface{}{}, "操作成功", true, c)
}

/*
可传消息的成功响应
*/
func OKWithMsg(msg string, c *gin.Context) {
	Result(SuccessCode, map[string]interface{}{}, msg, true, c)
}

/*
可传数据的成功响应
*/
func OkWithData(data interface{}, c *gin.Context) {
	Result(SuccessCode, data, "处理成功", true, c)
}

/*
失败响应
*/
func Fail(c *gin.Context) {
	Result(ERRORCode, map[string]interface{}{}, "处理失败", false, c)
}

/*
可传消息的失败响应
*/
func FailWithMsg(msg string, c *gin.Context) {
	Result(ERRORCode, map[string]interface{}{}, msg, false, c)
}

/*
可传数据的失败响应
*/
func FailWithData(data interface{}, c *gin.Context) {
	Result(ERRORCode, data, "处理失败", false, c)
}
