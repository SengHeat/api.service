package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
	GetUser(ctx context.Context, id int) (*User, error)
	GetUsers(ctx context.Context, page, pageSize int) ([]*User, int, error)
	UpdateUser(ctx context.Context, id int, req *UpdateUserRequest) (*User, error)
	Delete(ctx context.Context, id int) error
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	ValidateToken(ctx context.Context, token string) (*User, error)
}
