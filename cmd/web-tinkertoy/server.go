package main

import (
	"database/sql"
	"net/http"
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
