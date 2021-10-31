package db
import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./assets/db.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 迁移
	if err = db.AutoMigrate(&User{}, &FlightInfo{}, &TicketInfo{}); err != nil {
		fmt.Println("迁移scheme失败", err.Error())
		return nil, err
	}
	DB = db
	fmt.Println("初始化数据库成功")
	return DB, nil
}
