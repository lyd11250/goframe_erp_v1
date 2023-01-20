package model

type Role struct {
	RoleId   int64  // 角色ID，主键
	RoleName string // 角色名称
}

type GetRoleByIdInput struct {
	RoleId int64
}

type GetRoleByIdOutput struct {
	Role
}

type GetUserRoleListInput struct {
	UserId int64
}

type GetUserRoleListOutput struct {
	List []Role
}

type AddRoleInput struct {
	RoleName string
}

type AddRoleOutput struct {
	RoleId int64
}

type UpdateRoleInput struct {
	RoleId   *int64
	RoleName *string
}

type DeleteRoleInput struct {
	RoleId int64
}

type AddRoleAccessInput struct {
	RoleId   int64
	AccessId int64
}

type DeleteRoleAccessInput struct {
	RoleId   int64
	AccessId int64
}
