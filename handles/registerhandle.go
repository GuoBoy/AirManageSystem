package handles

import (
	"AirManageSystem/db"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context)  {
	var u db.User
	if err:= ctx.ShouldBindJSON(&u); err!=nil{
		// 解析请求失败
		ErrorJson(ctx)
		return
	}
	u.AddUser()
	SuccessJson(ctx, "token", u.IDCard)
}
