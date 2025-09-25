package poller

import (
	"asana/pkg/log"
	"context"
	"sync"
	"time"
)

type Task struct {
	Name     string
	Interval time.Duration
	Fn       func(ctx context.Context) error
}

type Poller struct {
	tasks []Task
}

func New() *Poller {
	return &Poller{}
}

func (p *Poller) AddTask(t Task) {
	p.tasks = append(p.tasks, t)
}

func (p *Poller) Start(ctx context.Context) error {
	var wg sync.WaitGroup

	for _, task := range p.tasks {
		t := task
		wg.Add(1)
		go func() {
			defer wg.Done()
			ticker := time.NewTicker(t.Interval)
			defer ticker.Stop()

			if err := t.Fn(ctx); err != nil {
				log.Error("Task %s failed: %v", t.Name, err)
			}

			for {
				select {
				case <-ticker.C:
					if err := t.Fn(ctx); err != nil {
						log.Error("Task %s failed: %v", t.Name, err)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	wg.Wait()
	return nil
}
