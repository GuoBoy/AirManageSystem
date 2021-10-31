package handles

import (
	"AirManageSystem/db"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context)  {
	data, length := db.GetAllUser()
	SuccessDataList(ctx, data, length)
}

func DelUser(ctx *gin.Context)  {
	ID := ctx.PostForm("ID")
	db.DelUser(ID)
	SuccessJson(ctx, "msg", "ok")
}
