package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"

	"github.com/jchenry/jchenry/db"
	"rsc.io/dbstore"
)

func (s *server) handleFortune() http.HandlerFunc {
	var (
		init    sync.Once
		dba     *db.Actor
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
			dba = &db.Actor{
				DB:         s.db,
				ActionChan: make(chan db.Func),
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
			results := make(chan interface{})
			err := make(chan error)

			dba.ActionChan <- DBStoreSelect(storage, err, results, &f, "ORDER BY RANDOM() LIMIT 1", "*")
			select {
			case r := <-results:
				if err := tmpl.Execute(w, map[string]string{"Message": r.(*fortuneWrapper).Fortune}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			case e := <-err:
				http.Error(w, e.Error(), http.StatusInternalServerError)
			}

		}

	}
}
