package models

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

func (Channel) TableName() string {
	return "channels" // 指定数据库名称
}

//func CreateUser(db *gorm.DB, user *User) error {
//	// 在数据库中创建用户
//	return db.Create(user).Error
//}
