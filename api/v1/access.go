package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-erp-v1/internal/model/entity"
)

type GetAccessListReq struct {
	g.Meta `path:"/access/list" method:"post" summary:"获取权限列表"`
}

type GetAccessListRes struct {
	List []entity.SysAccess `json:"list"`
}

type AddAccessReq struct {
	g.Meta      `path:"/access/add" method:"post" summary:"新增权限"`
	AccessTitle string `json:"accessTitle" dc:"权限标题" v:"required#请输入权限标题"`
	AccessUri   string `json:"accessUri"  dc:"权限Uri" v:"required#请输入权限Uri"`
}

type AddAccessRes struct {
	AccessId int64 `json:"accessId" dc:"权限ID"`
}

type UpdateAccessReq struct {
	g.Meta      `path:"/access/update" method:"post" summary:"修改权限"`
	AccessId    *int64  `json:"accessId" dc:"权限ID" v:"required#请输入权限ID"`
	AccessTitle *string `json:"accessTitle" dc:"权限标题" `
	AccessUri   *string `json:"accessUri"  dc:"权限Uri" `
}

type UpdateAccessRes struct {
}

type DeleteAccessReq struct {
	g.Meta   `path:"/access/delete" method:"post" summary:"删除权限"`
	AccessId int64 `json:"accessId" dc:"权限ID" v:"required#请输入权限ID"`
}

type DeleteAccessRes struct {
}
