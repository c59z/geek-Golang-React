package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Photo    string
}

func (User) TableName() string {
	return "userInfo" // 指定数据库名称
}

func CreateUser(db *gorm.DB, user *User) error {
	// 在数据库中创建用户
	return db.Create(user).Error
}
