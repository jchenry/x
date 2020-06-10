package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	s := server{
		router: http.NewServeMux(),
	}

	s.routes()

	if db, err := sql.Open("sqlite3", "foo.db"); err == nil {
		s.db = db
	} else {
		return err
	}
	return http.ListenAndServe(":8080", s.router)

}
