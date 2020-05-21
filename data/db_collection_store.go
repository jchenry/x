package data

import (
	"database/sql"

	"rsc.io/dbstore"
)

type DatabaseCollectionInstance struct {
	db    *sql.DB
	store *dbstore.Storage
}

func DatabaseCollection(
	db *sql.DB,
	createTables bool,
	types ...interface{}) (*DatabaseCollectionInstance, error) {

	dbi := &DatabaseCollectionInstance{
		db:    db,
		store: new(dbstore.Storage),
	}

	for _, t := range types {
		dbi.store.Register(t)
	}

	if err := dbi.store.CreateTables(db); err != nil {
		return nil, err
	}

	return dbi, nil
}

// Find returns a pointer to an array of the results found based on params
// or an error
func (d *DatabaseCollectionInstance) Find(entityArrPtr interface{}, params map[string]interface{}) error {
	return d.store.Select(d.db, entityArrPtr, "")
}

// Create returns the identifier for the newly accepted entity, or error
func (d *DatabaseCollectionInstance) Create(entityPtr interface{}) error {
	return d.store.Insert(d.db, entityPtr)
}

// Update returns the id of the newly updated entity, or error
func (d *DatabaseCollectionInstance) Update(entityPtr interface{}) error {
	return d.store.Insert(d.db, entityPtr)
	// return nil
}

// Delete returns whether the entity, specified by id, was successfully deleted
// or error
func (d *DatabaseCollectionInstance) Delete(entityPtr interface{}) error {
	return d.store.Delete(d.db, entityPtr)
}
