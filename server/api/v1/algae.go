package v1

import (
	"github.com/SukiEva/aldb/server/model"
	"github.com/SukiEva/aldb/server/util/e"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetData
// @Summary GetData
// @Description 获取所有藻类数据
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /data [get]
func GetData(c *gin.Context) {
	code := e.CODE.Success
	res := model.GetData()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

// AddAlga
// @Summary AddAlga
// @Description 添加藻类图片数据
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/add [post]
func AddAlga(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var alga model.Alga
	err := c.ShouldBindJSON(&alga)
	if err != nil {
		code = e.CODE.AlgaBindError
	} else {
		res, err := model.AddAlga(alga)
		id := res.(primitive.ObjectID)
		if err != nil {
			code = e.CODE.DataBaseError
		} else {
			if err := model.BindToRiver(alga.River, id); err != nil {
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

// AddRiver
// @Summary AddRiver
// @Description 添加河流数据
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/add [post]
func AddRiver(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var river model.River
	err := c.ShouldBindJSON(&river)
	if err != nil {
		code = e.CODE.RiverBindError
	} else if err := model.AddRiver(river); err != nil {
		code = e.CODE.RiverAlreadyExists
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// GetRivers
// @Summary GetRivers
// @Description 获取所有河流数据
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/all [get]
func GetRivers(c *gin.Context) {
	code := e.CODE.Success
	res := model.GetRivers()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}

// SearchAlga
// @Summary SearchAlga
// @Description 搜索藻类图像
// @Tags aldb
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/search [get]
func SearchAlga(c *gin.Context) {
	code := e.CODE.Success
	key := c.Query("k")
	res := model.SearchAlga(key)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": res,
	})
}
