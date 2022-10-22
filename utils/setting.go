package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DB      string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadDatabase(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DB_HOST = file.Section("database").Key("DB_HOST").MustString("localhost")
	DB_PORT = file.Section("database").Key("DB_PORT").MustString("3306")
	DB_USER = file.Section("database").Key("DB_USER").MustString("root")
	DB_PASS = file.Section("database").Key("DB_PASS").MustString("abc123321")
	DB_NAME = file.Section("database").Key("DB_NAME").MustString("go_web")
}
