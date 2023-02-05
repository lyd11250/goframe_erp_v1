package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type UserInfo struct {
	UserId       int64            `json:"userId"       ` // 用户ID，主键
	UserName     string           `json:"userName"     ` // 登录用户名
	UserRealName string           `json:"userRealName" ` // 用户真实姓名
	UserPhone    string           `json:"userPhone"    ` // 用户手机号，11位数字
	UserImage    string           `json:"userImage"    ` // 用户头像url
	UserStatus   uint             `json:"userStatus"   ` // 用户状态
	UserRoles    []entity.SysRole `json:"userRoles"`
}

type GetUserByIdReq struct {
	g.Meta `path:"/user/id" method:"post" summary:"通过ID获取用户信息"`
	UserId *int64 `json:"userId" dc:"用户ID"`
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

type UserLogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" summary:"用户登出"`
}

type UserLogoutRes struct {
}

type UpdateUserReq struct {
	g.Meta       `path:"/user/update" method:"post" summary:"修改用户"`
	UserId       *int64  `json:"userId" dc:"用户ID" v:"required#请输入用户名"`
	UserPassword *string `json:"userPassword" dc:"密码"`
	UserRealName *string `json:"userRealName" dc:"真实姓名"`
	UserPhone    *string `json:"userPhone" dc:"手机号"`
	UserImage    *string `json:"userImage" dc:"头像url"`
	UserStatus   *uint   `json:"userStatus" dc:"用户状态"`
}

type UpdateUserRes struct {
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

type AddUserRoleReq struct {
	g.Meta `path:"/user/role/add" method:"post" summary:"用户新增角色"`
	UserId int64 `json:"userId" dc:"用户ID" v:"required#请输入用户ID"`
	RoleId int64 `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
}

type AddUserRoleRes struct {
}

type DeleteUserRoleReq struct {
	g.Meta `path:"/user/role/delete" method:"post" summary:"用户删除角色"`
	UserId int64 `json:"userId" dc:"用户ID" v:"required#请输入用户ID"`
	RoleId int64 `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
}

type DeleteUserRoleRes struct {
}

type GetUserListReq struct {
	g.Meta `path:"/user/list" method:"post" summary:"获取所有用户"`
}

type GetUserListRes struct {
	List []UserInfo `json:"list"`
}
