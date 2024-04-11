package router

// RouteInfo is a structure for storing the Route Information
type RouteInfo struct {
	Method  string // HTTP Method
	Path    string // URL Path
	Handler string // Handler Function Name
}
