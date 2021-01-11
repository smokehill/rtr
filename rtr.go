package rtr

import (
	"strings"
	"regexp"
	"net/http"
	"net/url"
)

type Router struct {
	routes []*Route
}

type Route struct {
	methods []string          // request methods ["GET", "POST", ...]
	Pattern string            // URL regex pattern
	Handler http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) SetRoute(methods string, pattern string, handler http.HandlerFunc) *Route {
	route := &Route{prepareRouteMethods(methods), "^" + pattern + "$", handler}
	r.routes = append(r.routes, route)

	return route
}

func (r *Route) hasMethod(method string) bool {
	for _, m := range r.methods {
		if m == method {
			return true
		}
	}
	return false
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	match, allow := false, false

	for _, route := range r.routes {
		match, _ = regexp.MatchString(route.Pattern, rq.URL.Path)
		if route.hasMethod(rq.Method) == true && match == true {
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

func SplitURL(url *url.URL) []string {
	return strings.Split(strings.Trim(url.String(), "/"), "/")
}

// Prepares methods from string to array of values.
// "GET, POST, ..." => ["GET", "POST", ...]
func prepareRouteMethods(methods string) []string {
	m := methods
	m = strings.Replace(m, " ", "", -1)
	m = strings.TrimLeft(m, ",")
	m = strings.TrimRight(m, ",")

	return strings.Split(m, ",")
}