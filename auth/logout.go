package auth

import (
	"net/http"
	"net/url"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	domain := "dev-pb4s8m55.auth0.com"

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		panic(err.Error())
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost:3000")
	parameters.Add("client_id", "ae1e02bTwXA35O3r3Xxk4kbRf31j5ge9")
	Url.RawQuery = parameters.Encode()

	http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
}
