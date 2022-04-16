package inits

import (
	"fmt"

	"github.com/code4EE/cloud-sql/config"
	"github.com/code4EE/cloud-sql/mysql"
	"github.com/code4EE/cloud-sql/server"
)

func InitAll() {
	// 初始化config
	if err := config.InitConfig(""); err != nil {
		panic(fmt.Sprintf("config init failed:%v", err))
	}
	// 初始化Mysql
	if err := mysql.InitMysql(config.AppCfg.DBAddr); err != nil {
		panic(fmt.Sprintf("mysql init failed:%v", err))
	}
	// 初始化Server
	if err := server.InitServer(config.AppCfg.ServerAddr); err != nil {
		panic(fmt.Sprintf("server init failed:%v", err))
	}
}
