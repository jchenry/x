package http

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func BasicAuth(h http.Handler, htpasswd map[string]string, realm string) http.HandlerFunc {
	rlm := fmt.Sprintf(`Basic realm="%s"`, realm)
	sha1 := func(password string) string {
		s := sha1.New()
		_, _ = s.Write([]byte(password))
		passwordSum := []byte(s.Sum(nil))
		return base64.StdEncoding.EncodeToString(passwordSum)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if pw, ok := htpasswd[user]; !ok || !strings.EqualFold(pass, sha1(pw)) {
			w.Header().Set("WWW-Authenticate", rlm)
			http.Error(w, "Unauthorized", 401)
			return
		}
		h.ServeHTTP(w, r)
	}
}
