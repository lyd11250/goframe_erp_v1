package pojo

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileInput struct {
	File *ghttp.UploadFile
}

type UploadFileOutput struct {
	Filename string
}
