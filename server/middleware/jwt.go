package middleware

import (
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/SukiEva/aldb/server/util/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var claims *jwt.Claims
		code := e.CODE.Success
		// 从请求头部字段中获取 token
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			code = e.CODE.AuthEmpty
		} else {
			// 分割 token，获得 jwt 的三部分 header、payload、signature
			parts := strings.Split(authHeader, ".")
			if len(parts) != 3 { // token 格式正确
				code = e.CODE.AuthFormatError
			} else {
				var err error
				// 验证 token 是否存在
				claims, err = jwt.ParseToken(authHeader)
				if err != nil {
					code = e.CODE.AuthTokenError
				} else if time.Now().Unix() > claims.ExpiresAt {
					// token 过了有效期
					code = e.CODE.AuthTokenTimeOut
				}
			}
		}
		if code != e.CODE.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.ParseCode(code),
			})
			c.Abort()
			return
		}
		// 将当前请求信息保存到请求的上下文
		c.Set("UserEmail", claims.UserEmail)
		c.Set("UserPwd", claims.UserPwd)
		c.Next() //后续处理函数 c.Get("UserEmail") 获取当前请求的用户信息
	}
}
