package file

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/service"
)

type sFile struct {
}

func New() *sFile {
	return &sFile{}
}

func init() {
	service.RegisterFile(New())
}

func (s *sFile) Upload(ctx context.Context, input model.UploadFileInput) (output model.UploadFileOutput, err error) {
	// 储存文件到临时目录
	tempPath, err := g.Cfg().Get(ctx, "app.temp")
	if err != nil {
		return model.UploadFileOutput{}, err
	}
	filename, err := input.File.Save(tempPath.String(), true)
	if err != nil {
		return model.UploadFileOutput{}, err
	}
	// 获取文件md5值
	md5 := gmd5.MustEncryptFile(tempPath.String() + filename)
	// 判断该文件是否已存在
	count, err := dao.SysFile.Ctx(ctx).WherePri(md5).Count()
	if err != nil {
		return model.UploadFileOutput{}, err
	}
	if count > 0 {
		// 文件已存在，直接返回文件名，删除临时目录中的文件
		err = gfile.Remove(tempPath.String() + filename)
		if err != nil {
			return model.UploadFileOutput{}, err
		}
		var result entity.SysFile
		err = dao.SysFile.Ctx(ctx).WherePri(md5).Scan(&result)
		if err != nil {
			return model.UploadFileOutput{}, err
		}
		output = model.UploadFileOutput{Filename: result.FileName}
		return
	}
	// 获取文件存储路径配置
	path, err := g.Cfg().Get(ctx, "app.path")
	if err != nil {
		return model.UploadFileOutput{}, err
	}
	// 移动临时文件到主目录
	err = gfile.Move(tempPath.String()+filename, path.String()+filename)
	if err != nil {
		return model.UploadFileOutput{}, err
	}
	// 储存数据到数据库
	_, err = dao.SysFile.Ctx(ctx).Save(entity.SysFile{
		FileMd5:  md5,
		FileName: filename,
	})
	if err != nil {
		return model.UploadFileOutput{}, err
	}
	output = model.UploadFileOutput{Filename: filename}
	return
}
