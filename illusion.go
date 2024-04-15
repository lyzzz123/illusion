package illusion

import (
	"github.com/lyzzz123/illusion/config"
	"github.com/lyzzz123/illusion/container"
	"github.com/lyzzz123/illusion/log"
)

var mainContainer = container.MainContainer{}

func init() {
	mainContainer.InitContainer()
	mainContainer.Register(&log.DefaultLog{})
	mainContainer.Register(&config.AfterObjectDestroyRunner{})
	mainContainer.Register(&config.AfterObjectInjectRunner{})
}

func Register(object interface{}) {
	mainContainer.Register(object)
}

func Start() {
	mainContainer.Start()
}

func TestStart() {
	mainContainer.TestStart()
}
