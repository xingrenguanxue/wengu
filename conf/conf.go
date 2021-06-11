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

func init() {
	cfg, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Println("读取配置文件出错，请检查！")
	}
	loadServer(cfg)
	loadDatabase(cfg)
}

// 加载服务器配置
func loadServer(cfg *ini.File) {
	RunMode  = cfg.Section("server").Key("RunMode").String()
	HttpPort = cfg.Section("server").Key("HttpPort").String()
}

// 加载数据库配置
func loadDatabase(cfg *ini.File) {
	Db			= cfg.Section("database").Key("Db").String()
	DbHost		= cfg.Section("database").Key("DbHost").String()
	DbPort		= cfg.Section("database").Key("DbPort").String()
	DbUser		= cfg.Section("database").Key("DbUser").String()
	DbPassword	= cfg.Section("database").Key("DbPassword").String()
	DbName 		= cfg.Section("database").Key("DbName").String()
}