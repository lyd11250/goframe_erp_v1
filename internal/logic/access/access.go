package access

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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

func (s *sAccess) GetAccessList(ctx context.Context) (out model.GetAccessListOutput, err error) {
	err = dao.SysAccess.Ctx(ctx).Scan(&out.List)
	return
}

func (s *sAccess) GetAccessById(ctx context.Context, in model.GetAccessByIdInput) (out model.GetAccessByIdOutput, err error) {
	err = dao.SysAccess.Ctx(ctx).WherePri(in.AccessId).Scan(&out)
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
	err = dao.SysAccess.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除sys_access表中的数据
		_, e := tx.Model(dao.SysAccess.Table()).WherePri(in.AccessId).Delete()
		if e != nil {
			return e
		}

		// 删除sys_role_access表中的相关数据
		_, e = tx.Model(dao.SysRoleAccess.Table()).Where(dao.SysRoleAccess.Columns().AccessId, in.AccessId).Delete()
		if e != nil {
			return e
		}
		return nil
	})
	return
}
