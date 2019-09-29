package auth

import (
	"context"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
)

func NewCallbackHandler(c Config) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.URL.Query().Get("state") != session.Values["state"] {
			http.Error(w, "Invalid state parameter", http.StatusBadRequest)
			return
		}

		authenticator, err := NewAuthenticator(c.Domain, c.ClientID, c.ClientSecret, c.CallbackURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
		if err != nil {
			log.Printf("no token found: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}

		oidcConfig := &oidc.Config{
			ClientID: c.ClientID,
		}

		idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)

		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Getting now the userInfo
		user := User{}

		// var profile map[string]interface{}
		if err := idToken.Claims(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["id_token"] = rawIDToken
		session.Values["access_token"] = token.AccessToken
		session.Values["profile"] = user
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// if application ID is non existent, and therefore does not have a tenant
		// Create or associate?
		// Create:
		//  - Create Tenant
		//    - Specify plan
		//    - Specify payment info
		//	- Associate Tenant
		//    - by email address domain?
		//set tenant ID on application ID in App Metadata on user

		if c.CallbackFunc != nil {
			c.CallbackFunc(c, user)
		} else {
			// Redirect to logged in page
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		}

	}
}
