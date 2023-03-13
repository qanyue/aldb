package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var user model.Operator
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = e.CODE.UserBindError
	} else if err := model.AddUser(user); err != nil {
		code = e.CODE.UserAlreadyExists
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// UpdateUser
// @Summary UpdateUser
// @Description 修改用户信息
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /user/update [post]
func UpdateUser(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var user model.Operator
	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = e.CODE.UserBindError
	} else if err := model.UpdateUser(user); err != nil {
		code = e.CODE.DataBaseError
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// GetUser
// @Summary GetUser
// @Description 获取用户信息
// @Tags aldb
// @Accept json
// @Produce json
// @Param email query string true "用户邮箱"
// @Success 200 {object} object "{code, msg, data}"
// @Router /user/info [get]
func GetUser(c *gin.Context) {
	code := e.CODE.Success
	email := c.Query("email")
	res := model.GetUser(email)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

// GetUsers
// @Summary GetUsers
// @Description 获取所有用户信息
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /user/all [get]
func GetUsers(c *gin.Context) {
	res, _ := c.Get("UserEmail")
	email := res.(string)
	if !model.CheckAdmin(email) {
		code := e.CODE.AuthAccessError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": nil,
		})
	} else {
		code := e.CODE.Success
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": model.GetUsers(),
		})
	}
}

// DeleteUser
// @Summary DeleteUser
// @Description 删除用户
// @Tags aldb
// @Accept json
// @Produce json
// @Param email query string true "用户邮箱"
// @Success 200 {object} object "{code, msg, data}"
// @Router /user/delete [get]
func DeleteUser(c *gin.Context) {
	res, _ := c.Get("UserEmail")
	if !model.CheckAdmin(res.(string)) {
		code := e.CODE.AuthAccessError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": nil,
		})
	} else {
		code := e.CODE.Success
		email := c.Query("email")
		if err := model.DeleteUser(email); err != nil {
			code = e.CODE.DataBaseError
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": nil,
		})
	}
}

// ChangePassword
// @Summary ChangePassword
// @Description 修改用户密码
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /user/pwd [post]
func ChangePassword(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var ch change
	err := c.ShouldBindJSON(&ch)
	if err != nil {
		code = e.CODE.UserBindError
	} else {
		if model.CheckAuth(ch.Email, ch.Password) {
			if err := model.ChangePassword(ch.Email, ch.NewPassword); err != nil {
				code = e.CODE.DataBaseError
			}
		} else {
			code = e.CODE.AuthError
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

type change struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}
