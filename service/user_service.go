package service

import (
	"EcommersAPIHP/model/web"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) web.UserResponse
}
