package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/consts"
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

func (s *sUser) GetUserList(ctx context.Context) (out model.GetUserListOutput, err error) {
	err = dao.SysUser.Ctx(ctx).OrderDesc(dao.SysUser.Columns().UserStatus).Scan(&out.List)
	return
}

func (s *sUser) UserLogin(ctx context.Context, in model.UserLoginInput) (out model.UserLoginOutput, err error) {
	user := entity.SysUser{}
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().UserName:     in.UserName,
		dao.SysUser.Columns().UserPassword: encryptPassword(in.UserPassword),
	}).Scan(&user)
	if err != nil {
		return out, gerror.New("用户名或密码错误")
	}
	if user.UserStatus == consts.StatusDisabled {
		return model.UserLoginOutput{}, gerror.NewCodef(gcode.CodeInvalidParameter, "用户已被禁用")
	}
	out.UserInfo = convertDbEntityToOutput(user)
	return
}

func (s *sUser) UpdateUser(ctx context.Context, in model.UpdateUserInput) (err error) {
	data := g.Map{
		dao.SysUser.Columns().UserPhone:    in.UserPhone,
		dao.SysUser.Columns().UserName:     in.UserName,
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
		dao.SysUser.Columns().UserPassword: encryptPassword(consts.DefaultPassword),
		dao.SysUser.Columns().UserRealName: in.UserRealName,
		dao.SysUser.Columns().UserPhone:    in.UserPhone,
		dao.SysUser.Columns().UserImage:    in.UserImage,
		dao.SysUser.Columns().UserStatus:   in.UserStatus,
	}

	// 插入并返回自动生成的ID
	id, err := dao.SysUser.Ctx(ctx).InsertAndGetId(user)
	if err != nil {
		return model.AddUserOutput{}, err
	}
	out = model.AddUserOutput{UserId: id}
	return
}

func (s *sUser) GetUserAccessList(ctx context.Context, in model.GetUserAccessListInput) (out model.GetUserAccessListOutput, err error) {
	roleList, err := s.GetUserRoleList(ctx, model.GetUserRoleListInput{UserId: in.UserId})
	if err != nil {
		return model.GetUserAccessListOutput{}, err
	}
	for _, role := range roleList.List {
		accessList, err := service.Role().GetRoleAccessList(ctx, model.GetRoleAccessListInput{RoleId: role.RoleId})
		if err != nil {
			return model.GetUserAccessListOutput{}, err
		}
		for _, access := range accessList.List {
			out.List = append(out.List, access)
		}
	}
	return
}

func (s *sUser) AddUserRole(ctx context.Context, in model.AddUserRoleInput) (err error) {
	_, err = dao.SysUserRole.Ctx(ctx).Insert(g.Map{
		dao.SysUserRole.Columns().UserId: in.UserId,
		dao.SysUserRole.Columns().RoleId: in.RoleId,
	})
	return
}

func (s *sUser) DeleteUserRole(ctx context.Context, in model.DeleteUserRoleInput) (err error) {
	_, err = dao.SysUserRole.Ctx(ctx).Where(g.Map{
		dao.SysUserRole.Columns().UserId: in.UserId,
		dao.SysUserRole.Columns().RoleId: in.RoleId,
	}).Delete()
	return
}

func (s *sUser) GetUserRoleList(ctx context.Context, in model.GetUserRoleListInput) (out model.GetUserRoleListOutput, err error) {
	roleIds, err := dao.SysUserRole.Ctx(ctx).
		Where(dao.SysUserRole.Columns().UserId, in.UserId).
		Array(dao.SysUserRole.Columns().RoleId)
	if err != nil {
		return model.GetUserRoleListOutput{}, err
	}
	for _, roleId := range roleIds {
		output, err := service.Role().GetRoleById(model.GetRoleByIdInput{RoleId: roleId.Int64()})
		if err != nil {
			return model.GetUserRoleListOutput{}, err
		}
		out.List = append(out.List, output.SysRole)
	}
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
