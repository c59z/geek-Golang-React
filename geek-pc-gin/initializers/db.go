package initializers

import (
	"geek-pc-gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	log.Print("数据库用户:" + config.AppConfig.DBUser)
	dsn := config.AppConfig.DBUser + ":" + config.AppConfig.DBPass + "@tcp(127.0.0.1:3306)/" + config.AppConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
