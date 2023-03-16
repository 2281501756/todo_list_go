package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func ConnectMysql(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: false},
	})
	if err != nil {
		panic("数据库连接失败")
	}
	x, _ := db.DB()
	x.SetConnMaxIdleTime(20)
	x.SetMaxOpenConns(100)
	x.SetConnMaxLifetime(time.Second * 30)

	DB = db
}
