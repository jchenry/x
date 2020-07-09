package rest

import (
	"net/http"

	"github.com/jchenry/x/encoding"
	"github.com/jchenry/x/encoding/json"
	"github.com/jchenry/x/encoding/xml"
)

type EntityEncoder func(w http.ResponseWriter, e interface{})

func JSONEncoder(w http.ResponseWriter, e interface{}) error {
	return EntityResponseEncoder(w, "application/json", json.Encoder, e)
}

func XMLEncoder(w http.ResponseWriter, e interface{}) error {
	return EntityResponseEncoder(w, "application/xml", xml.Encoder, e)
}

func EntityResponseEncoder(w http.ResponseWriter, contentType string, encoder encoding.Encoder, e interface{}) error {
	w.Header().Set("content-type", contentType)
	return encoder(w, e)
}

func ErrorResponseEncoder(w http.ResponseWriter, contentType string, encoder encoding.Encoder, status int, err error) error {
	w.WriteHeader(status)
	return EntityResponseEncoder(w, contentType, encoder, map[string]interface{}{
		"status":  status,
		"message": err.Error,
	})
}
