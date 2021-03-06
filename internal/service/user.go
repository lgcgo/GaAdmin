// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"GaAdmin/internal/model"
	"GaAdmin/internal/model/entity"
	"context"
)

type IUser interface {
	SetupRoles(ctx context.Context, userId uint, role_ids []uint) error
	DeleteRolesByGroupId(ctx context.Context, roleId uint) error
	GetAllRoles(ctx context.Context) ([]*entity.UserAccess, error)
	CreateUser(ctx context.Context, in *model.UserCreateInput) (*entity.User, error)
	GetUser(ctx context.Context, userId uint) (*entity.User, error)
	GetUsers(ctx context.Context, userIds []uint) ([]*entity.User, error)
	GetUserByUuid(ctx context.Context, uuid string) (*entity.User, error)
	UpdateUser(ctx context.Context, in *model.UserUpdateInput) (*entity.User, error)
	DeleteUser(ctx context.Context, userId uint) error
	GetUserPage(ctx context.Context, in *model.Page) (*model.UserPageOutput, error)
	GetCurrentUser(ctx context.Context) (*entity.User, error)
	UpdateUserAccount(ctx context.Context, account string) error
	UpdateCurrentUserMobile(ctx context.Context, mobile string) error
	UpdateCurrentUserEmail(ctx context.Context, email string) error
	UpdateCurrentUserPassword(ctx context.Context, password string) error
	SignPassport(ctx context.Context, in *model.UserSignPassportInput) (*entity.User, error)
	SignMobile(ctx context.Context, in *model.UserSignMobile) (*entity.User, error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
