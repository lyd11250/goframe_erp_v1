package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/service"
)

type sUser struct {
}

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

func (s *sUser) GetUserById(ctx context.Context, in model.GetUserByIdInput) (out model.GetUserByIdOutput, err error) {
	user := entity.SysUser{}
	err = dao.SysUser.Ctx(ctx).WherePri(in.UserId).Scan(&user)
	if err != nil {
		return
	}
	out.UserInfo = convertDbEntityToOutput(user)
	return
}

func (s *sUser) GetUserByUserName(ctx context.Context, in model.GetUserByUserNameInput) (out model.GetUserByUserNameOutput, err error) {
	user := entity.SysUser{}
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserName, in.UserName).Scan(&user)
	if err != nil {
		return
	}
	out.UserInfo = convertDbEntityToOutput(user)
	return
}

func (s *sUser) UserLogin(ctx context.Context, in model.UserLoginInput) (out model.UserLoginOutput, err error) {
	user := entity.SysUser{}
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().UserName:     in.UserName,
		dao.SysUser.Columns().UserPassword: encryptPassword(in.UserPassword),
	}).Scan(&user)
	if err != nil {
		return
	}
	out.UserInfo = convertDbEntityToOutput(user)
	return
}

func (s *sUser) UpdateUserById(ctx context.Context, in model.UpdateUserByIdInput) (err error) {
	data := g.Map{
		dao.SysUser.Columns().UserPhone:    in.UserPhone,
		dao.SysUser.Columns().UserRealName: in.UserRealName,
		dao.SysUser.Columns().UserImage:    in.UserImage,
		dao.SysUser.Columns().UserStatus:   in.UserStatus,
	}
	if in.UserPassword != nil {
		data[dao.SysUser.Columns().UserPassword] = encryptPassword(*in.UserPassword)
	}
	_, err = dao.SysUser.Ctx(ctx).OmitEmpty().Data(data).WherePri(in.UserId).Update()
	return
}

func (s *sUser) AddUser(ctx context.Context, in model.AddUserInput) (out model.AddUserOutput, err error) {
	// 检查用户名是否存在
	count, err := dao.SysUser.Ctx(ctx).Count(dao.SysUser.Columns().UserName, in.UserName)
	if err != nil {
		return model.AddUserOutput{}, err
	}
	if count > 0 {
		return model.AddUserOutput{}, gerror.Newf("用户名%s已存在", in.UserName)
	}

	// 输入对象转换为DB对象
	user := g.Map{
		dao.SysUser.Columns().UserName:     in.UserName,
		dao.SysUser.Columns().UserPassword: encryptPassword(in.UserPassword),
		dao.SysUser.Columns().UserRealName: in.UserRealName,
		dao.SysUser.Columns().UserPhone:    in.UserPhone,
		dao.SysUser.Columns().UserImage:    in.UserImage,
	}

	// 插入并返回自动生成的ID
	id, err := dao.SysUser.Ctx(ctx).InsertAndGetId(user)
	if err != nil {
		return model.AddUserOutput{}, err
	}
	out = model.AddUserOutput{UserId: id}
	return
}

func encryptPassword(password string) string {
	return gmd5.MustEncrypt(password)
}

func convertDbEntityToOutput(entity entity.SysUser) (out model.UserInfo) {
	return model.UserInfo{
		UserId:       entity.UserId,
		UserName:     entity.UserName,
		UserRealName: entity.UserRealName,
		UserPhone:    entity.UserPhone,
		UserImage:    entity.UserImage,
		UserStatus:   entity.UserStatus,
	}
}
