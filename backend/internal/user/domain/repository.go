package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*User, error)
	CountByRole(ctx context.Context, role Role) (int64, error)
}
