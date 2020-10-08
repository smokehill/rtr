package rtr

import (
	// "fmt"
	"regexp"
	"net/http"
)

type Router struct {
	routes []*Route
}

type Route struct {
	Method string
	Scheme string
	Handler http.HandlerFunc
}

func (r *Router) SetRoute(method, scheme string, handler http.HandlerFunc) *Route {
	route := &Route{method, "^" + scheme + "$", handler}
	r.routes = append(r.routes, route)

	return route
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	match, allow := false, false

	for _, route := range r.routes {
		match, _ = regexp.MatchString(route.Scheme, rq.URL.Path)
		if rq.Method == route.Method && match == true {
			allow = true
			route.Handler(rw, rq)
			return
		}
	}

	if match == true && allow == false {
		http.Error(rw, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(rw, rq)
}