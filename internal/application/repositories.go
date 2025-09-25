package application

import (
	"asana/internal/infrastructure/httpx/projects"
	"asana/internal/infrastructure/httpx/users"
	"context"
)

func (a *Application) initRepositories(ctx context.Context) error {
	a.userRepo = users.New(a.asanaClient)
	a.prjRepo = projects.New(a.asanaClient)

	return nil
}
