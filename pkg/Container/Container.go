package Container

type Container struct {
	dependencies map[string]interface{}
}

func (app Container) Register(name string, dep interface{}) Container {
	app.dependencies[name] = dep
	return app
}

func (app Container) Get(name string) interface{} {
	return app.dependencies[name]
}

func (app Container) remove(name string) Container {
	delete(app.dependencies, name)
	return app
}
