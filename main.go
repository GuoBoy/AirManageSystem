package main

import (
	"AirManageSystem/db"
	"AirManageSystem/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化数据库/打开数据库
	_, err := db.InitDB()
	if err != nil {
		return
	}
	sigCh := make(chan os.Signal, 1)
	// 售票服务
	go func() {
		err = router.NewFront()
		if err != nil {
			return
		}
	}()
	// 管理服务
	go func() {
		err = router.NewManage()
		if err != nil {
			return
		}
	}()
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}