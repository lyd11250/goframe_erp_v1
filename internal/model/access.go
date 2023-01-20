package model

import "goframe-erp-v1/internal/model/entity"

type GetAccessListOutput struct {
	List []entity.SysAccess
}

type GetAccessByIdInput struct {
	AccessId int64
}

type GetAccessByIdOutput struct {
	entity.SysAccess
}

type AddAccessInput struct {
	AccessTitle string // 权限标题
	AccessUri   string // 权限路径
}

type AddAccessOutput struct {
	AccessId int64
}

type UpdateAccessInput struct {
	AccessId    int64   // 权限ID
	AccessTitle *string // 权限标题
	AccessUri   *string // 权限路径
}

type DeleteAccessInput struct {
	AccessId int64
}
