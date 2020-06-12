package crud

import (
	"github.com/jchenry/jchenry/pkg/db"
)

type Service interface {
	// Find returns a pointer to an array of the results found based on params
	// or an error
	Find(entityArrPtr interface{}, params map[string]interface{}) (err error)
	// Create returns the identifier for the newly accepted entity, or error
	Create(entityPtr interface{}) (err error)
	// Update returns the id of the newly updated entity, or error
	Update(entityPtr interface{}) (err error)
	// Delete returns whether the entity, specified by id, was successfully deleted
	// or error
	Delete(entityPtr interface{}) error
}

type Storage struct {
	Actor    db.Actor
	FindOp   func(entityArrPtr interface{}, params map[string]interface{}) db.Func
	CreateOp func(entityPtr interface{}) db.Func
	UpdateOp func(entityPtr interface{}) db.Func
	DeleteOp func(entityPtr interface{}) db.Func
}

func (s *Storage) Find(entityArrPtr interface{}, params map[string]interface{}) (err error) {
	s.Actor.ActionChan <- s.FindOp(entityArrPtr, params)
	return nil
}

func (s *Storage) Create(entityPtr interface{}) (err error) {
	s.Actor.ActionChan <- s.CreateOp(entityPtr)
	return nil
}
func (s *Storage) Update(entityPtr interface{}) (err error) {
	s.Actor.ActionChan <- s.UpdateOp(entityPtr)
	return nil
}
func (s *Storage) Delete(entityPtr interface{}) error {
	s.Actor.ActionChan <- s.DeleteOp(entityPtr)
	return nil
}
