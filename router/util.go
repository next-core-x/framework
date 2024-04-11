package router

import (
	"reflect"
	"runtime"
)

// nameOfFunction returns the name of the function
func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
