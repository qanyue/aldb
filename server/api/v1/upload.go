package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/util/e"
	"github.com/qanyue/aldb/server/util/upload"
	"net/http"
	"strconv"
	"time"
)

//TODO UploadFiles

func Upload(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})
	// 获取请求中的 multipart 头部
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		code = e.CODE.FileReceiveError
	} else {
		var fileUrl, fileName string
		cos := upload.NewHuaWeiObs()
		// 调用基于腾讯云官方 SDK 封装的上传方法
		fileUrl, fileName, err = cos.UploadFile(header)
		if err != nil {
			code = e.CODE.FileUploadError
		} else {
			// 返回文件名和文件链接
			data["name"] = fileName
			data["url"] = fileUrl
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}
func Uploadexamle(c *gin.Context) {
	code := e.CODE.Success
	data := make(map[string]interface{})

	data["name"] = "1679500111乌克丽丽_2014-12-27_11-16-37.jpg" + strconv.Itoa(int(time.Now().Unix()))
	data["url"] = "https://aldb.obs.cn-east-3.myhuaweicloud.com/img/1679500111乌克丽丽_2014-12-27_11-16-37.jpg"
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.ParseCode(code),
		"data": data,
	})
}
