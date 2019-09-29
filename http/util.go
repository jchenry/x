package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	HeaderContentType = "Content-Type"
	MimeJSON          = "application/json"
)

//ErrNotFound is returned when an entity could not be found in Find
var ErrNotFound = errors.New("entity not found")

func WriteEntity(w http.ResponseWriter, entityPtr interface{}) error {
	w.Header().Set(HeaderContentType, MimeJSON)

	var out []byte
	var err error

	if entityPtr == nil {
		return nil
	}

	out, err = json.MarshalIndent(entityPtr, " ", " ")

	if err != nil {
		return WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	if _, err = w.Write(out); err != nil {
		return err
	}
	return nil
}

func ReadEntity(request *http.Request, entityPtr interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(entityPtr)
	if err != nil {
		return err
	}

	return nil
}

// ErrorResponse - A structured Error HTTP response
type ErrorResponse struct {
	Status           int
	DeveloperMessage string
}

// WriteErrorResponse - Creates a new ErrorResponse and writes it to
// the response
func WriteErrorResponse(w http.ResponseWriter, status int, message string) error {
	log.Println(status, message)
	w.WriteHeader(status)
	return WriteEntity(w, ErrorResponse{status, message})
}
