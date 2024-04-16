package test

import (
	"fmt"
	"github.com/lyzzz123/illusion"
	"testing"
)

type TestInjectInterface interface {
	PrintMessage()
}

type TestAA struct {
}

func (testAA *TestAA) PrintMessage() {
	fmt.Println("this is test a")
}

type TestBB struct {
	MTestA TestInjectInterface `require:"true"`
}

func TestInterfaceInject(t *testing.T) {
	testA := &TestAA{}
	testB := &TestBB{}
	illusion.Register(testA)
	illusion.Register(testB)
	illusion.Start()
	fmt.Println()
}
