package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model/pojo"
	"goframe-erp-v1/internal/service"
)

type cFile struct {
}

var File cFile

func (c *cFile) Upload(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择上传的文件")
	}
	output, err := service.File().Upload(ctx, pojo.UploadFileInput{File: req.File})
	if err != nil {
		return nil, err
	}
	res = &v1.UploadFileRes{Filename: output.Filename}
	return
}
