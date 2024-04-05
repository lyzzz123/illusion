package container

import (
	"bufio"
	"fmt"
	"github.com/lyzzz123/illusion/converter"
	"github.com/lyzzz123/illusion/lifecycle"
	"github.com/lyzzz123/illusion/utils"
	"os"
	"reflect"
	"strings"
)

type MainContainer struct {
	ObjectContainer map[reflect.Type]interface{}

	PropertiesArray []map[string]string

	TypeConverterMap map[reflect.Type]converter.Converter

	BeforeInitPropertyArray []lifecycle.BeforeInitProperty

	AfterInitPropertyArray []lifecycle.AfterInitProperty

	AfterInitConverterArray []lifecycle.AfterInitConverter

	AfterInitInjectArray []lifecycle.AfterInitInject

	AfterRunArray []lifecycle.AfterRun
}

func (mainContainer *MainContainer) GetProperty(key string) string {
	for _, propertyMap := range mainContainer.PropertiesArray {
		value, ok := propertyMap[key]
		if !ok {
			continue
		}
		return value
	}
	return ""
}

func (mainContainer *MainContainer) GetPropertyMap(prefix string) map[string]string {
	returnMap := make(map[string]string)
	index := len(mainContainer.PropertiesArray) - 1
	for index >= 0 {
		propertiesMap := mainContainer.PropertiesArray[index]
		for key, value := range propertiesMap {
			if strings.HasPrefix(key, prefix) {
				rkey := strings.Trim(key, prefix)
				returnMap[rkey] = value
			}
		}
		index--
	}
	return returnMap
}

func (mainContainer *MainContainer) LoadFileProperty(propertyFile string) {
	propertyMap := make(map[string]string)
	file, err := os.Open(propertyFile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		propertyLine := scanner.Text()
		properties := strings.Split(propertyLine, "=")
		if len(properties) == 2 {
			key := strings.TrimSpace(properties[0])
			value := strings.TrimSpace(properties[1])
			propertyMap[key] = value
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
	mainContainer.PropertiesArray = append(mainContainer.PropertiesArray, propertyMap)
}

func (mainContainer *MainContainer) LoadSystemProperty() {
	propertyMap := make(map[string]string)
	env := os.Environ()
	for _, propertyLine := range env {
		properties := strings.Split(propertyLine, "=")
		key := strings.TrimSpace(properties[0])
		value := strings.TrimSpace(properties[1])
		propertyMap[key] = value
	}
	mainContainer.PropertiesArray = append(mainContainer.PropertiesArray, propertyMap)
}

func (mainContainer *MainContainer) LoadConverter(converter converter.Converter) {
	mainContainer.TypeConverterMap[converter.Support()] = converter
}

func (mainContainer *MainContainer) CanConvert(typ reflect.Type) bool {
	_, ok := mainContainer.TypeConverterMap[typ]
	return ok
}

func (mainContainer *MainContainer) Register(object interface{}) {
	pv1 := reflect.ValueOf(object)
	if pv1.Kind() != reflect.Ptr {
		objectType := reflect.TypeOf(object)
		panic("register must be a pointer:" + objectType.String())
	}
	objectType := reflect.TypeOf(object)
	mainContainer.ObjectContainer[objectType] = object
}

func (mainContainer *MainContainer) InjectBasicType(objectType reflect.Type, objectValue reflect.Value, index int) {
	fieldType := objectType.Field(index)
	property := fieldType.Tag.Get("property")
	if property != "" {
		propertyStringValue := mainContainer.GetProperty(property)
		propertyConverter, ok := mainContainer.TypeConverterMap[fieldType.Type]
		if ok {
			propertyValue, err := propertyConverter.Convert(propertyStringValue)
			if err != nil {
				panic(err)
			}
			objectValue.Field(index).Set(reflect.ValueOf(propertyValue))
		} else {
			panic("can not find property converter:" + fieldType.Type.String())
		}
	}
}

func (mainContainer *MainContainer) InjectSlice(objectType reflect.Type, objectValue reflect.Value, index int) {
	sliceType := objectType.Field(index)
	property := sliceType.Tag.Get("property")
	if property != "" {
		propertyStringValue := mainContainer.GetProperty(property)
		propertyConverter, ok := mainContainer.TypeConverterMap[sliceType.Type.Elem()]
		if ok {
			values := strings.Split(propertyStringValue, ",")
			for _, v := range values {
				propertyValue, err := propertyConverter.Convert(v)
				if err != nil {
					panic(err)
				}
				objectValue.Field(index).Set(reflect.Append(objectValue.Field(index), reflect.ValueOf(propertyValue)))
			}
		} else {
			panic("can not find property converter:" + sliceType.Type.Elem().String())
		}
	}
}

func (mainContainer *MainContainer) InjectMap(objectType reflect.Type, objectValue reflect.Value, index int) {
	fieldType := objectType.Field(index)
	fieldValue := objectValue.Field(index)

	property := fieldType.Tag.Get("property")
	if property != "" {
		//获取的配置中的值
		propertiesMap := mainContainer.GetPropertyMap(property)
		if len(propertiesMap) == 0 {
			panic("properties map not config:" + fieldType.Type.String())
		}
		//要注入的map的value的类型
		t := objectValue.Field(index).Type()
		fmt.Println(t)
		fmt.Println(t.Elem())
		propertyConverter, ok := mainContainer.TypeConverterMap[t.Elem()]
		if ok {
			injectMap := reflect.MakeMap(t)
			for configKey, configValue := range propertiesMap {
				propertyValue, err := propertyConverter.Convert(configValue)
				if err != nil {
					panic(err)
				}
				injectMap.SetMapIndex(reflect.ValueOf(configKey), reflect.ValueOf(propertyValue))
			}
			fieldValue.Set(reflect.ValueOf(injectMap.Interface()))
		} else {
			panic("can not find property converter:" + fieldType.Type.String())
		}
	}
}

func (mainContainer *MainContainer) InjectInterface(objectType reflect.Type, objectValue reflect.Value, index int) {

	fieldType := objectType.Field(index)
	fieldValue := objectValue.Field(index)
	require := fieldType.Tag.Get("require")
	if require == "" {
		return
	}
	injectCount := 0
	for _, registerObject := range mainContainer.ObjectContainer {
		if utils.Implement(registerObject, fieldType.Type) {
			injectCount++
			fieldValue.Set(reflect.ValueOf(registerObject))
		}
	}
	if injectCount >= 2 {
		panic(fieldType.Type.String() + " has more than two instances")
	}
	if require == "true" && injectCount == 0 {
		panic(fieldType.Type.String() + " must has one instances")
	}
}

func (mainContainer *MainContainer) InjectObject(objectType reflect.Type, objectValue reflect.Value, index int) {
	fieldType := objectType.Field(index)
	fieldValue := objectValue.Field(index)
	//防止循环依赖
	//if fieldValue.IsNil(){
	require := fieldType.Tag.Get("require")
	if require == "" {
		return
	}
	o := mainContainer.ObjectContainer[fieldValue.Type()]
	if require == "true" && o == nil {
		panic("inject object can not find " + fieldValue.Type().String())
	}
	if require != "true" && o == nil {
		return
	}
	fieldValue.Set(reflect.ValueOf(o))
	//}
}

func (mainContainer *MainContainer) Inject() {
	for _, object := range mainContainer.ObjectContainer {
		objectType := reflect.TypeOf(object).Elem()
		objectValue := reflect.ValueOf(object).Elem()
		for i := 0; i < objectType.NumField(); i++ {
			fieldType := objectType.Field(i)
			if mainContainer.CanConvert(fieldType.Type) {
				mainContainer.InjectBasicType(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Slice {
				mainContainer.InjectSlice(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Map {
				mainContainer.InjectMap(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Interface {
				mainContainer.InjectInterface(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Ptr {
				mainContainer.InjectObject(objectType, objectValue, i)
			} else {
				panic("inject object must be pointer")
			}
		}
	}
}

func (mainContainer *MainContainer) InitProperty() {
	mainContainer.LoadSystemProperty()
	mainContainer.LoadFileProperty("application.property")
}

func (mainContainer *MainContainer) InitConverter() {
	mainContainer.LoadConverter(&converter.BoolConvert{})
	mainContainer.LoadConverter(&converter.BoolPtrConvert{})
	mainContainer.LoadConverter(&converter.Float32Convert{})
	mainContainer.LoadConverter(&converter.Float32PtrConvert{})
	mainContainer.LoadConverter(&converter.Float64Convert{})
	mainContainer.LoadConverter(&converter.Float64PtrConvert{})
	mainContainer.LoadConverter(&converter.Int8Converter{})
	mainContainer.LoadConverter(&converter.Int8PtrConverter{})
	mainContainer.LoadConverter(&converter.Int16Converter{})
	mainContainer.LoadConverter(&converter.Int16PtrConverter{})
	mainContainer.LoadConverter(&converter.Int32Converter{})
	mainContainer.LoadConverter(&converter.Int32PtrConverter{})
	mainContainer.LoadConverter(&converter.Int64Converter{})
	mainContainer.LoadConverter(&converter.Int64PtrConverter{})
	mainContainer.LoadConverter(&converter.IntConverter{})
	mainContainer.LoadConverter(&converter.IntPtrConverter{})
	mainContainer.LoadConverter(&converter.StringConvert{})
	mainContainer.LoadConverter(&converter.StringPtrConvert{})
	mainContainer.LoadConverter(&converter.Uint8Converter{})
	mainContainer.LoadConverter(&converter.Uint8PtrConverter{})
	mainContainer.LoadConverter(&converter.Uint16Converter{})
	mainContainer.LoadConverter(&converter.Uint16PtrConverter{})
	mainContainer.LoadConverter(&converter.Uint32Converter{})
	mainContainer.LoadConverter(&converter.Uint32PtrConverter{})
	mainContainer.LoadConverter(&converter.Uint64Converter{})
	mainContainer.LoadConverter(&converter.Uint64PtrConverter{})
	mainContainer.LoadConverter(&converter.UintConverter{})
	mainContainer.LoadConverter(&converter.UintPtrConverter{})
}

func (mainContainer *MainContainer) InitContainer() {
	mainContainer.ObjectContainer = make(map[reflect.Type]interface{})
	mainContainer.PropertiesArray = make([]map[string]string, 0)
	mainContainer.TypeConverterMap = make(map[reflect.Type]converter.Converter)
	mainContainer.BeforeInitPropertyArray = make([]lifecycle.BeforeInitProperty, 0)
	mainContainer.AfterInitPropertyArray = make([]lifecycle.AfterInitProperty, 0)
	mainContainer.AfterInitConverterArray = make([]lifecycle.AfterInitConverter, 0)
	mainContainer.AfterInitInjectArray = make([]lifecycle.AfterInitInject, 0)
	mainContainer.AfterRunArray = make([]lifecycle.AfterRun, 0)
}

func (mainContainer *MainContainer) RegisterBeforeInitProperty(beforeInitProperty lifecycle.BeforeInitProperty) {
	mainContainer.BeforeInitPropertyArray = append(mainContainer.BeforeInitPropertyArray, beforeInitProperty)
}

func (mainContainer *MainContainer) RegisterAfterInitProperty(afterInitProperty lifecycle.AfterInitProperty) {
	mainContainer.AfterInitPropertyArray = append(mainContainer.AfterInitPropertyArray, afterInitProperty)
}

func (mainContainer *MainContainer) RegisterAfterInitConverter(afterInitConverter lifecycle.AfterInitConverter) {
	mainContainer.AfterInitConverterArray = append(mainContainer.AfterInitConverterArray, afterInitConverter)
}
func (mainContainer *MainContainer) RegisterAfterInitInject(afterInitInject lifecycle.AfterInitInject) {
	mainContainer.AfterInitInjectArray = append(mainContainer.AfterInitInjectArray, afterInitInject)
}
func (mainContainer *MainContainer) RegisterAfterRun(afterRun lifecycle.AfterRun) {
	mainContainer.AfterRunArray = append(mainContainer.AfterRunArray, afterRun)
}

func (mainContainer *MainContainer) InitLifeCycle() {
	for _, registerObject := range mainContainer.ObjectContainer {
		BeforeInitPropertyObject, ok := registerObject.(lifecycle.BeforeInitProperty)
		if ok {
			mainContainer.BeforeInitPropertyArray = append(mainContainer.BeforeInitPropertyArray, BeforeInitPropertyObject)
		}
		AfterInitPropertyObject, ok := registerObject.(lifecycle.AfterInitProperty)
		if ok {
			mainContainer.AfterInitPropertyArray = append(mainContainer.AfterInitPropertyArray, AfterInitPropertyObject)
		}
		AfterInitConverterObject, ok := registerObject.(lifecycle.AfterInitConverter)
		if ok {
			mainContainer.AfterInitConverterArray = append(mainContainer.AfterInitConverterArray, AfterInitConverterObject)
		}
		AfterInitInjectObject, ok := registerObject.(lifecycle.AfterInitInject)
		if ok {
			mainContainer.AfterInitInjectArray = append(mainContainer.AfterInitInjectArray, AfterInitInjectObject)
		}
		AfterRunObject, ok := registerObject.(lifecycle.AfterRun)
		if ok {
			mainContainer.AfterRunArray = append(mainContainer.AfterRunArray, AfterRunObject)
		}
	}
}

func (mainContainer *MainContainer) Start() {

	mainContainer.InitLifeCycle()

	for _, value := range mainContainer.BeforeInitPropertyArray {
		if err := value.BeforeInitPropertyAction(); err != nil {
			panic(err)
		}
	}
	mainContainer.InitProperty()
	for _, value := range mainContainer.AfterInitPropertyArray {
		if err := value.AfterInitPropertyAction(mainContainer.PropertiesArray); err != nil {
			panic(err)
		}
	}
	mainContainer.InitConverter()
	for _, value := range mainContainer.AfterInitConverterArray {
		if err := value.AfterInitConverterAction(mainContainer.TypeConverterMap); err != nil {
			panic(err)
		}
	}
	mainContainer.Inject()
	for _, value := range mainContainer.AfterInitInjectArray {
		if err := value.AfterInitInjectAction(mainContainer.ObjectContainer); err != nil {
			panic(err)
		}
	}
	for _, value := range mainContainer.AfterRunArray {
		if err := value.AfterRunAction(mainContainer.ObjectContainer); err != nil {
			panic(err)
		}
	}
}