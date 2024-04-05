package illusioncontainer

import (
	"github.com/lyzzz123/illusion/container"
)

var mainContainer = container.MainContainer{}

func init() {
	mainContainer.InitContainer()
}

func Register(object interface{}) {
	mainContainer.Register(object)
}

func Start() {
	mainContainer.Start()
}
