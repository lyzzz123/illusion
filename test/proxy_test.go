package test

import (
	"fmt"
	"github.com/lyzzz123/illusion"
	"reflect"
	"testing"
)

type TestTargetInterface interface {
	PrintMessage()
}

type TestTarget struct {
}

func (testTarget *TestTarget) PrintMessage() {
	fmt.Println("this is proxy target")
}

type TestTargetProxy struct {
	Target interface{}
}

func (testTargetProxy *TestTargetProxy) PrintMessage() {
	targetProxy, _ := testTargetProxy.Target.(TestTargetInterface)
	fmt.Println("before target run")
	targetProxy.PrintMessage()
	fmt.Println("after target run")

}

type InjectObject struct {
	Target TestTargetInterface `require:"true"`
}

func (testTargetProxy *TestTargetProxy) SupportInterface() reflect.Type {
	return reflect.TypeOf(*new(TestTargetInterface))
}
func (testTargetProxy *TestTargetProxy) SetTarget(target interface{}) {
	testTargetProxy.Target = target
}

func TestProxy(t *testing.T) {
	injectObject := &InjectObject{}
	illusion.Register(&TestTarget{})
	illusion.Register(&TestTargetProxy{})
	illusion.Register(injectObject)
	illusion.Start()
	fmt.Println()
}
