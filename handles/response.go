package handles

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessJson 成功请求json响应
func SuccessJson(ctx *gin.Context, name string, data interface{})  {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		name: data,
	})
}
// ErrorJson 错误请求json响应
func ErrorJson(ctx *gin.Context)  {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
	})
}

// SuccessDataList 响应列表数据
func SuccessDataList(ctx *gin.Context,  data, total interface{})  {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"total": total,
	})
}
