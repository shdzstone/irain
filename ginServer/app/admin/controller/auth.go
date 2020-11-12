package controller

import (
	"github.com/gin-gonic/gin"
	"irain/ginServer/app/models"
	"irain/ginServer/util/jwt"
	"irain/ginServer/util/resp"
	"log"
	"net/http"
)

// @Tags 授权相关接口
// @Description 后台注册接口
// @Summary 注册接口
// @Accept application/json
// @Produce application/json
// @Param account query string true "账号必传"
// @Param password query string true "密码必传"
// @Success 200  {string} string "{"msg": "success","code":2000}"
// @Failure 400 {string} string "{"msg": "错误信息","code":2001}"
// @Router /register [post]
func Register(c *gin.Context)  {
	var admin models.Admin
	err := c.ShouldBindJSON(&admin)
	if err!=nil {
		resp.Failure(c,2001,"无效的参数")
		return
	}
	err = models.CreateAnAdmin(&admin)
	if err != nil{
		resp.Failure(c,2001,err.Error())
	}else{
		//c.JSON(http.StatusOK, todo)
		resp.Success(c,models.RespResult{
			Code: 2000,
			Msg: "success",
			Data: admin,
		})
	}
}

// @Tags 授权相关接口
// @Description 后台管理登录接口
// @Summary 登录接口
// @Accept application/json
// @Produce application/json
// @Param account query string true "账号"
// @Param password query string true "密码必传"
// @Success 200  {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /login [post]
func Login(c *gin.Context)  {

	// GET请求参数(query string)：/admin/login
	// 初始化结构体时指定初始参数
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

	// 返回响应
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


