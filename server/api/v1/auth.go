package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/captcha"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/SukiEva/aldb/server/util/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	CaptchaId    string `json:"captchaId" binding:"required"`
	CaptchaValue string `json:"captchaValue" binding:"required"`
}

func GetAuth(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var a auth
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.CODE.AuthBindError,
			"msg":  e.ParseCode(code),
			"data": data,
		})
		return
	}
	// 判断验证码是否正确
	ok := captcha.Verify(a.CaptchaId, a.CaptchaValue)
	// 判断账号密码是否正确
	if ok && model.CheckAuth(a.Email, a.Password) {
		// 生成令牌 token
		token, err := jwt.GenToken(a.Email, a.Password)
		if err != nil {
			code = e.CODE.AuthTokenError
		} else {
			// token 返回
			data["token"] = token
		}
	} else if !ok {
		code = e.CODE.CaptchaMismatch
	} else {
		code = e.CODE.AuthError
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

func GetCaptcha(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})
	id, b64s, err := captcha.Generate()
	if err != nil {
		code = e.CODE.CaptchaError
	} else {
		data["captchaId"] = id
		data["captchaValue"] = b64s
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}
