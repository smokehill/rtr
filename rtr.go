package rtr

import (
	"strings"
	"regexp"
	"net/http"
	"net/url"
)

// Router registers routes.
type Router struct {
	routes []*Route
}

// Route stores information about set route.
type Route struct {
	// Request method [GET, POST, ...].
	Method string
	// URL regex pattern.
	Pattern string
	// Route handler function.
	Handler http.HandlerFunc
}

// NewRouter returns a new router instance.
func NewRouter() *Router {
	return &Router{}
}

// SetRoute returns a new router instance.
func (r *Router) SetRoute(method, pattern string, handler http.HandlerFunc) *Route {
	route := &Route{method, "^" + pattern + "$", handler}
	r.routes = append(r.routes, route)

	return route
}

// ServeHTTP dispatches the handler registered in set routes.
func (r *Router) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	match, allow := false, false

	for _, route := range r.routes {
		match, _ = regexp.MatchString(route.Pattern, rq.URL.Path)
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

// Helper functions
// -----------------------------------------------------------------------------

// SplitURL splits url into array of parts.
func SplitURL(url *url.URL) []string {
	return strings.Split(strings.Trim(url.String(), "/"), "/")
}