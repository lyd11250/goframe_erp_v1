package controller

import (
	"context"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model/pojo"
	"goframe-erp-v1/internal/service"
)

type cRole struct {
}

var Role cRole

func (c *cRole) AddRole(ctx context.Context, req *v1.AddRoleReq) (res *v1.AddRoleRes, err error) {
	output, err := service.Role().AddRole(ctx, pojo.AddRoleInput{RoleName: req.RoleName})
	if err != nil {
		return nil, err
	}
	res = &v1.AddRoleRes{RoleId: output.RoleId}
	return
}

func (c *cRole) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error) {
	err = service.Role().UpdateRole(ctx, pojo.UpdateRoleInput{
		RoleId:   req.RoleId,
		RoleName: req.RoleName,
	})
	return
}

func (c *cRole) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	err = service.Role().DeleteRole(ctx, pojo.DeleteRoleInput{RoleId: req.RoleId})
	return
}

func (c *cRole) AddRoleAccess(ctx context.Context, req *v1.AddRoleAccessReq) (res *v1.AddRoleAccessRes, err error) {
	err = service.Role().AddRoleAccess(ctx, pojo.AddRoleAccessInput{
		RoleId:   req.RoleId,
		AccessId: req.AccessId,
	})
	return
}

func (c *cRole) DeleteRoleAccess(ctx context.Context, req *v1.DeleteRoleAccessReq) (res *v1.DeleteRoleAccessRes, err error) {
	err = service.Role().DeleteRoleAccess(ctx, pojo.DeleteRoleAccessInput{
		RoleId:   req.RoleId,
		AccessId: req.AccessId,
	})
	return
}

func (c *cRole) GetRoleList(ctx context.Context, req *v1.GetRoleListReq) (res *v1.GetRoleListRes, err error) {
	output, err := service.Role().GetRoleList()
	if err != nil {
		return nil, err
	}
	res = &v1.GetRoleListRes{List: output.List}
	return
}

func (c *cRole) GetRoleAccessList(ctx context.Context, req *v1.GetRoleAccessReq) (res *v1.GetRoleAccessRes, err error) {
	output, err := service.Role().GetRoleAccessList(ctx, pojo.GetRoleAccessListInput{RoleId: req.RoleId})
	if err != nil {
		return nil, err
	}
	res = &v1.GetRoleAccessRes{List: output.List}
	return
}
