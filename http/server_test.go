package http_test

import (
	"os"

	"github.com/jchenry/jchenry/http"
	"github.com/jchenry/jchenry/rest"
)

func ExampleServer() {
	type contact struct {
		ID    int64  `json:"id"`
		First string `json:"firstName"`
		Last  string `json:"lastName"`
		Email string `json:"emailAddress"`
	}

	s := http.NewServer().
		Service("", rest.Collection(new(contact), crud.NewInMemoryCrudService()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s.Run(":" + port)
}
