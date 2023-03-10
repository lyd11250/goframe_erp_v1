package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetRoleListReq struct {
	g.Meta `path:"/role/list" method:"post" summary:"获取角色列表" tags:"角色管理"`
}

type GetRoleListRes struct {
	List []entity.SysRole `json:"list"`
}

type AddRoleReq struct {
	g.Meta   `path:"/role/add" method:"post" summary:"新增角色" tags:"角色管理"`
	RoleName string `json:"roleName" dc:"角色名称" v:"required#请输入角色名称" `
}

type AddRoleRes struct {
	RoleId int64 `json:"roleId" dc:"角色ID"`
}

type UpdateRoleReq struct {
	g.Meta   `path:"/role/update" method:"post" summary:"修改角色" tags:"角色管理"`
	RoleId   *int64  `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
	RoleName *string `json:"roleName" dc:"角色名称" v:"required#请输入角色名称"`
}

type UpdateRoleRes struct {
}

type DeleteRoleReq struct {
	g.Meta `path:"/role/delete" method:"post" summary:"删除角色" tags:"角色管理"`
	RoleId int64 `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
}

type DeleteRoleRes struct {
}

type AddRoleAccessReq struct {
	g.Meta   `path:"/role/access/add" method:"post" summary:"角色新增权限" tags:"角色管理"`
	RoleId   int64 `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
	AccessId int64 `json:"accessId" dc:"权限ID" v:"required#请输入权限ID"`
}

type AddRoleAccessRes struct {
}

type DeleteRoleAccessReq struct {
	g.Meta   `path:"/role/access/delete" method:"post" summary:"角色删除权限" tags:"角色管理"`
	RoleId   int64 `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
	AccessId int64 `json:"accessId" dc:"权限ID" v:"required#请输入权限ID"`
}

type DeleteRoleAccessRes struct {
}

type GetRoleAccessReq struct {
	g.Meta `path:"/role/access/list" method:"post" summary:"获取角色所拥有的权限列表" tags:"角色管理"`
	RoleId int64 `json:"roleId" dc:"角色ID" v:"required#请输入角色ID"`
}

type GetRoleAccessRes struct {
	List []entity.SysAccess `json:"list"`
}
