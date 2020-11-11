package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"irain/ginServer/config"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GenToken 生成token
func GenToken(account string) (string, error) {
	// 创建自己的消息体
	c := config.Claims{
		Account: account, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			Issuer:    "irain",                             				//token签发人
			//IssuedAt: time.Now().Unix(),									//token签发时间
			//Subject: "token for irain",										//适用范围
			//Audience: "irain admins",										//使用者
			//NotBefore: time.Now().Unix(),									//token有效期起始时间
			ExpiresAt: time.Now().Add(config.TokenExpireDuration).Unix(), //token有效期结束时间 
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(config.SecretKey)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*config.Claims, error) {
	// 解析token
	var claims config.Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (i interface{}, err error) {
		return config.SecretKey, nil
	})
	if err != nil {
		log.Println("出错了")
		//token已过期，但未过刷新时间，刷新时间为过期时间后20秒
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired && claims.ExpiresAt > time.Now().Unix()-(120) {
			log.Println("token已过期但尚未过刷新时间，刷新时间为20秒")
			claims.ExpiresAt = time.Now().Add(config.TokenExpireDuration).Unix()
			return &claims,nil
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid { // 校验token
		return claims, nil
	}

	log.Println("invalid token")
	return nil, errors.New("invalid token")
}

// AuthMiddleware 基于JWT的认证中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 具体实现方式要依据实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		//URL HEADER解码
		decodeAuthToken, err := url.QueryUnescape(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "Authorization解码失败",
			})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(decodeAuthToken, " ", 2)
		log.Println(parts[1])

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2006,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的account信息保存到请求的上下文c上
		c.Set("account", mc.Account)
		c.Next() // 后续的处理函数可以用过c.Get("account")来获取当前请求的用户信息
	}
}
