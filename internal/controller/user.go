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
	if req.UserId == nil {
		id := redis.Ctx(ctx).CheckLogin()
		req.UserId = &id
	}
	user, err := service.User().GetUserById(ctx, model.GetUserByIdInput{
		UserId: *req.UserId,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(user, &res)
	if err != nil {
		return nil, err
	}
	roleList, err := service.User().GetUserRoleList(ctx, model.GetUserRoleListInput{UserId: user.UserId})
	if err != nil {
		return nil, err
	}
	res.UserRoles = roleList.List
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
	roleList, err := service.User().GetUserRoleList(ctx, model.GetUserRoleListInput{UserId: user.UserId})
	if err != nil {
		return nil, err
	}
	res.UserRoles = roleList.List
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
	roleList, err := service.User().GetUserRoleList(ctx, model.GetUserRoleListInput{UserId: user.UserId})
	if err != nil {
		return nil, err
	}
	res.UserRoles = roleList.List
	return
}

func (c *cUser) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	redis.Ctx(ctx).Logout()
	return
}

func (c *cUser) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	input := model.UpdateUserInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.User().UpdateUser(ctx, input)
	if err != nil {
		return nil, err
	}
	res = &v1.UpdateUserRes{}
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

	// 用户默认拥有”系统用户“角色
	err = service.User().AddUserRole(ctx, model.AddUserRoleInput{
		UserId: output.UserId,
		RoleId: 1,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.AddUserRes{UserId: output.UserId}
	return
}

func (c *cUser) AddUserRole(ctx context.Context, req *v1.AddUserRoleReq) (res *v1.AddUserRoleRes, err error) {
	err = service.User().AddUserRole(ctx, model.AddUserRoleInput{
		UserId: req.UserId,
		RoleId: req.RoleId,
	})
	return
}

func (c *cUser) DeleteUserRole(ctx context.Context, req *v1.DeleteUserRoleReq) (res *v1.DeleteUserRoleRes, err error) {
	err = service.User().DeleteUserRole(ctx, model.DeleteUserRoleInput{
		UserId: req.UserId,
		RoleId: req.RoleId,
	})
	return
}

func (c *cUser) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	out, err := service.User().GetUserList(ctx)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(out, &res)
	if err != nil {
		return nil, err
	}
	for i, user := range res.List {
		roleList, err := service.User().GetUserRoleList(ctx, model.GetUserRoleListInput{UserId: user.UserId})
		if err != nil {
			return nil, err
		}
		res.List[i].UserRoles = roleList.List
	}
	return
}
