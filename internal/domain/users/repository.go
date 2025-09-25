package users

import (
	"context"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]*User, error)
}
