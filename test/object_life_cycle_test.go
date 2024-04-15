package test

import (
	"fmt"
	"github.com/lyzzz123/illusion"
	"testing"
)

type ObjectLifeCycleTest struct {
}

func (objectLifeCycleTest *ObjectLifeCycleTest) AfterObjectInjectAction() error {
	fmt.Println("AfterObjectInjectAction")
	return nil
}

func (objectLifeCycleTest *ObjectLifeCycleTest) AfterObjectDestroyAction() error {
	fmt.Println("AfterObjectDestroyAction")
	return nil
}

func TestObjectLifeCycle(t *testing.T) {

	illusion.Register(&ObjectLifeCycleTest{})
	illusion.Start()
}
