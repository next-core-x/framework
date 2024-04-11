package router

import "github.com/gin-gonic/gin"

// Create a new Structure CustomGroup for implementing Gin's RouterGroup
type CustomGroup struct {
	group  *gin.RouterGroup
	router *CustomRouter
}

// Create a Route for handling the Request
func (g *CustomGroup) Handle(method, path string, handlers ...gin.HandlerFunc) *CustomGroup {
	g.group.Handle(method, path, handlers...)
	g.router.routes = append(g.router.routes, RouteInfo{
		Method:  method,
		Path:    path,
		Handler: nameOfFunction(handlers[0]),
	})
	return g
}

// Create a Sub-Group
func (g *CustomGroup) Group(path string, handlers ...gin.HandlerFunc) *CustomGroup {
	return &CustomGroup{
		group:  g.group.Group(path, handlers...),
		router: g.router,
	}
}
