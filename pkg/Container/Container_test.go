package Container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Demo struct {
	xxx string
}

func TestNew(t *testing.T) {
	di := new()
	assert.Equal(t, di, &Container{
		dependencies: make(map[string]interface{}),
	})
}

func TestRegister(t *testing.T) {
	di := new()
	c := &Container{
		dependencies: make(map[string]interface{}),
	}
	di.Register("test", Demo{xxx: string("xxx")})
	c.dependencies["test"] = Demo{xxx: string("xxx")}
	assert.Equal(t, di, c)
}

func TestGet(t *testing.T) {
	di := new()
	demo := Demo{xxx: string("xxx")}
	di.Register("test", demo)
	assert.Equal(t, di.Get("test"), demo)
}
