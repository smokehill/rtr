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
		t.Errorf("r.routes result must be equal to 2, but got %d", len(r.routes))
	}
}

func TestSplitURL(t *testing.T) {
	u, _ := url.Parse("/news/1")
	r := SplitURL(u)

	if len(r) != 2 {
		t.Errorf("SplitURL(u) result must be equal to 2, buy got %d", len(r))
	}
}