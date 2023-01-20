package model

type Access struct {
	AccessId    int64  // 权限ID
	AccessTitle string // 权限标题
	AccessUri   string // 权限路径
}

type GetAccessByIdInput struct {
	AccessId int64
}

type GetAccessByIdOutput struct {
	Access
}

type GetRoleAccessListInput struct {
	RoleId int64
}

type GetRoleAccessListOutput struct {
	List []Access
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
