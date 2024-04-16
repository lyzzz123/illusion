package test

import (
	"fmt"
	"github.com/lyzzz123/illusion"
	"testing"
)

type TestA struct {
}

type TestB struct {
	MTestA *TestA `require:"true"`
}

func TestObjectInject(t *testing.T) {
	testA := &TestA{}
	testB := &TestB{}
	illusion.Register(testA)
	illusion.Register(testB)
	illusion.Start()
	fmt.Println()
}
