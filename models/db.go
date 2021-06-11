package models

// 使用gorm 详见https://github.com/go-gorm/gorm/

import (
	"fmt"
	"wengu/conf"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	// DB 用于SQL操作
	DB *gorm.DB
	err error
)

// InitDB 初始化数据库
func InitDB() {
	// dsm example: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
						conf.DbUser,
						conf.DbPassword,
						conf.DbHost,
						conf.DbPort,
						conf.DbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println("数据库连接失败，请检查配置后重新尝试！")
	}
}
