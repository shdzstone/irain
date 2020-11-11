package routers

import (
	"github.com/gin-gonic/gin"
	"irain/ginServer/config"
	"irain/ginServer/util/logger"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	//设置gin的模式
	switch config.ServerConf.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)	
	}
	
	e := gin.Default()
	
	// 注册zap相关中间件
	e.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 告诉gin框架模板文件引用的静态文件去哪里找
	e.Static("/static", "static")
	
	for _, opt := range options {
		opt(e)
	}
	return e
}

