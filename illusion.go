package illusion

import (
	"github.com/lyzzz123/illusion/container"
	"reflect"
)

var mainContainer = container.MainContainer{}

func init() {
	mainContainer.InitContainer()
}

func Register(object interface{}) {
	mainContainer.Register(object)
}

func GetObject(typ reflect.Type) interface{} {
	object, ok := mainContainer.ObjectContainer[typ]
	if ok {
		return object
	}
	return nil
}

func Start() {
	mainContainer.Start()
}
