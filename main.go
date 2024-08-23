package main

import (
	conf "payconfig/core/config"
	dbmysql "payconfig/core/db"
	"payconfig/core/logger"
	myrd "payconfig/core/redis"
	"payconfig/interfaces/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	conf.LoadConfig()

	// 日志组件
	logger.InitLogger()

	//初始化默认db
	dbmysql.NewMysql(&conf.GlobalCfg.Db)

	// 加载redis
	myrd.NewRedis(&conf.GlobalCfg.Redis)

	// 初始化路由组
	r := api.InitRouter()

	// 设置gin启动模式
	gin.SetMode(conf.GlobalCfg.Server.Mode)

	// 启动gin服务
	err := r.Run(conf.GlobalCfg.Server.Port)
	if err != nil {
		panic(err.Error())
	}
}
