package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type actionFunc func(s string, w http.ResponseWriter, r *http.Request) error

var pageDir *string

func main() {
	p, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	pageDir = flag.String("pageDir", path.Join(p, "pages"), "the directory in which pages exist")
	httpAddr := flag.String("http", "127.0.0.1:8080", " HTTP service address")
	help := flag.Bool("help", false, "this help.")
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	for path, action := range map[string]actionFunc{"/wiki/": view, "/edit/": edit, "/save/": save, "/search/": search} {
		register(path, action)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { view("WelcomeVisitors", w, r) })))
	log.Printf("using log/pass: %s/%s", os.Getenv("WIKI_USERNAME"), os.Getenv("WIKI_PASSWORD"))
	log.Printf("wiki has started listening at %s", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

func register(path string, action actionFunc) {
	http.Handle(path, http.StripPrefix(path, auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "" {
			if err := action(r.URL.Path, w, r); err != nil {
				log.Fatal(err)
			}
		}
	}))))
}
