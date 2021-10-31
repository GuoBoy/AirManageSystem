package handles

import (
	"AirManageSystem/db"
	"github.com/gin-gonic/gin"
)

func AddTicket(ctx *gin.Context)  {
	var t db.TicketInfo
	if err:= ctx.BindJSON(&t); err!=nil || t.UserIDCard=="" || t.FlightFid == 0{
		ErrorJson(ctx)
		return
	}
	t.AddTicket()
	SuccessJson(ctx, "msg", "success")
}

func DelTicket(ctx *gin.Context)  {
	ID := ctx.PostForm("ID")
	db.DelTicket(ID)
	SuccessJson(ctx, "msg", "success")
}

// GetTickets 获取买票信息
func GetTickets(ctx *gin.Context)  {
	token := ctx.PostForm("token")
	if token == ""{
		ErrorJson(ctx)
		return
	}
	data, length := db.GetTicketByUid(token)
	SuccessDataList(ctx, data, length)
}

// GetAllTickets 获取所有人员所有机票列表
func GetAllTickets(ctx *gin.Context)  {
	data, length := db.GetTickets()
	SuccessDataList(ctx, data, length)
}