package pojo

import "goframe-erp-v1/internal/model/entity"

type GetRoleListOutput struct {
	List []entity.SysRole
}

type GetRoleByIdInput struct {
	RoleId int64
}

type GetRoleByIdOutput struct {
	entity.SysRole
}

type AddRoleInput struct {
	RoleName string
}

type AddRoleOutput struct {
	RoleId int64
}

type UpdateRoleInput struct {
	RoleId   *int64
	RoleName *string
}

type DeleteRoleInput struct {
	RoleId int64
}

type AddRoleAccessInput struct {
	RoleId   int64
	AccessId int64
}

type DeleteRoleAccessInput struct {
	RoleId   int64
	AccessId int64
}

type GetRoleAccessListInput struct {
	RoleId int64
}

type GetRoleAccessListOutput struct {
	List []entity.SysAccess
}
