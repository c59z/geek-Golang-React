package services

import (
	"context"
	"errors"
	"fmt"
	"geek-pc-gin/initializers"
	"geek-pc-gin/models"
	"geek-pc-gin/utils"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var ctx = context.Background()

func Login(db *gorm.DB, rdb *redis.Client, username, password string) (string, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	// 使用 CheckPasswordHash 验证用户输入的密码是否匹配数据库中存储的加密密码
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("incorrect password")
	}

	// 生成 JWT Token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	// 将用户信息存储在 Redis 中
	err = rdb.Set(ctx, user.Username, token, 0).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func RegisterUser(db *gorm.DB, username, password string) error {

	// 对密码进行加密
	encryptedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := models.User{
		Username: username,
		Password: encryptedPassword,
	}

	return models.CreateUser(db, &newUser)
}

func GetUserInfoByID(userId int) models.User {
	var user models.User
	if err := initializers.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		fmt.Println("未找到用户")
		return models.User{}
	}
	return user
}
