package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/jchenry/jchenry/auth"
	jch_http "github.com/jchenry/jchenry/http"
	"github.com/jchenry/jchenry/payments"
)

func main() {
	auth.Init()
	StartServer()
}

func StartServer() {
	auth.PrintConfig()
	payments.PrintConfig()
	s := jch_http.NewServer(negroni.New()).
		Static("/public/*filepath", http.Dir("public/")).
		Service("", auth.Service(auth.FromEnv())).
		Service("", payments.Service(payments.FromEnv())).
		GET("/", "", http.HandlerFunc(HomeHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	s.Run(":" + port)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	jch_http.RenderTemplate(w, "home", nil)
}
