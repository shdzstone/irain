package controller

import (
	"github.com/gin-gonic/gin"
	"irain/ginServer/app/models"
	"irain/ginServer/util/jwt"
	"log"
	"net/http"
)

func Register(c *gin.Context)  {
	var admin models.Admin
	err := c.ShouldBindJSON(&admin)
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	err = models.CreateAnAdmin(&admin)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		//c.JSON(http.StatusOK, todo)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "success",
			"data": admin,
		})
	}
}

func Login(c *gin.Context)  {

	// 用户发送用户名和密码过来
	var user models.Admin
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	
	admin,err  := models.GetExistAdmin(user.Account,user.Password)
	log.Println(c.Request.URL)

	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "该账号"+user.Account+"没有注册,"+"请注册",
		})
	}

	nodes,err := models.GetAllNode(admin.ID)
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "获取节点失败",
		})
	}
	// 生成Token
	tokenString, err := jwt.GenToken(admin.Account)
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "生成token失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{
			"token": tokenString,
			"info":admin,
			"nodes":nodes,
		},
	})
	return
}

func Logout(c *gin.Context)  {
	log.Println("处理登出请求")
	account := c.MustGet("account").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"msg": "账号"+account+"已退出"},
	})
}

//func GetNodes(adminId uint) []models.Node {
//	
//}


