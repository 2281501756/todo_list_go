package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"todo_list/global"
	"todo_list/repository/db/model"
)

func InitMysql() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       createDSN(), // data source name
		DefaultStringSize:         256,         // default size for string fields
		DisableDatetimePrecision:  true,        // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,        // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,        // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,       // a
	}))
	if err != nil {
		panic("数据库配置错误")
	}
	err = db.AutoMigrate(&model.UserModel{})
	if err != nil {
		panic("数据库表迁移失败")
	}
	global.DB = db
}

func createDSN() (res string) {
	c := global.Config.MysqlConfig
	return strings.Join([]string{c.Root, ":", c.Password, "@tcp(", c.Path, ":", c.Port, ")/", c.Name, "?charset=utf8&parseTime=True&loc=Local"}, "")
}
