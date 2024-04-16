package Container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Demo struct {
	xxx string
}

func di() *Container {
	return &Container{
		dependencies: make(map[string]interface{}),
	}
}

func TestRegister(t *testing.T) {
	di := di()
	c := &Container{
		dependencies: make(map[string]interface{}),
	}
	di.Register("test", Demo{xxx: string("xxx")})
	c.dependencies["test"] = Demo{xxx: string("xxx")}
	assert.Equal(t, di, c)
}

func TestGet(t *testing.T) {
	di := di()
	demo := Demo{xxx: string("xxx")}
	di.Register("test", demo)
	assert.Equal(t, di.Get("test"), demo)
}

func TestRemove(t *testing.T) {
	di := di()
	demo := Demo{xxx: string("xxx")}
	c := &Container{
		dependencies: make(map[string]interface{}),
	}
	di.Register("test", demo)
	assert.Equal(t, di.Get("test"), demo)
	di.remove("test")
	assert.Nil(t, di.Get("test"))
	assert.Equal(t, di, c)
}
