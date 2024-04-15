package test

import (
	"fmt"
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusion/converter"
	"reflect"
	"testing"
)

type ContainerLifeCycleTest struct {
}

func (containerLifeCycleTest *ContainerLifeCycleTest) BeforeContainerInitPropertyAction() error {
	fmt.Println("BeforeContainerInitPropertyAction")
	return nil
}
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInitPropertyAction(propertiesArray []map[string]string) error {
	fmt.Println("AfterContainerInitPropertyAction")
	return nil
}
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInitConverterAction(typeConverterMap map[reflect.Type]converter.Converter) error {
	fmt.Println("AfterContainerInitConverterAction")
	return nil
}
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInjectAction(objectContainer map[reflect.Type]interface{}) error {
	fmt.Println("AfterContainerInjectAction")
	return nil
}
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterRunAction(objectContainer map[reflect.Type]interface{}) error {
	fmt.Println("AfterRunAction")
	return nil
}

func TestContainerLifeCycle(t *testing.T) {

	illusion.Register(&ContainerLifeCycleTest{})
	illusion.Start()
}
