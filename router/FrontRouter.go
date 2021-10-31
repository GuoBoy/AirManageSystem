package router

import (
	"AirManageSystem/handles"
	"AirManageSystem/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewFront()  error{
	addr := fmt.Sprintf("%s:%d","0.0.0.0", 12345)
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	r.POST("/login", handles.Login)
	r.POST("/register", handles.Register)

	// 前台
	front := r.Group("/v1")
	{
		front.POST("/tickets", handles.GetTickets)
		front.POST("/add/ticket", handles.AddTicket)
		front.POST("/del/ticket", handles.DelTicket)
		front.GET("/flights", handles.GetFlight)
	}
	if err := r.Run(addr); err != nil {
		fmt.Println("开启服务后台管理api失败")
		return err
	}
	return nil
}

