package admin

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	
	_ "irain/docs" // 千万不要忘了导入把你上一步生成的docs
	"irain/ginServer/app/admin/controller"
	"irain/ginServer/util/jwt"
)

func Routers(e *gin.Engine) {
	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// v1
	adminGroup := e.Group("admin")
	{
		adminGroup.POST("/register", controller.Register)
		adminGroup.POST("/login", controller.Login)
		//e.Use(jwt.AuthMiddleware())

		//以下接口需要进行jwt验证
		{
			adminGroup.POST("/logout", jwt.AuthMiddleware(), controller.Logout)

			// 待办事项
			// 添加
			adminGroup.POST("/todo", controller.CreateTodo)
			// 查看
			adminGroup.GET("/todo", controller.GetTodoList)
			// 修改
			adminGroup.PUT("/todo/:id", controller.UpdateATodo)
			// 删除
			adminGroup.DELETE("/todo/:id", controller.DeleteATodo)
		}

	}

}
