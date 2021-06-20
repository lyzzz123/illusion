package test

import (
	"fmt"
)

type TestListener struct {
}

func (testListener *TestListener) ServiceStartup() error {
	//delete(requestconverter.ParamConverterMap,"GetConverter")
	//delete(responsewriter.ResponseWriterMap,"GetConverter")
	//delete(converter.ParamConverterMap, reflect.Int)
	fmt.Println("ServiceStartup listener")
	return nil
}

func (testListener *TestListener) ServiceShutdown() error {
	//delete(requestconverter.ParamConverterMap,"GetConverter")
	//delete(responsewriter.ResponseWriterMap,"GetConverter")
	//delete(converter.ParamConverterMap, reflect.Int)
	fmt.Println("ServiceShutdown listener")
	return nil
}

func (testListener *TestListener) GetPriority() int {
	return 1
}
