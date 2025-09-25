package application

import (
	"context"
	"fmt"

	"asana/internal/infrastructure/httpx"
	"asana/internal/infrastructure/poller"
	"asana/pkg/config"
	"asana/pkg/log"

	"asana/internal/domain/projects"
	"asana/internal/domain/users"
)

type Application struct {
	cfg         *config.Config
	asanaClient *httpx.Client
	poller      *poller.Poller

	userRepo users.Repository
	prjRepo  projects.Repository
}

func New() *Application {
	return &Application{
		cfg:         config.Get(),
		asanaClient: httpx.New(),
	}
}

func (a *Application) Run(ctx context.Context) error {
	systems := []struct {
		name string
		fn   func(context.Context) error
	}{
		{"Logger", a.setupLogger},
		{"repositories", a.initRepositories},
		{"poller", a.startPoller},
	}

	for _, system := range systems {
		fmt.Printf("Initializing %s...\n", system.name)
		if err := system.fn(ctx); err != nil {
			return fmt.Errorf("failed to initialize %s: %w", system.name, err)
		}
	}

	return nil
}

func (a *Application) setupLogger(ctx context.Context) error {
	if err := log.Setup(log.Console, ""); err != nil {
		return fmt.Errorf("failed to set up logger: %w", err)
	}
	return nil
}
