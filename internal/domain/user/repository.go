package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindById(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindAll(ctx context.Context, limit int, offset int) ([]*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
}
