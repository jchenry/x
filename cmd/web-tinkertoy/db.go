package main

import (
	"context"
	"database/sql"
	"fmt"

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

func DBStoreInsert(store *dbstore.Storage, e interface{}) DBFunc {
	return func(db *sql.DB) {
		err := store.Insert(db, e)
		fmt.Println(err)

	}
}

func DBStoreDelete(store *dbstore.Storage, e interface{}) DBFunc {
	return func(db *sql.DB) {
		store.Delete(db, e)
	}
}

func DBStoreSelect(store *dbstore.Storage, e interface{}, query string, args ...interface{}) DBFunc {
	return func(db *sql.DB) {
		err := store.Select(db, e, query, args...)
		fmt.Println(err)
	}
}

// func DBStoreRead(store *dbstore.Storage, e interface{}, IDCol string) DBFunc {
// 	return func(db *sql.DB) {
// 		store.Read(db, e, IDCol)
// 	}
// }
