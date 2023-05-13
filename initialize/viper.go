package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"todo_list/global"
)

func InitConf() {
	vip := viper.New()
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	vip.AddConfigPath("./config")
	err := vip.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	err = vip.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("配置文件序列化失败")
	}
	fmt.Println(&global.Config)
}
