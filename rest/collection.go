package rest

import (
	"net/http"
	"net/url"
	"path/filepath"
	"sync"

	"github.com/jchenry/x/encoding"
	"github.com/jchenry/x/log"
)

type CollectionStore interface {
	All(params map[string][]string) (interface{}, error)
	Get(id string) (interface{}, error)
	Delete(id string) error
	Update(e interface{}) error
	New(e interface{}) error
}

// Example: Collection(p, c, JSONEncoder, json.Decode(func()interface{}{return &foo{}}), log.None{})
func Collection(pool *sync.Pool, store CollectionStore, encode EntityEncoder, decode encoding.Decoder, log log.Logger) http.HandlerFunc {
	return EntityHandler(
		collectionGet(store, encode, log),
		collectionPost(store, encode, decode, pool, log),
		collectionPut(store, encode, decode, pool, log),
		collectionDelete(store, encode, log),
	)
}

func collectionGet(store CollectionStore, encode EntityEncoder, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // GET
		if id := filepath.Base(r.URL.Path); id != "" {
			if e, err := store.Get(id); err == nil { // handle individual entity
				encode(w, e)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				encode(w, err)
				log.Printf("Error: %s", err)
			}
		} else {
			if params, err := url.ParseQuery(r.URL.RawQuery); err == nil {
				if e, err := store.All(params); err == nil { // handle all entities
					encode(w, e)
				} else {
					// TODO: we really should write a header here, but need to figure out what it should be
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("Error: %s", err)
				}
			} else {
				//	encode(w, err)
				w.WriteHeader(http.StatusBadRequest)

			}
		}
	}
}

func collectionPost(store CollectionStore, encode EntityEncoder, decode encoding.Decoder, pool *sync.Pool, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // POST TODO
		e := pool.New()
		defer pool.Put(e)
		if err := decode(r.Body, e); err == nil {
			if err = store.New(e); err == nil {
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error: %s", err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func collectionPut(store CollectionStore, encode EntityEncoder, decode encoding.Decoder, pool *sync.Pool, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // PUT TODO
		e := pool.New()
		defer pool.Put(e)
		if err := decode(r.Body, e); err == nil {
			if err = store.Update(e); err == nil {
				w.WriteHeader(http.StatusAccepted)
				encode(w, e)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error: %s", err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, err)

		}
	}
}

func collectionDelete(store CollectionStore, encode EntityEncoder, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // DELETE TODO
		if id := filepath.Base(r.URL.Path); id != "" {
			if err := store.Delete(id); err == nil {
				w.WriteHeader(http.StatusNoContent)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error: %s", err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
