package httpRouter

import "net/http"

type Router struct {
	routes []*Route
}

type Handler struct {
	ServeHTTP func(w http.ResponseWriter, r *http.Request)
}

func New() *Router {
	return &Router{}
}

func (r *Router) Register(method, path string, handler http.Handler) {
	route := &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	r.routes = append(r.routes, route)
}

func (r *Router) Match(req *http.Request) *Route {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			return route
		}
	}
	return nil
}
