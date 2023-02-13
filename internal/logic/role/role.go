package role

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/model/entity"
	"goframe-erp-v1/internal/service"
)

type sRole struct {
	RoleList []entity.SysRole
}

func New() *sRole {
	roleService := &sRole{}
	err := dao.SysRole.Ctx(gctx.New()).Scan(&roleService.RoleList)
	if err != nil {
		return nil
	}
	return roleService
}

func init() {
	service.RegisterRole(New())
}

func (s *sRole) GetRoleList() (out model.GetRoleListOutput, err error) {
	out.List = s.RoleList
	return
}

func (s *sRole) GetRoleById(in model.GetRoleByIdInput) (out model.GetRoleByIdOutput, err error) {
	for _, role := range s.RoleList {
		if role.RoleId == in.RoleId {
			out.SysRole = role
			return
		}
	}
	return model.GetRoleByIdOutput{}, gerror.NewCode(gcode.CodeInvalidParameter, "角色不存在")
}

func (s *sRole) AddRole(ctx context.Context, in model.AddRoleInput) (out model.AddRoleOutput, err error) {
	id, err := dao.SysRole.Ctx(ctx).InsertAndGetId(g.Map{
		dao.SysRole.Columns().RoleName: in.RoleName,
	})
	if err != nil {
		return model.AddRoleOutput{}, err
	}
	out.RoleId = id
	return
}

func (s *sRole) UpdateRole(ctx context.Context, in model.UpdateRoleInput) (err error) {
	_, err = dao.SysRole.Ctx(ctx).WherePri(in.RoleId).Data(g.Map{
		dao.SysRole.Columns().RoleName: in.RoleName,
	}).Update()
	return
}

func (s *sRole) DeleteRole(ctx context.Context, in model.DeleteRoleInput) (err error) {
	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除sys_role表中的数据
		_, e := tx.Model(dao.SysRole.Table()).WherePri(in.RoleId).Delete()
		if e != nil {
			return e
		}

		// 删除sys_role_access表中的相关数据
		_, e = tx.Model(dao.SysRoleAccess.Table()).Where(dao.SysRoleAccess.Columns().RoleId, in.RoleId).Delete()
		if e != nil {
			return e
		}

		// 删除sys_user_role表中的相关数据
		_, e = tx.Model(dao.SysUserRole.Table()).Where(dao.SysUserRole.Columns().RoleId, in.RoleId).Delete()
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func (s *sRole) AddRoleAccess(ctx context.Context, in model.AddRoleAccessInput) (err error) {
	_, err = dao.SysRoleAccess.Ctx(ctx).Insert(g.Map{
		dao.SysRoleAccess.Columns().RoleId:   in.RoleId,
		dao.SysRoleAccess.Columns().AccessId: in.AccessId,
	})
	return
}

func (s *sRole) DeleteRoleAccess(ctx context.Context, in model.DeleteRoleAccessInput) (err error) {
	_, err = dao.SysRoleAccess.Ctx(ctx).Where(g.Map{
		dao.SysRoleAccess.Columns().AccessId: in.AccessId,
		dao.SysRoleAccess.Columns().RoleId:   in.RoleId,
	}).Delete()
	return
}

func (s *sRole) GetRoleAccessList(ctx context.Context, in model.GetRoleAccessListInput) (out model.GetRoleAccessListOutput, err error) {
	result, err := dao.SysRoleAccess.Ctx(ctx).
		Fields(dao.SysRoleAccess.Columns().AccessId).
		All(dao.SysRoleAccess.Columns().RoleId, in.RoleId)
	if err != nil {
		return model.GetRoleAccessListOutput{}, err
	}
	for _, accessId := range result.Array() {
		output, err := service.Access().GetAccessById(model.GetAccessByIdInput{AccessId: accessId.Int64()})
		if err != nil {
			return model.GetRoleAccessListOutput{}, err
		}
		out.List = append(out.List, output.SysAccess)
	}
	return
}
