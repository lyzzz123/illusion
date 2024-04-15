package config

import (
	"github.com/lyzzz123/illusion/lifecycle"
	"reflect"
)

type AfterObjectInjectRunner struct {
}

func (afterObjectInjectRunner *AfterObjectInjectRunner) AfterContainerInjectAction(objectContainer map[reflect.Type]interface{}) error {
	for _, registerObject := range objectContainer {
		if afterObjectInjectObject, ok := registerObject.(lifecycle.AfterObjectInject); ok {
			if err := afterObjectInjectObject.AfterObjectInjectAction(); err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func (afterObjectInjectRunner *AfterObjectInjectRunner) GetPriority() int {
	return 1
}
