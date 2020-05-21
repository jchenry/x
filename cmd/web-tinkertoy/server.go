package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"

	"rsc.io/dbstore"
)

type server struct {
	db          *sql.DB
	router      *http.ServeMux
	echoMessage string
}

func (s *server) routes() {
	s.router.HandleFunc("/time", s.handleTime())
	s.router.HandleFunc("/echo", s.handleEcho())
	s.router.HandleFunc("/fortune", s.handleFortune())

}

func (s *server) handleTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, time.Now().String())
	}
}

func (s *server) handleEcho() http.HandlerFunc {
	var (
		init   sync.Once
		tmpl   *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tmpl, tplerr = template.ParseFiles("echoform.tmpl")
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case http.MethodPost:
			r.ParseForm()
			s.echoMessage = r.Form.Get("msg")
			http.Redirect(w, r, "/echo", 301)
			return
		case http.MethodGet:
			if err := tmpl.Execute(w, map[string]string{"Message": s.echoMessage}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func (s *server) handleFortune() http.HandlerFunc {
	var (
		init    sync.Once
		dba     *DBActor
		storage *dbstore.Storage
		tmpl    *template.Template
		tplerr  error
	)

	type fortuneWrapper struct {
		Fortune string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			ctx, _ := context.WithCancel(context.Background())
			dba = &DBActor{
				DB:         s.db,
				ActionChan: make(chan DBFunc),
			}
			go dba.Run(ctx)
			storage = new(dbstore.Storage)
			storage.Register(&fortuneWrapper{})
			err := storage.CreateTables(dba.DB)
			if err != nil {
				fmt.Println(err)
			}
			tmpl, tplerr = template.ParseFiles("fortunesupload.tmpl")
		})

		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case http.MethodPost:
			r.ParseMultipartForm(10 << 20)
			file, _, err := r.FormFile("fortunes")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var fortune string
			for scanner.Scan() {
				switch scanner.Text() {
				case "%":
					dba.ActionChan <- DBStoreInsert(storage, &fortuneWrapper{Fortune: fortune})
				default:
					fortune = scanner.Text()
				}

			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		case http.MethodGet:
			f := fortuneWrapper{}
			DBStoreSelect(storage, &f, "ORDER BY RANDOM() LIMIT 1", "*")(s.db)
			if err := tmpl.Execute(w, map[string]string{"Message": f.Fortune}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	}
}
