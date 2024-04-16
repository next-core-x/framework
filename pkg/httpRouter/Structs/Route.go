package httpRouter

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}
