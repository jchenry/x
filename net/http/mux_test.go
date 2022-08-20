package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tests = []struct {
	// RouteMethod  string
	RoutePattern string

	Method string
	Path   string
	Match  bool
	Params map[string]string
}{
	// simple path matching
	{
		"/one",
		"GET", "/one", true, nil,
	},
	{
		"/two",
		"GET", "/two", true, nil,
	},
	{
		"/three",
		"GET", "/three", true, nil,
	},
	// methods
	{
		"/methodcase",
		"GET", "/methodcase", true, nil,
	},
	{
		"/methodcase",
		"get", "/methodcase", true, nil,
	},
	{
		"/methodcase",
		"get", "/methodcase", true, nil,
	},
	{
		"/method1",
		"POST", "/method1", true, nil,
	},
	{
		"/method2",
		"GET", "/method2", true, nil,
	},
	{
		"/method3",
		"PUT", "/method3", true, nil,
	},
	// all methods
	{
		"/all-methods",
		"GET", "/all-methods", true, nil,
	},
	{
		"/all-methods",
		"POST", "/all-methods", true, nil,
	},
	{
		"/all-methods",
		"PUT", "/all-methods", true, nil,
	},
	// nested
	{
		"/parent/child/one",
		"GET", "/parent/child/one", true, nil,
	},
	{
		"/parent/child/two",
		"GET", "/parent/child/two", true, nil,
	},
	{
		"/parent/child/three",
		"GET", "/parent/child/three", true, nil,
	},
	// slashes
	{
		"slashes/one",
		"GET", "/slashes/one", true, nil,
	},
	{
		"/slashes/two",
		"GET", "slashes/two", true, nil,
	},
	{
		"slashes/three/",
		"GET", "/slashes/three", true, nil,
	},
	{
		"/slashes/four",
		"GET", "slashes/four/", true, nil,
	},
	// prefix
	{
		"/prefix/",
		"GET", "/prefix/anything/else", true, nil,
	},
	{
		"/not-prefix",
		"GET", "/not-prefix/anything/else", false, nil,
	},
	{
		"/prefixdots...",
		"GET", "/prefixdots/anything/else", true, nil,
	},
	{
		"/prefixdots...",
		"GET", "/prefixdots", true, nil,
	},
	// path params
	{
		"/path-param/{id}",
		"GET", "/path-param/123", true, map[string]string{"id": "123"},
	},
	{
		"/path-params/{era}/{group}/{member}",
		"GET", "/path-params/60s/beatles/lennon", true, map[string]string{
			"era":    "60s",
			"group":  "beatles",
			"member": "lennon",
		},
	},
	{
		"/path-params-prefix/{era}/{group}/{member}/",
		"GET", "/path-params-prefix/60s/beatles/lennon/yoko", true, map[string]string{
			"era":    "60s",
			"group":  "beatles",
			"member": "lennon",
		},
	},
	// misc no matches
	{
		"/not/enough",
		"GET", "/not/enough/items", false, nil,
	},
	{
		"/not/enough/items",
		"GET", "/not/enough", false, nil,
	},
}

func TestWay(t *testing.T) {
	for _, test := range tests {
		r := NewServeMux()
		match := false
		var ctx context.Context
		r.Handle(test.RoutePattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			match = true
			ctx = r.Context()
		}))
		req, err := http.NewRequest(test.Method, test.Path, nil)
		if err != nil {
			t.Errorf("NewRequest: %s", err)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if match != test.Match {
			t.Errorf("expected match %v but was %v: %s %s", test.Match, match, test.Method, test.Path)
		}
		if len(test.Params) > 0 {
			for expK, expV := range test.Params {
				// check using helper
				actualValStr := Param(ctx, expK)
				if actualValStr != expV {
					t.Errorf("Param: context value %s expected \"%s\" but was \"%s\"", expK, expV, actualValStr)
				}
			}
		}
	}
}

func TestMultipleRoutesDifferentMethods(t *testing.T) {
	r := NewServeMux()
	var match string

	r.Handle("/route", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			match = "GET /route"
		case http.MethodDelete:
			match = "DELETE /route"
		case http.MethodPost:
			match = "POST /route"
		}
	}))

	r.Handle("/route", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		match = "GET /route"
	}))
	r.Handle("/route", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		match = "DELETE /route"
	}))
	r.Handle("/route", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		match = "POST /route"
	}))

	req, err := http.NewRequest(http.MethodGet, "/route", nil)
	if err != nil {
		t.Errorf("NewRequest: %s", err)
	}
	r.ServeHTTP(httptest.NewRecorder(), req)
	if match != "GET /route" {
		t.Errorf("unexpected: %s", match)
	}

	req, err = http.NewRequest(http.MethodDelete, "/route", nil)
	if err != nil {
		t.Errorf("NewRequest: %s", err)
	}
	r.ServeHTTP(httptest.NewRecorder(), req)
	if match != "DELETE /route" {
		t.Errorf("unexpected: %s", match)
	}

	req, err = http.NewRequest(http.MethodPost, "/route", nil)
	if err != nil {
		t.Errorf("NewRequest: %s", err)
	}
	r.ServeHTTP(httptest.NewRecorder(), req)
	if match != "POST /route" {
		t.Errorf("unexpected: %s", match)
	}

}
