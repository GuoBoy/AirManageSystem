package handles

import (
	"AirManageSystem/db"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if have, u := db.IsExist(username, password); have {
		SuccessJson(ctx, "token", u.IDCard)
	}else {
		ErrorJson(ctx)
	}
}
