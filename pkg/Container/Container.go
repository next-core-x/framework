package Container

import "fmt"

type Container struct {
	dependencies map[string]interface{}
}

func new() *Container {
	return &Container{
		dependencies: make(map[string]interface{}),
	}
}

func (app Container) Register(name string, dep interface{}) Container {
	app.dependencies[name] = dep
	return app
}

func (app Container) Get(name string) interface{} {
	return app.dependencies[name]
}

func main() {
	di := new()
	fmt.Println(di.dependencies)
}
