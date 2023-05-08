package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/model"
	"github.com/qanyue/aldb/server/util/e"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

// GetRiverInfo
// @Summary GetRiverInfo
// @Description 获取所有数据集藻类列表数据
// @Tags aldb
// @Accept json
// @Produce json
// @param RiverID body string false "数据集ID" example(640fd8c403c8e6ead93ea295)
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/info [get]
// TODO 修改前端使用逻辑为获取数据集图片而不是所有图片
func GetRiverInfo(c *gin.Context) {
	code := e.CODE.Success
	rId := c.Query("riverId")
	if len(rId) <= 0 {
		code = e.CODE.RiverQueryError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": gin.H{
				"err": "数据集ID为空",
			},
		})

	} else {

		r := model.GetRiverByID(rId)
		if r == nil {
			code = e.CODE.RiverQueryError
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": r,
		})
	}
}

// AddAlga
// @Summary AddAlga
// @Description 添加藻类图片数据
// @Tags aldb
// @Accept json
// @Produce json
// @param  riverId body string false "数据集ID" example(640fd8c403c8e6ead93ea295)
// @param  name body string false "图片ID" example(64117bf11c55d4d364da3828)
// @param  src body string false "图片连接" example(https://aldb.obs.cn-east-3.myhuaweicloud.com/img/1679155481苹果_2014-12-27_11-16-14.jpg)
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/add [post]
func AddAlga(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var algaUpload model.AlgaUpload
	err := c.ShouldBindJSON(&algaUpload)
	if err != nil {
		code = e.CODE.AlgaBindError
		data["err"] = err.Error()
	} else {
		aID, err := model.AddAlga(algaUpload.Alga)
		if err != nil {
			code = e.CODE.DataBaseError
			data["err"] = err.Error()
		} else {
			rId, err := primitive.ObjectIDFromHex(algaUpload.RiverId)
			if err != nil {
				code = e.CODE.DataBaseError
				data["err"] = err.Error()
			} else {
				if err := model.BindToRiver(rId, aID); err != nil {
					code = e.CODE.DataBindError
					data["err"] = err.Error()
				}

			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// GetAllAlgas
// @Summary getAllAlgas
// @Description 获取一个数据集所有图片
// @Tags aldbs
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/all [Get]
func GetAllAlgas(c *gin.Context) {
	code := e.CODE.Success
	rId := c.Query("riverId")
	if len(rId) <= 0 {
		code = e.CODE.RiverQueryError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": gin.H{
				"err": "数据集ID为空",
			},
		})

	} else {
		r := model.GetAllAlga(rId)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": r,
		})
	}

}

// AddAlgas
// @Summary AddAlgas
// @Description 批量添加藻类图片数据
// @Tags aldbs
// @Accept json
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/addMul [post]
func AddAlgas(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var algaUpload model.AlgaUpload
	algaForms := c.PostFormArray("algaforms")

	for _, v := range algaForms {
		//fmt.Println(v)
		err := json.Unmarshal([]byte(v), &algaUpload)
		if err != nil {
			code = e.CODE.AlgaBindError
			data["err"] = err.Error()
			zap.L().Error("algaFormBinder", zap.String("information", err.Error()), zap.String("content", v))
		} else {
			aID, err := model.AddAlga(algaUpload.Alga)
			if err != nil {
				code = e.CODE.DataBaseError
				data["err"] = err.Error()
			} else {
				rId, err := primitive.ObjectIDFromHex(algaUpload.RiverId)
				if err != nil {
					code = e.CODE.DataBaseError
					data["err"] = err.Error()
				} else {

					if err := model.BindToRiver(rId, aID); err != nil {
						code = e.CODE.DataBindError
						data["err"] = err.Error()
					}

				}
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
// @param  userEmail body string false "用户邮箱" example(aaa)
// @param  name body string false "数据集名称ID" example(test)
// @param  description body string false "数据集描述" example(testing)
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/add [post]
func AddRiver(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	var riverAdd model.RiverAdd
	err := c.ShouldBindJSON(&riverAdd)
	if err != nil {
		code = e.CODE.RiverBindError
		data["err"] = err.Error()
	} else {
		rId, err := model.AddRiver(riverAdd.River)
		if err != nil {
			code = e.CODE.RiverAlreadyExists
			data["err"] = err.Error()
		} else {
			err = model.RiverBindToUser(riverAdd.UserEmail, rId)
			if err != nil {
				code = e.CODE.RiverAlreadyExists
				data["err"] = err.Error()
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// GetRivers
// @Summary GetRivers
// @Description 获取用户所有数据集
// @Tags aldb
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/all [post]
func GetRivers(c *gin.Context) {
	code := e.CODE.Success
	userEmail := c.PostForm("userEmail")
	if len(userEmail) <= 0 {
		code = e.CODE.RiverQueryError
	} else {
		res := model.GetRivers(userEmail)
		if len(res) >= 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.ParseCode(code),
				"data": res,
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": gin.H{
			"err": "查询用户所有数据集错误",
		},
	})
}

// SearchAlga
// @Summary SearchAlga
// @Description 搜索藻类图像
// @Tags aldb
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/search [post]
func SearchAlga(c *gin.Context) {
	code := e.CODE.Success
	key := c.PostForm("k")
	riverId := c.PostForm("riverId")
	rId, err := primitive.ObjectIDFromHex(riverId)
	if err != nil {
		code = e.CODE.RiverQueryError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": "",
		})
	} else {
		res := model.SearchAlga(rId, key)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": res,
		})
	}
}

// GetAlga
// @Summary getAlga
// @Description 得到具体藻类图像数据
// @Tags aldb
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} object "{code, msg, data}"
// @Router /alga/get [post]
func GetAlga(c *gin.Context) {
	code := e.CODE.Success
	algaId := c.PostForm("algaId")
	aId, err := primitive.ObjectIDFromHex(algaId)
	if err != nil {
		code = e.CODE.AlgaQueryError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": "",
		})
	} else {
		res := model.GetAlgaData(aId)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": *res,
		})
	}
}
