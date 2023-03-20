package main

import (
	"gin_vue_project/core"
	"gin_vue_project/global"
	"gin_vue_project/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在：%s", addr)
	router.Run(addr)
}
