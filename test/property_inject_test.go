package test

import (
	"github.com/lyzzz123/illusion"
	"testing"
)

type PropertyInjectTest struct {
	Boolbool       bool     `property:"bool.bool"`
	Boolptr        *bool    `property:"bool.ptr, true"`
	Float32float32 float32  `property:"float32.float32, false"`
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
	Uintptrq     *uint   `property:"uint.ptrq"`

	IntSlice []int          `property:"int.slice, true"`
	IntMap   map[string]int `property:"int.map, false"`
}

func TestPropertyInject(t *testing.T) {
	propertyInjectTest := &PropertyInjectTest{}
	illusion.Register(propertyInjectTest)
	illusion.Start()
}
