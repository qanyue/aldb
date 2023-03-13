package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetAnnotationByAlga
// @Summary GetAnnotationByAlga
// @Description 藻类图像获取标注
// @Tags aldb
// @Accept json
// @Produce json
// @Param alga query string true "藻类名称"
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/anno [get]
func GetAnnotationByAlga(c *gin.Context) {
	code := e.CODE.Success
	alga := c.Query("alga")
	res := model.GetAnnotationByAlga(alga)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

// GetAnnotationByUser
// @Summary GetAnnotationByUser
// @Description 用户获取个人标注
// @Tags aldb
// @Accept json
// @Produce json
// @Param user query string true "用户邮箱"
// @Success 200 {object} object "{code, msg, data}"
// @Router /user/anno [get]
func GetAnnotationByUser(c *gin.Context) {
	code := e.CODE.Success
	user := c.Query("user")
	res := model.GetAnnotationByUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

// AddAnnotation
// @Summary AddAnnotation
// @Description 添加标注
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /anno/add [post]
func AddAnnotation(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var anno model.Anno
	err := c.ShouldBindJSON(&anno)
	if err != nil {
		code = e.CODE.AnnoBindError
	} else {
		// 添加藻类标注到数据库
		res, err := model.AddAnnotation(anno)
		// 从结果中断言生成的唯一 id
		id := res.(primitive.ObjectID)
		if err != nil {
			code = e.CODE.DataBaseError
		} else {
			// 绑定到藻类图像
			err1 := model.BindToAlga(anno.Alga, id)
			// 绑定到用户
			err2 := model.BindToUser(anno.User, id)
			if err1 != nil || err2 != nil {
				code = e.CODE.DataBindError
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// DeleteAnnotation
// @Summary DeleteAnnotation
// @Description 删除标注
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /anno/delete [get]
func DeleteAnnotation(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})
	id := c.Query("id")
	if err := model.DeleteAnnotation(id); err != nil {
		code = e.CODE.DataBaseError
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// UpdateAnnotation
// @Summary UpdateAnnotation
// @Description 修改标注信息
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /anno/update [post]
func UpdateAnnotation(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var anno model.Annotation
	err := c.ShouldBindJSON(&anno)
	if err != nil {
		code = e.CODE.AnnoBindError
	} else if err := model.UpdateAnnotation(anno); err != nil {
		code = e.CODE.DataBaseError
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}
