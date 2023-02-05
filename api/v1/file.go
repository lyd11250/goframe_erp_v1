package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/file/upload" method:"post" mime:"multipart/form-data" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type UploadFileRes struct {
	Filename string `json:"filename"`
}
