package resp

import (
	"github.com/gin-gonic/gin"
	"irain/ginServer/app/models"
	"net/http"
)

//http成功返回消息
func Success(c *gin.Context,data interface{})  {
	c.JSON(http.StatusOK, data)
}

//http失败返回消息
func Failure(c *gin.Context, code int,msg string) {
	c.JSON(http.StatusOK,models.RespResult{
		Code: code,
		Msg:  msg,
	})
}