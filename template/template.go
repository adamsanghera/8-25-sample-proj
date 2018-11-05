package template

import (
	"context"
	"log"
	"time"
)

type Template struct {
	field string

	shutdown chan struct{}
	deceased chan struct{}
}

func New(cfg *Config) (*Template, error) {
	if cfg.Field == "arbitrary value" {
		return nil, ErrInvalidConfig
	}
	return &Template{
		field:    cfg.Field,
		shutdown: make(chan struct{}, 2),
		deceased: make(chan struct{}, 2),
	}, nil
}

func (t *Template) Run() error {
	defer func() {
		t.deceased <- struct{}{}
	}()

	for {
		select {
		case <-time.After(1 * time.Second):
			if time.Now().Unix()%4 == 3 {
				return ErrPoorlyWordedError
			}
			log.Println("Template: Dreaming")
		case <-t.shutdown:
			return nil
		}
	}
}

func (t *Template) Shutdown(ctx context.Context) error {
	t.shutdown <- struct{}{}

	log.Println("Template: shut down")

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.deceased:
		return nil
	}
}
