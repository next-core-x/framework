package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Create a new Structure CustomRouter for implementing Gin's Router
type CustomRouter struct {
	router *gin.Engine
	routes []RouteInfo
}

// Create a new CustomRouter and return the pointer to it
func NewCustomRouter() *CustomRouter {
	return &CustomRouter{
		router: gin.Default(),
		routes: make([]RouteInfo, 0),
	}
}

// Create a new Router Group
func (r *CustomRouter) Group(path string, handlers ...gin.HandlerFunc) *CustomGroup {
	return &CustomGroup{
		group:  r.router.Group(path, handlers...),
		router: r,
	}
}

// List the All Routes
func (r *CustomRouter) ListRoutes() {
	for _, route := range r.routes {
		fmt.Printf("%-7s %-50s %s\n", route.Method, route.Path, route.Handler)
	}
}

// Run the Router and Listen on the Address
func (r *CustomRouter) Run(addr ...string) error {
	return r.router.Run(addr...)
}
