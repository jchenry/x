package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
)

var patterns [5]*regexp.Regexp
var renderers [5]func([]byte) []byte

func init() {
	// /*autoLinkRegexp*/ patterns[0], renderers[0] = regexp.MustCompile("[A-Z][a-z0-9]+([A-Z][a-z0-9]+)+"), func(s []byte) []byte { return []byte(fmt.Sprintf(`<a href="/wiki/%s/">%s</a>`, string(s), string(s))) }
	/*BracketedAutoLinkRegexp*/
	patterns[0], renderers[0] = regexp.MustCompile("\\[\\[[A-Za-z0-9 ]+([A-Za-z0-9 ]+)+\\]\\]"), func(s []byte) []byte { return []byte(fmt.Sprintf(`<a href="/wiki/%s/">%s</a>`, string(s), string(s))) }

	/*searchRegexp*/
	patterns[1], renderers[1] = regexp.MustCompile("\\[Search\\]"), func(s []byte) []byte {
		return []byte(`<form id="search_form" action="/search" onsubmit="searchHelper()"><input type="text" size="40" name="search" value=""><input type="submit" value="Search"></form>`)
	}
	/*youTubeLinkRegexp*/ patterns[2], renderers[2] = regexp.MustCompile("https://(www.)?youtube.com/watch\\?v=([-\\w]+)"), func(s []byte) []byte {
		return []byte(fmt.Sprintf(`<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/%s?rel=0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>`, strings.Split(string(s), "=")[1]))
	}
	/*isbnLinkRegexp*/ patterns[3], renderers[3] = regexp.MustCompile("ISBN:*([0-9]{10,})"), func(s []byte) []byte {
		return []byte(fmt.Sprintf(`<a href="http://www.amazon.com/exec/obidos/ISBN=%s" rel="nofollow">ISBN %s</a>`, bytes.Replace(bytes.Split(s, []byte(":"))[1], []byte("-"), []byte(""), -1), bytes.Split(s, []byte(":"))[1]))
	}
	/*alltextRegexp*/ patterns[4], renderers[4] = regexp.MustCompile(".*"), func(s []byte) []byte {
		return blackfriday.Run(s, blackfriday.WithExtensions(blackfriday.CommonExtensions))
	}
}

func view(pageName string, w http.ResponseWriter, r *http.Request) (err error) {
	var body []byte
	if body, err = getFile(pageName, os.O_RDWR); os.IsNotExist(err) {
		http.Redirect(w, r, fmt.Sprintf("/edit/%s", pageName), http.StatusTemporaryRedirect) // no page? redirect to edit/create it.
		return nil
	}
	for i := range renderers {
		body = patterns[i].ReplaceAllFunc(body, renderers[i])
	}
	return render(pageName, "view", body, w)
}
