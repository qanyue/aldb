package upload

import (
	"context"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/qanyue/aldb/server/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

var cfg = config.Conf.TencentCOS

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + cfg.Bucket + ".cos." + cfg.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretID,
			SecretKey: cfg.SecretKey,
		},
	})
	return client
}

func NewOBSClient() *obs.ObsClient {
	endpoint, _ := url.Parse("https://" + cfg.Bucket + ".obs." + cfg.Region + ".myhuaweicloud.com")
	cilent, err := obs.New(cfg.SecretID, cfg.SecretKey, endpoint.String())
	if err != nil {
		zap.L().Error("New obs cilent error ",
			zap.String("message", err.Error()))
	}
	return cilent

}

type TencentCOS struct{}
type HuaWeiOBS struct{}

// UploadFile to OBS
func (h HuaWeiOBS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewOBSClient()
	f, openError := file.Open()
	if openError != nil {
		// log
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	input := &obs.PutObjectInput{}
	input.Bucket = cfg.Bucket
	input.Key = cfg.PathPrefix + "/" + fileKey
	input.Body = f
	_, err := client.PutObject(input)
	if err != nil {
		if obsError, ok := err.(obs.ObsError); ok {
			zap.L().Error("Upload to Obs error",
				zap.String("message", obsError.Message))
		}
	}
	return cfg.BaseURL + "/" + cfg.PathPrefix + "/" + fileKey, fileKey, err
}

func (h HuaWeiOBS) DeleteFile(key string) error {
	client := NewOBSClient()
	name := cfg.PathPrefix + "/" + key
	input := &obs.DeleteObjectInput{}
	input.Bucket = cfg.Bucket
	input.Key = name
	_, err := client.DeleteObject(input)
	if err != nil {
		// log
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		// log
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), cfg.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return cfg.BaseURL + "/" + cfg.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := cfg.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		// log
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}
