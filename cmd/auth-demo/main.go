package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/jchenry/jchenry/internal/auth"
	_http "github.com/jchenry/jchenry/internal/http"
)

func main() {
	auth.Init()
	StartServer()
}

func StartServer() {
	auth.PrintConfig()
	s := _http.NewServer(negroni.New(), _http.NewJulienschmidtHTTPRouter()).
		Static("/public/*filepath", http.Dir("public/")).
		Service("", auth.Service(auth.FromEnv())).
		Get("/", "", http.HandlerFunc(HomeHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	s.Run(":" + port)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_http.RenderTemplate(w, "home", nil)
}
