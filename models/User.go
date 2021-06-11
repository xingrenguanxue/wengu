package models

import (
	"gorm.io/gorm"
	"time"
	"wengu/utils/msg"
)

// User 用户结构体
type User struct {
	UID 	  int
	Username  string
	Password  string
	Role 	  int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt // gorm 的软删除
}

// CreateUser 创建新用户
func CreateUser(user *User) (code int) {
	err := DB.Create(user).Error
	if err != nil {
		return msg.ERROR
	}
	return msg.SUSSCE
}