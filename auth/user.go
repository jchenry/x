package auth

import (
	"net/http"

	jchenry_http "github.com/jchenry/jchenry/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := Store.Get(r, SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jchenry_http.RenderTemplate(w, "user", session.Values["profile"])
}

type User struct {
	ID        string                 `json:"sub"`
	Email     string                 `json:"email"`
	FirstName string                 `json:"given_name"`
	LastName  string                 `json:"family_name"`
	Picture   string                 `json:"picture"`
	Nickname  string                 `json:"nickname"`
	Apps      map[string]interface{} `json:"app_metadata,omitempty"`

	//UserMetadata UserMetadata `json:"user_metadata"`
}

// type UserMetadata struct {
// }
