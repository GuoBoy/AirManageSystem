package handles

import (
	"AirManageSystem/db"
	"github.com/gin-gonic/gin"
)

func AddFlight(ctx *gin.Context)  {
	var f db.FlightInfo
	if err := ctx.ShouldBindJSON(&f);err!=nil{
		ErrorJson(ctx)
		return
	}
	f.AddFlight()
	SuccessJson(ctx, "msg", "ok")
}

func DelFlight(ctx *gin.Context)  {
	ID := ctx.PostForm("ID")
	db.DelFlight(ID)
	SuccessJson(ctx, "msg", "ok")
}

func UpdateFlight(ctx *gin.Context)  {
	var flight db.FlightInfo
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		ErrorJson(ctx)
		return
	}
	flight.UpdateFlightInfo()
	SuccessJson(ctx, "msg", "ok")
}

func GetFlight(ctx *gin.Context)  {
	data, length := db.GetAllFlight()
	SuccessDataList(ctx, data, length)
}