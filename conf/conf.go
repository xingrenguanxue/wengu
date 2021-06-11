package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	RunMode  	string
	HttpPort 	string
	Db			string
	DbHost		string
	DbPort		string
	DbUser		string
	DbPassword	string
	DbName 		string
)

// SetUp 初始化设置
func SetUp() {
	file, err := ini.Load("./conf.ini")
	if err != nil {
		fmt.Println("读取配置文件出错，请检查！")
	}
	loadServer(file)
	loadDatabase(file)
}

// 加载服务器配置
func loadServer(file *ini.File) {
	RunMode  = file.Section("server").Key("RunMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8555")
}

// 加载数据库配置
func loadDatabase(file *ini.File) {
	Db			= file.Section("database").Key("Db").MustString("mysql")
	DbHost		= file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort		= file.Section("database").Key("DbPort").MustString("3306")
	DbUser		= file.Section("database").Key("DbUser").MustString("root")
	DbPassword	= file.Section("database").Key("DbPassword").MustString("123456")
	DbName 		= file.Section("database").Key("DbName").MustString("WenGu")
}