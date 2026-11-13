package users

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id string) error
}
