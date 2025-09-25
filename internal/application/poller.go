package application

import (
	"asana/internal/infrastructure/poller"
	"context"

	"asana/pkg/log"
	"asana/pkg/storage"
)

func (a *Application) startPoller(ctx context.Context) error {
	a.poller = poller.New()

	path := a.cfg.Cache.Path
	if path == "" {
		path = "cache"
	}
	storage.CreatePath(path)

	a.poller.AddTask(poller.Task{
		Name:     "First Task",
		Interval: a.cfg.Fetch.PollInterval,
		Fn: func(ctx context.Context) error {
			if err := a.Fetch(ctx, path); err != nil {
				log.Error("Failed to fetch data: %v", err)
			}

			return nil
		},
	})

	a.poller.AddTask(poller.Task{
		Name:     "Second Task",
		Interval: a.cfg.Fetch.SecondPollInterval,
		Fn: func(ctx context.Context) error {
			if err := a.Fetch(ctx, path); err != nil {
				log.Error("Failed to fetch data: %v", err)
			}

			return nil
		},
	})

	if err := a.poller.Start(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Application) Fetch(ctx context.Context, path string) error {
	log.Info("Fetch data from Asana")

	users, err := a.userRepo.GetUsers(ctx)
	if err != nil {
		return err
	}
	log.Info("Fetched %d users", len(users))

	for _, user := range users {
		if _, err := storage.WriteJson(path, user.Name, user); err != nil {
			log.Error("Failed to write user %s: %v", user.Gid, err)
		}
	}

	prj, err := a.prjRepo.GetProjects(ctx)
	if err != nil {
		return err
	}
	log.Info("Fetched %d projects", len(prj))

	for _, project := range prj {
		if _, err := storage.WriteJson(path, project.Name, project); err != nil {
			log.Error("Failed to write project %s: %v", project.Gid, err)
		}
	}

	return nil
}
