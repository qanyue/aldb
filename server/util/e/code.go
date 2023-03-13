package e

type code struct {
	Success            int
	CaptchaError       int
	CaptchaMismatch    int
	AuthBindError      int
	AuthEmpty          int
	AuthFormatError    int
	AuthTokenError     int
	AuthTokenTimeOut   int
	AuthError          int
	AuthAccessError    int
	UserBindError      int
	UserAlreadyExists  int
	RiverBindError     int
	RiverAlreadyExists int
	FileUploadError    int
	FileReceiveError   int
	AlgaBindError      int
	AnnoBindError      int
	DataBaseError      int
	DataBindError      int
}

var CODE code

func init() {
	CODE = code{
		Success:            200,
		CaptchaError:       301,
		CaptchaMismatch:    302,
		AuthBindError:      400,
		AuthEmpty:          401,
		AuthFormatError:    402,
		AuthTokenError:     403,
		AuthTokenTimeOut:   404,
		AuthError:          405,
		AuthAccessError:    406,
		UserBindError:      500,
		UserAlreadyExists:  501,
		RiverBindError:     601,
		RiverAlreadyExists: 602,
		FileReceiveError:   701,
		FileUploadError:    702,
		AlgaBindError:      800,
		AnnoBindError:      801,
		DataBaseError:      900,
		DataBindError:      901,
	}
}

func ParseCode(num int) string {
	var msg = ""
	switch num {
	case CODE.Success:
		msg = "OK"
	case CODE.CaptchaError:
		msg = "无法获取验证码"
	case CODE.CaptchaMismatch:
		msg = "验证码错误"
	case CODE.AuthBindError:
		msg = "授权格式错误"
	case CODE.AuthEmpty:
		msg = "未授权"
	case CODE.AuthFormatError:
		msg = "授权Token格式错误"
	case CODE.AuthTokenError:
		msg = "授权错误"
	case CODE.AuthTokenTimeOut:
		msg = "授权超时，请重新登录"
	case CODE.AuthError:
		msg = "密码错误"
	case CODE.AuthAccessError:
		msg = "权限不足"
	case CODE.UserBindError:
		msg = "用户格式错误"
	case CODE.UserAlreadyExists:
		msg = "用户已被注册"
	case CODE.RiverBindError:
		msg = "河流格式错误"
	case CODE.RiverAlreadyExists:
		msg = "河流已经存在"
	case CODE.FileReceiveError:
		msg = "接收文件失败"
	case CODE.FileUploadError:
		msg = "文件上传失败"
	case CODE.AlgaBindError:
		msg = "图像提交数据错误"
	case CODE.AnnoBindError:
		msg = "标注提交数据错误"
	case CODE.DataBaseError:
		msg = "数据库错误"
	case CODE.DataBindError:
		msg = "数据绑定错误"
	default:
		msg = "未知错误"
	}
	return msg
}
