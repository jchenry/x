package rest

import (
	"net/http"
	"net/url"
	"path/filepath"
	"sync"

	"github.com/jchenry/x/encoding"
	"github.com/jchenry/x/log"
)

// Example: Resource(p, c, JSONEncoder, json.Decode(func()interface{}{return &foo{}}), log.None{})
func Resource(p *sync.Pool, g Gateway, e EntityEncoder, d encoding.Decoder, l log.Logger) http.HandlerFunc {
	return restVerbHandler(
		GetResource(g, e, l),
		PostResource(g, d, p, l),
		PutResource(g, e, d, p, l),
		DeleteResource(g, l),
	)
}

func GetResource(store Readable, encode EntityEncoder, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // GET
		if id := filepath.Base(r.URL.Path); id != "" {
			if e, err := store.Read(id); err == nil { // handle individual entity
				encode(w, e)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error: %s", err)
			}
		} else {
			if params, err := url.ParseQuery(r.URL.RawQuery); err == nil {
				if e, err := store.All(params); err == nil { // handle all entities
					encode(w, e)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("Error: %s", err)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}

func PostResource(store Creatable, decode encoding.Decoder, pool *sync.Pool, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // POST TODO
		e := pool.Get()
		defer pool.Put(e)
		if err := decode(r.Body, e); err == nil {
			if err = store.Create(e); err == nil {
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

func PutResource(store Updatable, encode EntityEncoder, decode encoding.Decoder, pool *sync.Pool, log log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // PUT TODO
		e := pool.Get()
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
		}
	}
}

func DeleteResource(store Deletable, log log.Logger) http.HandlerFunc {
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
