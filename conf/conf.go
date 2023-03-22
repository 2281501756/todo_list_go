package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"todo_list/model"
)

var (
	AppMode    string
	AppPort    string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbDatabase string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic("读取配置失败")
	}
	loadService(file)
	loadMysql(file)
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbDatabase)
	model.ConnectMysql(path)
	model.Migration()
}

func loadService(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	AppPort = file.Section("service").Key("AppPort").String()
}

func loadMysql(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbDatabase = file.Section("mysql").Key("DbDatabase").String()
}
