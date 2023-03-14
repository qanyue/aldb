package upload

import "mime/multipart"

type COS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	//Todo deletefile
	DeleteFile(key string) error
}

func NewCos() COS {
	return &TencentCOS{}
}

func NewHuaWeiObs() COS {
	return &HuaWeiOBS{}
}
