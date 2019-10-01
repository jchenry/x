package payments

import (
	"net/http"

	"github.com/jchenry/jchenry/auth"
)

func HasTenantAndSubscription(productID string) func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		session, err := auth.Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if u, ok := session.Values["profile"]; ok {
			user := u.(auth.User)
			if _, exist := user.AppMetadata.Apps[productID]; !exist {
				next(w, r)
			} else {
				http.Redirect(w, r, "/subscription", http.StatusSeeOther)
			}
		}
	}
}
