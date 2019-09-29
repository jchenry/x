package auth

import (
	"net/http"

	jchenry_http "github.com/jchenry/jchenry/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jchenry_http.RenderTemplate(w, "user", session.Values["profile"])
}
