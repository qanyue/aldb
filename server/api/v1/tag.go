package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/model"
	"github.com/qanyue/aldb/server/util/e"
	"net/http"
)

// AddTag
// @Summary AddTag
// @Description 添加标签数据
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /tag/add [post]
func AddTag(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var tag model.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		code = e.CODE.TagBindError
	} else if err := model.AddTag(tag); err != nil {
		code = e.CODE.TagAlreadyExists
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// GetTags
// @Summary GetTags
// @Description 获取所有标签数据
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /tag/all [get]
func GetTags(c *gin.Context) {
	code := e.CODE.Success
	res := model.GetTags()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}
