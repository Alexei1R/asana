package projects

import (
	"context"
)

type Repository interface {
	GetProjects(ctx context.Context) ([]*Project, error)
}
