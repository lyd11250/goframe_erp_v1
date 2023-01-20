package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "goframe-erp-v1/api/v1"
	"goframe-erp-v1/internal/model"
	"goframe-erp-v1/internal/service"
	"goframe-erp-v1/utility/redis"
)

type cUser struct {
}

var User cUser

func (c *cUser) GetUserById(ctx context.Context, req *v1.GetUserByIdReq) (res *v1.GetUserByIdRes, err error) {
	user, err := service.User().GetUserById(ctx, model.GetUserByIdInput{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(user, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cUser) GetUserByUserName(ctx context.Context, req *v1.GetUserByUserNameReq) (res *v1.GetUserByUserNameRes, err error) {
	user, err := service.User().GetUserByUserName(ctx, model.GetUserByUserNameInput{
		UserName: req.UserName,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(user, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cUser) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	user, err := service.User().UserLogin(ctx, model.UserLoginInput{
		UserName:     req.UserName,
		UserPassword: req.UserPassword,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(user, &res)
	if err != nil {
		return nil, err
	}
	redis.Ctx(ctx).Login(res.UserId)
	return
}

func (c *cUser) UpdateUserById(ctx context.Context, req *v1.UpdateUserByIdReq) (res *v1.UpdateUserByIdRes, err error) {
	input := model.UpdateUserByIdInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.User().UpdateUserById(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.UpdateUserByIdRes{}
	return
}

func (c *cUser) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {
	input := model.AddUserInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	output, err := service.User().AddUser(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.AddUserRes{UserId: output.UserId}
	return
}
