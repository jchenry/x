package model

// import "database/sql"

// type Storage interface {
// 	Insert(ctxt StorageContext, val interface{}) error
// 	Read(ctxt StorageContext, val interface{}, columns ...string) error
// 	Select(ctxt StorageContext, val interface{}, query string, args ...interface{}) error
// 	Write(ctxt StorageContext, val interface{}, columns ...string) error
// 	Delete(ctxt StorageContext, val interface{}) error
// }

// type StorageContext interface {
// 	Exec(query string, args ...interface{}) (sql.Result, error)
// 	Query(query string, args ...interface{}) (*sql.Rows, error)
// }

// type ModelService struct {
// 	db  Storage
// 	ctx *StorageContext
// }

// // Find returns a pointer to an array of the results found based on params
// // or an error
// func (s *ModelService) Find(entityArrPtr interface{}, params map[string]interface{}) (err error) {
// 	s.db.Select(s.ctx, )

// }

// // Create returns the identifier for the newly accepted entity, or error
// func (s *ModelService) Create(entityPtr interface{}) (id interface{}, err error) {}

// // Update returns the id of the newly updated entity, or error
// func (s *ModelService) Update(entityPtr interface{}) (id interface{}, err error) {}

// // Delete returns whether the entity, specified by id, was successfully deleted
// // or error
// func (s *ModelService) Delete(entityPtr interface{}) error {}
