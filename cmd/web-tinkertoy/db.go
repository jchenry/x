package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jchenry/jchenry/db"
	"rsc.io/dbstore"
)

type DBFunc func(db *sql.DB)

type DBActor struct {
	DB         *sql.DB
	ActionChan chan DBFunc
}

func (a *DBActor) Run(ctx context.Context) error {
	for {
		select {
		case f := <-a.ActionChan:
			f(a.DB)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func DBStoreInsert(store *dbstore.Storage, e interface{}) db.Func {
	return func(db *sql.DB) {
		err := store.Insert(db, e)
		fmt.Println(err)
	}
}

func DBStoreDelete(store *dbstore.Storage, e interface{}) db.Func {
	return func(db *sql.DB) {
		store.Delete(db, e)
	}
}

func DBStoreSelect(store *dbstore.Storage,
	err chan error,
	results chan interface{},
	ent interface{},
	query string,
	args ...interface{}) db.Func {
	return func(db *sql.DB) {
		if e := store.Select(db, ent, query, args...); e != nil {
			err <- e
		} else {
			results <- ent
		}
	}
}

func DBStoreRead(store *dbstore.Storage,
	err chan error,
	results chan interface{},
	ent interface{},
	columns ...string) db.Func {
	return func(db *sql.DB) {
		if e := store.Read(db, ent, columns...); e != nil {
			err <- e
		} else {
			results <- ent
		}
	}
}
