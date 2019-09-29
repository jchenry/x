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

type User struct {
	Email       string      `json:"email"`
	FirstName   string      `json:"given_name"`
	LastName    string      `json:"family_name"`
	Picture     string      `json:"picture"`
	Nickname    string      `json:"nickname"`
	AppMetadata AppMetadata `json:"app_metadata"`

	//UserMetadata UserMetadata `json:"user_metadata"`
}

type AppMetadata struct {
	Apps map[string]string // an association between the unique applicationID and the tenantID that the user is associated with
	// Apps []struct {
	// 	ApplicationID string
	// 	TenantID      string
	// }
}

// type UserMetadata struct {
// }
