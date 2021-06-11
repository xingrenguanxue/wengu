package main

import (
	"wengu/routes"
	"wengu/models"
)

func main() {
	// 初始化数据库
	models.InitDB()
	// 初始化路由组件
	routes.InitRouter()
}