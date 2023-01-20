package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfo struct {
	UserId       int64  `json:"userId"       ` // 用户ID，主键
	UserName     string `json:"userName"     ` // 登录用户名
	UserRealName string `json:"userRealName" ` // 用户真实姓名
	UserPhone    string `json:"userPhone"    ` // 用户手机号，11位数字
	UserImage    string `json:"userImage"    ` // 用户头像url
	UserStatus   uint   `json:"userStatus"   ` // 用户状态
}

type GetUserByIdReq struct {
	g.Meta `path:"/user/id" method:"post" summary:"通过ID获取用户信息"`
	UserId int64 `json:"userId" dc:"用户ID" v:"required#请输入用户ID"`
}

type GetUserByIdRes struct {
	UserInfo
}

type GetUserByUserNameReq struct {
	g.Meta   `path:"/user/name" method:"post" summary:"通过用户名获取用户信息"`
	UserName string `json:"userName" dc:"用户名" v:"required#请输入用户名"`
}

type GetUserByUserNameRes struct {
	UserInfo
}

type UserLoginReq struct {
	g.Meta       `path:"/user/login" method:"post" summary:"用户登录"`
	UserName     string `json:"userName" dc:"用户名" v:"required#请输入用户名"`
	UserPassword string `json:"userPassword" dc:"密码" v:"required#请输入密码"`
}

type UserLoginRes struct {
	UserInfo
}

type UpdateUserByIdReq struct {
	g.Meta       `path:"/user/update/id" method:"post" summary:"通过ID修改用户信息"`
	UserId       *int64  `json:"userId" dc:"用户ID" v:"required#请输入用户名"`
	UserPassword *string `json:"userPassword" dc:"密码"`
	UserRealName *string `json:"userRealName" dc:"真实姓名"`
	UserPhone    *string `json:"userPhone" dc:"手机号"`
	UserImage    *string `json:"userImage" dc:"头像url"`
	UserStatus   *uint   `json:"userStatus" dc:"用户状态"`
}

type UpdateUserByIdRes struct {
}

type AddUserReq struct {
	g.Meta       `path:"/user/add" method:"post" summary:"新增用户"`
	UserName     string `json:"userName" dc:"用户名" v:"required#请输入用户名"`
	UserPassword string `json:"userPassword" dc:"密码" v:"required#请输入密码"`
	UserRealName string `json:"userRealName" dc:"真实姓名" v:"required#请输入真实姓名"`
	UserPhone    string `json:"userPhone" dc:"手机号码" v:"required#请输入手机号码"`
	UserImage    string `json:"userImage" dc:"头像url" `
}

type AddUserRes struct {
	UserId int64 `json:"userId" dc:"用户ID"`
}
