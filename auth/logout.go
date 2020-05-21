package auth

import (
	"net/http"
	"net/url"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	if cook, err := r.Cookie(SessionName); err == nil {
		cook.MaxAge = -1
		http.SetCookie(w, cook)
	}

	domain := "dev-pb4s8m55.auth0.com"

	// var Url *url.URL
	URL, err := url.Parse("https://" + domain)

	if err != nil {
		panic(err.Error())
	}

	session, err := Store.Get(r, SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	URL.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost:3000")
	parameters.Add("client_id", "ae1e02bTwXA35O3r3Xxk4kbRf31j5ge9")
	URL.RawQuery = parameters.Encode()

	http.Redirect(w, r, URL.String(), http.StatusTemporaryRedirect)
}
