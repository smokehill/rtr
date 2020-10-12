package rtr

import (
	"testing"
	"net/url"
	"net/http"
)

func TestRoutes(t *testing.T) {
	r := NewRouter()
	r.SetRoute("GET", "/news", func (w http.ResponseWriter, r *http.Request) {})
	r.SetRoute("GET", "/news/([0-9]+)", func (w http.ResponseWriter, r *http.Request) {})

	if len(r.routes) != 2 {
		t.Errorf("len(r.routes) must be equal to 2, but got %d", len(r.routes))
	}
}

func TestSplitURL(t *testing.T) {
	u, _ := url.Parse("/news/1")
	r := SplitURL(u)

	if len(r) != 2 {
		t.Errorf("len(r) must be equal to 2, but got %d", len(r))
	}
}

func TestPrepareRouteMethods(t *testing.T) {
	methods := "GET,POST,PUT,DELETE"
	r := prepareRouteMethods(methods)

	if len(r) != 4 {
		t.Errorf("len(r) must be equal to 4, but got %d", len(r))
	} else {
		if r[0] != "GET" {
			t.Errorf("r[0] must be equal to \"GET\", but got %s", r[0])
		} else if r[1] != "POST" {
			t.Errorf("r[1] must be equal to \"POST\", but got %s", r[1])
		} else if r[2] != "PUT" {
			t.Errorf("r[2] must be equal to \"PUT\", but got %s", r[2])
		} else if r[3] != "DELETE" {
			t.Errorf("r[3] must be equal to \"DELETE\", but got %s", r[3])
		}
	}
}