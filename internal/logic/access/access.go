package access

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-erp-v1/internal/dao"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
)

type sAccess struct {
}

func New() *sAccess {
	return &sAccess{}
}

func init() {
	service.RegisterAccess(New())
}

func (s *sAccess) GetAccessById(ctx context.Context, in model.GetAccessByIdInput) (out model.GetAccessByIdOutput, err error) {
	err = dao.SysAccess.Ctx(ctx).WherePri(in.AccessId).Scan(&out)
	return
}

func (s *sAccess) GetRoleAccessList(ctx context.Context, in model.GetRoleAccessListInput) (out model.GetRoleAccessListOutput, err error) {
	result, err := dao.SysRoleAccess.Ctx(ctx).All(dao.SysRoleAccess.Columns().RoleId, in.RoleId)
	if err != nil {
		return model.GetRoleAccessListOutput{}, err
	}
	for _, row := range result.List() {
		accessId := gconv.Int64(row["access_id"])
		access, err := s.GetAccessById(ctx, model.GetAccessByIdInput{AccessId: accessId})
		if err != nil {
			return model.GetRoleAccessListOutput{}, err
		}
		out.List = append(out.List, access.Access)
	}
	return
}

func (s *sAccess) AddAccess(ctx context.Context, in model.AddAccessInput) (out model.AddAccessOutput, err error) {
	id, err := dao.SysAccess.Ctx(ctx).InsertAndGetId(g.Map{
		dao.SysAccess.Columns().AccessTitle: in.AccessTitle,
		dao.SysAccess.Columns().AccessUri:   in.AccessUri,
	})
	if err != nil {
		return model.AddAccessOutput{}, err
	}
	out.AccessId = id
	return
}

func (s *sAccess) UpdateAccess(ctx context.Context, in model.UpdateAccessInput) (err error) {
	_, err = dao.SysAccess.Ctx(ctx).OmitEmpty().WherePri(in.AccessId).Data(g.Map{
		dao.SysAccess.Columns().AccessTitle: in.AccessTitle,
		dao.SysAccess.Columns().AccessUri:   in.AccessUri,
	}).Update()
	return
}

func (s *sAccess) DeleteAccess(ctx context.Context, in model.DeleteAccessInput) (err error) {
	_, err = dao.SysAccess.Ctx(ctx).WherePri(in.AccessId).Delete()
	return
}
