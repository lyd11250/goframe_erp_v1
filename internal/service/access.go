// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-erp-v1/internal/model"
)

type (
	IAccess interface {
		GetAccessById(ctx context.Context, in model.GetAccessByIdInput) (out model.GetAccessByIdOutput, err error)
		GetRoleAccessList(ctx context.Context, in model.GetRoleAccessListInput) (out model.GetRoleAccessListOutput, err error)
		AddAccess(ctx context.Context, in model.AddAccessInput) (out model.AddAccessOutput, err error)
		UpdateAccess(ctx context.Context, in model.UpdateAccessInput) (err error)
		DeleteAccess(ctx context.Context, in model.DeleteAccessInput) (err error)
	}
)

var (
	localAccess IAccess
)

func Access() IAccess {
	if localAccess == nil {
		panic("implement not found for interface IAccess, forgot register?")
	}
	return localAccess
}

func RegisterAccess(i IAccess) {
	localAccess = i
}
