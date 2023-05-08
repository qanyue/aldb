package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/model"
	"github.com/qanyue/aldb/server/util/e"
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
	alga := c.Query("algaId")
	algaId, err := primitive.ObjectIDFromHex(alga)
	if err != nil {
		code = e.CODE.AlgaQueryError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": "",
		})
	}
	res := model.GetAnnotations(algaId)
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
/*func GetAnnotationByUser(c *gin.Context) {
	code := e.CODE.Success
	user := c.Query("user")
	res := model.GetAnnotationByUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}
*/

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
		aId, err := primitive.ObjectIDFromHex(anno.AlgaId)
		if err != nil {
			code = e.CODE.AlgaQueryError
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.ParseCode(code),
				"data": "",
			})
		}
		err = model.AddAnnotation(aId, model.Annotation{
			Description: anno.Description,
			Tag:         &anno.Tag,
		})
		if err != nil {
			code = e.CODE.DataBaseError
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
	var anno model.Anno
	err := c.ShouldBindJSON(&anno)
	if err != nil {
		code = e.CODE.AnnoBindError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": "",
		})
	}
	aId, err := primitive.ObjectIDFromHex(anno.AlgaId)
	if err != nil {
		code = e.CODE.AnnoBindError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": "",
		})
	}
	if err := model.DeleteAnnotation(aId, model.Annotation{
		Description: anno.Description,
		Tag:         &anno.Tag,
	}); err != nil {
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
/*func UpdateAnnotation(c *gin.Context) {
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
*/
