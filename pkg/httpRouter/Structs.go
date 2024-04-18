package httpRouter

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler *Handler
}

type Handler struct {
	ServeHTTP func(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	routes []*Route
}
