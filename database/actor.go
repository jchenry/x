package database

import (
	"context"
	"database/sql"
)

type Func func(db *sql.DB)

type Actor struct {
	DB         *sql.DB
	ActionChan <-chan Func
}

func (a *Actor) Run(ctx context.Context) error {
	for {
		select {
		case f := <-a.ActionChan:
			f(a.DB)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
