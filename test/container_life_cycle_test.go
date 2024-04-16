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

//在整个程序启动前执行，例如可以在这个扩展点输出一些banner
func (containerLifeCycleTest *ContainerLifeCycleTest) BeforeContainerInitPropertyAction() error {
	fmt.Println("BeforeContainerInitPropertyAction")
	return nil
}

//属性加载完成之后，可以在这个扩展点对加载完的属性做一些操作
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInitPropertyAction(propertiesArray []map[string]string) error {
	fmt.Println("AfterContainerInitPropertyAction")
	return nil
}

//数据类型转换器加载完成之后，可以在这里加载一些自定义的类型转换器
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInitConverterAction(typeConverterMap map[reflect.Type]converter.Converter) error {
	fmt.Println("AfterContainerInitConverterAction")
	return nil
}

//容器中的所有对象的属性都注入完成之后，可以在这里对托管对象做一些自定义操作
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInjectAction(objectContainer map[reflect.Type]interface{}) error {
	fmt.Println("AfterContainerInjectAction")
	return nil
}

//容器启动的最后一个扩展点，可以在这里集成一些其他的程序框架，比如illusionmvc
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterRunAction(objectContainer map[reflect.Type]interface{}) error {
	fmt.Println("AfterRunAction")
	return nil
}
func (containerLifeCycleTest *ContainerLifeCycleTest) GetPriority() int {
	return 1
}

func TestContainerLifeCycle(t *testing.T) {

	illusion.Register(&ContainerLifeCycleTest{})
	illusion.Start()
}
