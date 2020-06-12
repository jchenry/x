package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const resultFmt = "<a href=/wiki/%s>%s</a> . . . . . .  %s<br>\n"

func search(keyword string, w http.ResponseWriter, r *http.Request) (err error) {
	var results string
	if files, err := ioutil.ReadDir(*pageDir); err == nil {
		re := regexp.MustCompile(keyword)
		for _, f := range files {
			if f.Name() == keyword {
				results += fmt.Sprintf(resultFmt, f.Name(), f.Name(), f.Name())
			}
			if body, err := getFile(f.Name(), os.O_RDWR); err == nil {
				for _, occur := range re.FindSubmatch(body) {
					results += fmt.Sprintf(resultFmt, f.Name(), f.Name(), occur)
				}
			} else {
				return err
			}
		}
		render("search", "view", []byte(results), w)
	}
	return err
}
