package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	ChannelID uint           `gorm:"not null"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Author    string         `json:"author"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除字段
}
