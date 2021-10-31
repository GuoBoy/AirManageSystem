package router

import (
	"AirManageSystem/handles"
	"AirManageSystem/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewManage()  error{
	addr := fmt.Sprintf("%s:%d","0.0.0.0", 23456)
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	// 后台管理
	admin := r.Group("/admin")
	{
		admin.POST("/login", handles.Login)
		admin.POST("/register", handles.Register)
		admin.GET("/list", handles.GetUser)
		admin.POST("/del", handles.DelUser)
	}

	// 航班管理
	flight := r.Group("/flight")
	{
		flight.POST("/add", handles.AddFlight)
		flight.GET("/list", handles.GetFlight)
		flight.POST("/del", handles.DelFlight)
		flight.POST("/update", handles.UpdateFlight)
	}
	// 机票管理
	ticket := r.Group("/ticket")
	{
		ticket.GET("/list", handles.GetAllTickets)
	}

	if err := r.Run(addr); err!=nil{
		fmt.Println("开启服务后台管理api失败")
		return err
	}
	return nil
}

