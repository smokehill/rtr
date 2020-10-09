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
	// Request methods ["GET", "POST", ...].
	methods []string
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
func (r *Router) SetRoute(methods string, pattern string, handler http.HandlerFunc) *Route {
	route := &Route{prepareRouteMethods(methods), "^" + pattern + "$", handler}
	r.routes = append(r.routes, route)

	return route
}

// HasMethod checks if route contains method.
func (r *Route) HasMethod(method string) bool {
	for _, m := range r.methods {
		if m == method {
			return true
		}
	}
	return false
}

// ServeHTTP dispatches the handler registered in set routes.
func (r *Router) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	match, allow := false, false

	for _, route := range r.routes {
		match, _ = regexp.MatchString(route.Pattern, rq.URL.Path)
		if route.HasMethod(rq.Method) == true && match == true {
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

// prepareRouteMethods prepares methods from string to array of values.
// "GET, POST, ..." => ["GET", "POST", ...]
func prepareRouteMethods(methods string) []string {
	m := methods
	m = strings.Replace(m, " ", "", -1)
	m = strings.TrimLeft(m, ",")
	m = strings.TrimRight(m, ",")

	return strings.Split(m, ",")
}