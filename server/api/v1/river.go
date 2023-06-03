package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/model"
	"github.com/qanyue/aldb/server/util/e"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// AddRiver
// @Summary AddRiver
// @Description 添加数据集
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
				code = e.CODE.RiverBindError
				data["err"] = err.Error()
			}
		}
		data["riverId"] = rId.Hex()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}

// ShareRiver
// @Summary ShareRiver
// @Description 共享数据集
// @Tags aldb
// @Accept json
// @Produce json
// @param  userEmail body string false "用户邮箱" example(aaa)
// @param  riverId  body string false "数据集名称ID" example(test)
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/share [post]
func ShareRiver(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})
	userEmail := c.PostForm("userEmail")
	riverId := c.PostForm("riverId")
	if userEmail == "" || riverId == "" {
		code = e.CODE.UserBindError
	} else {
		id, err := primitive.ObjectIDFromHex(riverId)
		if err != nil {
			code = e.CODE.RiverQueryError
			data["err"] = err.Error()
		}
		err = model.RiverBindToUser(userEmail, id)
		if err != nil {
			code = e.CODE.RiverBindError
			data["err"] = err.Error()
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
		fmt.Println("查询所有用户邮箱" + userEmail)
		code = e.CODE.RiverQueryError
	} else {
		res := model.GetRiversWithoutAlgae(userEmail)
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

// GetRiverInfo
// @Summary GetRiverInfo
// @Description 获取所有数据集藻类列表数据
// @Tags aldb
// @Accept json
// @Produce json
// @param RiverID body string false "数据集ID" example(640fd8c403c8e6ead93ea295)
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/info [get]
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

// SearchRiver
// @Summary SearchRiver
// @Description 搜索数据集
// @Tags aldb
// @Accept json
// @Produce json
// @param Key body string false "数据集key" example("黄河")
// @Success 200 {object} object "{code, msg, data}"
// @Router /river/search [post]
func SearchRiver(c *gin.Context) {
	code := e.CODE.Success
	key := c.PostForm("key")
	userEmail := c.PostForm("userEmail")
	if len(key) <= 0 || len(userEmail) <= 0 {
		code = e.CODE.RiverQueryError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.ParseCode(code),
			"data": gin.H{
				"err": "数据集key为空",
			},
		})
	} else {
		r := model.SearchRiverByKeyWithoutAlgae(userEmail, key)
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
