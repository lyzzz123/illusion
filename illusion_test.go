package illusion

import (
	"fmt"
	"github.com/lyzzz123/illusion/converter"
	"reflect"
	"testing"
)

type TestValue struct {
	Boolbool       bool     `property:"bool.bool"`
	Boolptr        *bool    `property:"bool.ptr"`
	Float32float32 float32  `property:"float32.float32"`
	Float32ptr     *float32 `property:"float32.ptr"`
	Float64float64 float64  `property:"float64.float64"`
	Float64ptr     *float64 `property:"float64.ptr"`

	Int8int8     int8    `property:"int8.int8"`
	Int8ptr      *int8   `property:"int8.ptr"`
	Int16int16   int16   `property:"int16.int16"`
	Int16ptr     *int16  `property:"int16.ptr"`
	Int32int32   int32   `property:"int32.int32"`
	Int32ptr     *int32  `property:"int32.ptr"`
	Int64int64   int64   `property:"int64.int64"`
	Int64ptr     *int64  `property:"int64.ptr"`
	Intint       int     `property:"int.int"`
	Intptr       *int    `property:"int.ptr"`
	Stringstring string  `property:"string.string"`
	Stringptr    *string `property:"string.ptr"`
	Uint8uint8   uint8   `property:"uint8.uint8"`
	Uint8ptr     *uint8  `property:"uint8.ptr"`
	Uint16uint16 uint16  `property:"uint16.uint16"`
	Uint16ptr    *uint16 `property:"uint16.ptr"`
	Uint32uint32 uint32  `property:"uint32.uint32"`
	Uint32ptr    *uint32 `property:"uint32.ptr"`
	Uint64uint64 uint64  `property:"uint64.uint64"`
	Uint64ptr    *uint64 `property:"uint64.ptr"`
	Uintuint     uint    `property:"uint.uint"`
	Uintptr      *uint   `property:"uint.ptr"`

	IntSlice []int          `property:"int.slice"`
	IntMap   map[string]int `property:"int.map"`
}

func (testValue *TestValue) Hello() {

}

type TestValueInterface interface {
	Hello()
}

type TestInject struct {
	TestValueValue TestValueInterface `require:"true"`
}

type TestLifeCycle struct {
}

func (testLifeCycle *TestLifeCycle) BeforeInitPropertyAction() error {
	fmt.Println("BeforeInitPropertyAction")
	return nil
}
func (testLifeCycle *TestLifeCycle) AfterInitPropertyAction(propertiesArray []map[string]string) error {
	fmt.Println("AfterInitPropertyAction")
	return nil
}

func (testLifeCycle *TestLifeCycle) AfterInitConverterAction(typeConverterMap map[reflect.Type]converter.Converter) error {
	fmt.Println("AfterInitConverterAction")
	return nil
}
func (testLifeCycle *TestLifeCycle) AfterInitInjectAction(objectContainer map[reflect.Type]interface{}) error {
	fmt.Println("AfterInitInjectAction")
	return nil
}
func (testLifeCycle *TestLifeCycle) AfterRunAction(objectContainer map[reflect.Type]interface{}) error {
	fmt.Println("AfterRunAction")
	return nil
}

func TestToRegex(t *testing.T) {
	Register(&TestValue{})
	Register(&TestInject{})
	Register(&TestLifeCycle{})

	Start()

}
