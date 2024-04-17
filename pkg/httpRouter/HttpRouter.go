package httpRouter

import "net/http"

func New() *Router {
	return &Router{}
}

func (r *Router) Register(method, path string, handler Handler) *Router {
	route := &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	r.routes = append(r.routes, route)
	return r
}

func (r *Router) Match(req *http.Request) *Route {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			return route
		}
	}
	return nil
}
