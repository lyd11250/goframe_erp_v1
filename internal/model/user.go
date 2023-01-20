package model

type UserInfo struct {
	UserId       int64
	UserName     string
	UserRealName string
	UserPhone    string
	UserImage    string
	UserStatus   uint
}

type GetUserByIdInput struct {
	UserId int64
}

type GetUserByIdOutput struct {
	UserInfo
}

type GetUserByUserNameInput struct {
	UserName string
}

type GetUserByUserNameOutput struct {
	UserInfo
}

type UserLoginInput struct {
	UserName     string
	UserPassword string
}

type UserLoginOutput struct {
	UserInfo
}

type UpdateUserByIdInput struct {
	UserId       int64
	UserPassword *string
	UserRealName *string
	UserPhone    *string
	UserImage    *string
	UserStatus   *uint
}

type AddUserInput struct {
	UserName     string
	UserPassword string
	UserRealName string
	UserPhone    string
	UserImage    string
}

type AddUserOutput struct {
	UserId int64
}

type GetUserAccessListInput struct {
	UserId int64
}

type GetUserAccessListOutput struct {
	List []Access
}

type AddUserRoleInput struct {
	UserId int64
	RoleId int64
}

type DeleteUserRoleInput struct {
	UserId int64
	RoleId int64
}
