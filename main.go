package main

import (
	"fmt"
	"irain/ginServer/app/admin"
	"irain/ginServer/app/blog"
	"irain/ginServer/config"
	"irain/ginServer/routers"
	"irain/ginServer/util/logger"
	"log"
)

func main() {
	//初始化配置
	if err := config.Init(); err != nil {
		log.Println("load config failed:", err)
		return
	}
	
	//日志初始化
	if err := logger.InitLogger(config.LoggerConf); err != nil {
		log.Println("migrate logger failed, err:", err)
		return 
	}

	
	//路由配置
	routers.Include(admin.Routers, blog.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(fmt.Sprintf(":%v", config.ServerConf.Port)); err != nil {
		log.Println("startup service failed:", err)
	}
}

