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

// @title irain-admin接口文档
// @version 0.1
// @description 后台管理文档

// @contact.name stone
// @contact.url http://129.211.48.125:80
// @contact.email shdzstone@gmail.com
// @termsOfService http://129.211.48.125:80

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8088
// @BasePath /admin
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

