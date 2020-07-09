package rss

import (
	"io"

	"golang.org/x/net/html"
)

func Parse(r io.Reader) *Feed {

	// z := html.NewTokenizer(r)

	return nil
}

func parseFeed(z html.Tokenizer) *Feed {
	z.Next()
	return nil
}

func parseVersion(z html.Tokenizer) {}
