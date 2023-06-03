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

	var alga model.Alga
	algaForms := c.PostFormArray("algaforms")
	riverId := c.PostForm("riverId")
	preSuffix := c.PostForm("preSuffix")
	if len(riverId) <= 0 || len(preSuffix) <= 0 {
		code = e.CODE.RiverQueryError
		data["err"] = "数据集ID为空或前缀为空"
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": data,
		})
		return
	}
	for _, v := range algaForms {
		err := json.Unmarshal([]byte(v), &alga)

		if err != nil {
			code = e.CODE.AlgaBindError
			data["err"] = err.Error()
			zap.L().Error("algaFormBinder", zap.String("information", err.Error()), zap.String("content", v))
		} else {
			alga.Name = preSuffix + "_" + alga.Name
			aID, err := model.AddAlga(alga)
			if err != nil {
				code = e.CODE.DataBaseError
				data["err"] = err.Error()
			} else {
				rId, err := primitive.ObjectIDFromHex(riverId)
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
	key := c.PostForm("key")
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
