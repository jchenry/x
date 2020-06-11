package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/rsc/rsc/google"
	"github.com/rsc/rsc/imap"
)

type Pair struct {
	Key   string
	Value uint64
}

type PairList []*Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	acct := google.Acct("")

	c, err := imap.NewClient(imap.TLS, "imap.gmail.com", acct.Email, acct.Password, "/")
	if err != nil {
		log.Fatal(err)
	}

	all := c.Inbox()
	// all := c.Box("[Gmail]/All Mail")

	if err := all.Check(); err != nil {
		log.Fatal(err)
	}

	msgs := all.Msgs()

	counts := make(map[string]*Pair)
	labels := make(map[string]struct{})
	var pairs PairList = make([]*Pair, 0)

	for _, m := range msgs {
		ls := m.GmailLabels
		for _, l := range ls {
			labels[l] = struct{}{}
		}

		email := m.Hdr.From[0].Email
		if _, ok := counts[email]; ok {
			counts[email].Value = counts[email].Value + 1
		} else {
			p := Pair{email, 1}
			pairs = append(pairs, &p)
			counts[email] = &p
		}
	}

	keys := []string{}

	for k := range counts {
		keys = append(keys, k)
	}

	sort.Sort(pairs)

	for _, e := range pairs {
		fmt.Printf("%d %s\t\t\thttps://mail.google.com/mail/u/0/#search/in%%3Ainbox+from%%3A(%s)\n", e.Value, e.Key, e.Key)
	}

	// for k, _ := range labels {
	// 	fmt.Printf("%s, ", k)
	// }
	fmt.Println()

}
