package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"goframe-erp-v1/internal/consts"
	"goframe-erp-v1/internal/model/pojo"
	"goframe-erp-v1/internal/service"
	"goframe-erp-v1/utility/redis"
	"goframe-erp-v1/utility/response"
)

func (s *sMiddleware) AccessHandler(r *ghttp.Request) {
	ctx := r.GetCtx()
	uri := r.RequestURI
	// 跳过拦截登录请求
	if uri == "/user/login" {
		r.Middleware.Next()
		return
	}

	// 登录状态验证
	loginId := redis.Ctx(ctx).CheckLogin()
	if loginId == 0 {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "登录状态失效")
		return
	}

	// 账号禁用验证
	userInfo, err := service.User().GetUserById(ctx, pojo.GetUserByIdInput{UserId: loginId})
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "获取登录信息失败")
		return
	}
	if userInfo.UserStatus == consts.StatusDisabled {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "账号已被禁用")
		return
	}

	// 权限验证
	accessList, err := service.User().GetUserAccessList(ctx, pojo.GetUserAccessListInput{UserId: loginId})
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "权限认证失败，请联系系统管理员")
		return
	}
	for _, access := range accessList.List {
		// 用户此次访问的uri存在与权限列表中
		if matchUri(access.AccessUri, uri) {
			r.Middleware.Next()
			return
		}
	}
	response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "用户权限不足")
}

func matchUri(str, uri string) bool {
	/*
		/user/*		/user/add	->	true
		/user/a		/user/add 	->	false
		/user		/user/add	-> 	false
	*/
	if str == uri {
		return true
	}
	if str[len(str)-1] == "*"[0] {
		// "/user/*" -> "/user"
		prefix := gstr.StrTillEx(str, "/*")
		return 0 == gstr.Pos(uri, prefix)
	}
	return false
}
