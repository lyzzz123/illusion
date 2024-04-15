package config

import (
	"github.com/lyzzz123/illusion/lifecycle"
	"reflect"
)

type AfterObjectDestroyRunner struct {
}

func (afterObjectDestroyRunner *AfterObjectDestroyRunner) AfterContainerDestroyAction(objectContainer map[reflect.Type]interface{}) error {
	for _, registerObject := range objectContainer {
		if afterObjectDestroyObject, ok := registerObject.(lifecycle.AfterObjectDestroy); ok {
			if err := afterObjectDestroyObject.AfterObjectDestroyAction(); err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func (afterObjectDestroyRunner *AfterObjectDestroyRunner) GetPriority() int {
	return 1
}
