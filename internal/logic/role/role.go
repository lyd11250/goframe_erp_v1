package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sRole struct {
}

func New() *sRole {
	return &sRole{}
}

func init() {
	service.RegisterRole(New())
}

func (s *sRole) GetRoleById(ctx context.Context, in model.GetRoleByIdInput) (out model.GetRoleByIdOutput, err error) {
	err = dao.SysRole.Ctx(ctx).WherePri(in.RoleId).Scan(&out)
	return
}

func (s *sRole) GetUserRoleList(ctx context.Context, in model.GetUserRoleListInput) (out model.GetUserRoleListOutput, err error) {
	result, err := dao.SysUserRole.Ctx(ctx).All(dao.SysUserRole.Columns().UserId, in.UserId)
	for _, row := range result.List() {
		roleId := gconv.Int64(row["role_id"])
		role, err := s.GetRoleById(ctx, model.GetRoleByIdInput{RoleId: roleId})
		if err != nil {
			continue
		}
		out.List = append(out.List, role.Role)
	}
	return
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
	_, err = dao.SysRole.Ctx(ctx).WherePri(in.RoleId).Delete()
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
