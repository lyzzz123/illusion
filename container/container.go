package container

import (
	"bufio"
	"github.com/lyzzz123/illusion/converter"
	"github.com/lyzzz123/illusion/lifecycle"
	"github.com/lyzzz123/illusion/proxy"
	"github.com/lyzzz123/illusion/utils"
	"os"
	"os/signal"
	"reflect"
	"sort"
	"strings"
	"syscall"
)

type MainContainer struct {
	ObjectContainer map[reflect.Type]interface{}

	ProxyMap map[reflect.Type]proxy.Proxy

	PropertiesArray []map[string]string

	TypeConverterMap map[reflect.Type]converter.Converter

	BeforeContainerInitPropertyArray []lifecycle.BeforeContainerInitProperty

	AfterContainerInitPropertyArray []lifecycle.AfterContainerInitProperty

	AfterContainerInitConverterArray []lifecycle.AfterContainerInitConverter

	AfterContainerInjectArray []lifecycle.AfterContainerInject

	AfterRunArray []lifecycle.AfterRun

	AfterContainerDestroyArray []lifecycle.AfterContainerDestroy

	AfterObjectInjectArray []lifecycle.AfterObjectInject

	AfterObjectDestroyArray []lifecycle.AfterObjectDestroy
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

func (mainContainer *MainContainer) LoadCmdProperty() {
	propertyMap := make(map[string]string)
	cmdProperties := os.Args[1:]
	for i := 0; i < len(cmdProperties); i++ {
		cmdProperty := cmdProperties[i]
		if strings.HasPrefix(cmdProperty, "#") {
			cmdPropertyKV := strings.Split(strings.TrimLeft(cmdProperty, "#"), "=")
			if len(cmdPropertyKV) == 2 {
				propertyMap[strings.TrimSpace(cmdPropertyKV[0])] = strings.TrimSpace(cmdPropertyKV[1])
			}
		}

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
	if proxyObject, ok := object.(proxy.Proxy); ok {
		mainContainer.ProxyMap[proxyObject.SupportInterface()] = proxyObject
	} else {
		pv1 := reflect.ValueOf(object)
		if pv1.Kind() != reflect.Ptr {
			objectType := reflect.TypeOf(object)
			panic("register must be a pointer, register object type:" + objectType.String())
		}
		objectType := reflect.TypeOf(object)
		mainContainer.ObjectContainer[objectType] = object
	}
}

func (mainContainer *MainContainer) InjectProperty(objectType reflect.Type, objectValue reflect.Value, index int) {
	fieldType := objectType.Field(index)
	property := fieldType.Tag.Get("property")
	//require := fieldType.Tag.Get("require")
	if property != "" {
		properties := strings.Split(property, ",")
		propertyStringValue := mainContainer.GetProperty(strings.TrimSpace(properties[0]))
		if propertyStringValue != "" {
			if propertyConverter, ok := mainContainer.TypeConverterMap[fieldType.Type]; ok {
				if propertyValue, err := propertyConverter.Convert(propertyStringValue); err != nil {
					panic(err)
				} else {
					objectValue.Field(index).Set(reflect.ValueOf(propertyValue))
				}
			} else {
				panic("not find a type converter for property:" + properties[0])
			}
		} else {
			if len(properties) > 1 && strings.TrimSpace(properties[1]) == "true" {
				panic("can not find value for property:" + properties[0])
			}
		}
	}
}

func (mainContainer *MainContainer) InjectSlice(objectType reflect.Type, objectValue reflect.Value, index int) {
	sliceType := objectType.Field(index)
	property := sliceType.Tag.Get("property")
	if property != "" {
		properties := strings.Split(property, ",")
		propertyStringValue := mainContainer.GetProperty(strings.TrimSpace(properties[0]))
		if propertyStringValue != "" {
			if propertyConverter, ok := mainContainer.TypeConverterMap[sliceType.Type.Elem()]; ok {
				values := strings.Split(propertyStringValue, ",")
				for _, v := range values {
					propertyValue, err := propertyConverter.Convert(v)
					if err != nil {
						panic(err)
					}
					objectValue.Field(index).Set(reflect.Append(objectValue.Field(index), reflect.ValueOf(propertyValue)))
				}
			} else {
				panic("not find a type converter for property:" + properties[0])
			}
		} else {
			if len(properties) > 1 && strings.TrimSpace(properties[1]) == "true" {
				panic("can not find value for property:" + properties[0])
			}
		}
	}
}

func (mainContainer *MainContainer) InjectMap(objectType reflect.Type, objectValue reflect.Value, index int) {
	fieldType := objectType.Field(index)
	fieldValue := objectValue.Field(index)
	property := fieldType.Tag.Get("property")
	if property != "" {
		properties := strings.Split(property, ",")
		propertiesMap := mainContainer.GetPropertyMap(strings.TrimSpace(properties[0]))
		if len(propertiesMap) != 0 {
			//要注入的map的value的类型
			t := objectValue.Field(index).Type()
			if propertyConverter, ok := mainContainer.TypeConverterMap[t.Elem()]; ok {
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
				panic("not find a type converter for property:" + properties[0])
			}
		} else {
			if len(properties) > 1 && strings.TrimSpace(properties[1]) == "true" {
				panic("can not find value for property:" + properties[0])
			}
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
	var findInterface interface{} = nil
	for _, registerObject := range mainContainer.ObjectContainer {
		if utils.Implement(registerObject, fieldType.Type) {
			injectCount++
			findInterface = registerObject
		}
	}
	if injectCount >= 2 {
		panic(fieldType.Type.String() + " has more than two instances")
	} else if injectCount == 1 {
		if proxy, ok := mainContainer.ProxyMap[fieldType.Type]; ok {
			proxy.SetTarget(findInterface)
			fieldValue.Set(reflect.ValueOf(proxy))
		} else {
			fieldValue.Set(reflect.ValueOf(findInterface))
		}
	} else if injectCount == 0 && require == "true" {
		panic(fieldType.Type.String() + " must has one instance")
	}

}

func (mainContainer *MainContainer) InjectObject(objectType reflect.Type, objectValue reflect.Value, index int) {
	fieldType := objectType.Field(index)
	fieldValue := objectValue.Field(index)

	require := fieldType.Tag.Get("require")
	if require == "" {
		return
	}
	o := mainContainer.ObjectContainer[fieldValue.Type()]
	if o != nil {
		fieldValue.Set(reflect.ValueOf(o))
	} else {
		if require == "true" {
			panic("inject object can not find " + fieldValue.Type().String())
		}
	}
}

func (mainContainer *MainContainer) Inject() {
	for _, object := range mainContainer.ObjectContainer {
		objectType := reflect.TypeOf(object).Elem()
		objectValue := reflect.ValueOf(object).Elem()
		for i := 0; i < objectType.NumField(); i++ {
			fieldType := objectType.Field(i)
			if mainContainer.CanConvert(fieldType.Type) {
				mainContainer.InjectProperty(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Slice {
				mainContainer.InjectSlice(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Map {
				mainContainer.InjectMap(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Interface {
				mainContainer.InjectInterface(objectType, objectValue, i)
			} else if fieldType.Type.Kind() == reflect.Ptr {
				mainContainer.InjectObject(objectType, objectValue, i)
			} else {
				panic("inject object must be pointer:" + objectType.String())
			}
		}
	}
}

func (mainContainer *MainContainer) InitProperty() {
	mainContainer.LoadCmdProperty()
	mainContainer.LoadSystemProperty()
	mainContainer.LoadFileProperty("application.property")
	active := mainContainer.GetProperty("environment.active")
	if active != "" {
		activeNamePath := "application-" + active + ".property"
		mainContainer.LoadFileProperty(activeNamePath)
	}
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
	mainContainer.ProxyMap = make(map[reflect.Type]proxy.Proxy)
	mainContainer.PropertiesArray = make([]map[string]string, 0)
	mainContainer.TypeConverterMap = make(map[reflect.Type]converter.Converter)
	mainContainer.BeforeContainerInitPropertyArray = make([]lifecycle.BeforeContainerInitProperty, 0)
	mainContainer.AfterContainerInitPropertyArray = make([]lifecycle.AfterContainerInitProperty, 0)
	mainContainer.AfterContainerInitConverterArray = make([]lifecycle.AfterContainerInitConverter, 0)
	mainContainer.AfterContainerInjectArray = make([]lifecycle.AfterContainerInject, 0)
	mainContainer.AfterRunArray = make([]lifecycle.AfterRun, 0)
	mainContainer.AfterContainerDestroyArray = make([]lifecycle.AfterContainerDestroy, 0)
	mainContainer.AfterObjectInjectArray = make([]lifecycle.AfterObjectInject, 0)
	mainContainer.AfterObjectDestroyArray = make([]lifecycle.AfterObjectDestroy, 0)

}

func (mainContainer *MainContainer) RegisterBeforeInitProperty(beforeInitProperty lifecycle.BeforeContainerInitProperty) {
	mainContainer.BeforeContainerInitPropertyArray = append(mainContainer.BeforeContainerInitPropertyArray, beforeInitProperty)
}

func (mainContainer *MainContainer) RegisterAfterInitProperty(afterInitProperty lifecycle.AfterContainerInitProperty) {
	mainContainer.AfterContainerInitPropertyArray = append(mainContainer.AfterContainerInitPropertyArray, afterInitProperty)
}

func (mainContainer *MainContainer) RegisterAfterInitConverter(afterInitConverter lifecycle.AfterContainerInitConverter) {
	mainContainer.AfterContainerInitConverterArray = append(mainContainer.AfterContainerInitConverterArray, afterInitConverter)
}

func (mainContainer *MainContainer) RegisterAfterInitInject(afterInitInject lifecycle.AfterContainerInject) {
	mainContainer.AfterContainerInjectArray = append(mainContainer.AfterContainerInjectArray, afterInitInject)
}

func (mainContainer *MainContainer) RegisterAfterRun(afterRun lifecycle.AfterRun) {
	mainContainer.AfterRunArray = append(mainContainer.AfterRunArray, afterRun)
}

func (mainContainer *MainContainer) InitLifeCycle() {
	for _, registerObject := range mainContainer.ObjectContainer {
		BeforeInitPropertyObject, ok := registerObject.(lifecycle.BeforeContainerInitProperty)
		if ok {
			mainContainer.BeforeContainerInitPropertyArray = append(mainContainer.BeforeContainerInitPropertyArray, BeforeInitPropertyObject)
		}
		AfterInitPropertyObject, ok := registerObject.(lifecycle.AfterContainerInitProperty)
		if ok {
			mainContainer.AfterContainerInitPropertyArray = append(mainContainer.AfterContainerInitPropertyArray, AfterInitPropertyObject)
		}
		AfterInitConverterObject, ok := registerObject.(lifecycle.AfterContainerInitConverter)
		if ok {
			mainContainer.AfterContainerInitConverterArray = append(mainContainer.AfterContainerInitConverterArray, AfterInitConverterObject)
		}
		AfterInitInjectObject, ok := registerObject.(lifecycle.AfterContainerInject)
		if ok {
			mainContainer.AfterContainerInjectArray = append(mainContainer.AfterContainerInjectArray, AfterInitInjectObject)
		}
		AfterRunObject, ok := registerObject.(lifecycle.AfterRun)
		if ok {
			mainContainer.AfterRunArray = append(mainContainer.AfterRunArray, AfterRunObject)
		}
		AfterContainerDestroyObject, ok := registerObject.(lifecycle.AfterContainerDestroy)
		if ok {
			mainContainer.AfterContainerDestroyArray = append(mainContainer.AfterContainerDestroyArray, AfterContainerDestroyObject)
		}
		afterObjectInjectObject, ok := registerObject.(lifecycle.AfterObjectInject)
		if ok {
			mainContainer.AfterObjectInjectArray = append(mainContainer.AfterObjectInjectArray, afterObjectInjectObject)
		}
		afterObjectDestroyObject, ok := registerObject.(lifecycle.AfterObjectDestroy)
		if ok {
			mainContainer.AfterObjectDestroyArray = append(mainContainer.AfterObjectDestroyArray, afterObjectDestroyObject)
		}
	}
}

func (mainContainer *MainContainer) TestStart() {
	mainContainer.InitLifeCycle()
	mainContainer.executeBeforeContainerInitPropertyArray()
	mainContainer.InitProperty()
	mainContainer.executeAfterContainerInitPropertyArray()
	mainContainer.InitConverter()
	mainContainer.executeAfterContainerInitConverterArray()
	mainContainer.Inject()
	mainContainer.executeAfterObjectInjectArray()
	mainContainer.executeAfterContainerInjectArray()
	mainContainer.executeAfterRunArray()
}

func (mainContainer *MainContainer) Start() {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	closeChannel := make(chan int, 1)

	go func() {
		defer func() {
			closeChannel <- 1
		}()
		mainContainer.InitLifeCycle()
		mainContainer.executeBeforeContainerInitPropertyArray()
		mainContainer.InitProperty()
		mainContainer.executeAfterContainerInitPropertyArray()
		mainContainer.InitConverter()
		mainContainer.executeAfterContainerInitConverterArray()
		mainContainer.Inject()
		mainContainer.executeAfterObjectInjectArray()
		mainContainer.executeAfterContainerInjectArray()
		mainContainer.executeAfterRunArray()
	}()
	select {
	case <-quitChannel:
	case <-closeChannel:
	}
	mainContainer.executeAfterObjectDestroyArray()
	mainContainer.executeAfterContainerDestroyArray()
}

func (mainContainer *MainContainer) executeBeforeContainerInitPropertyArray() {
	sort.SliceStable(mainContainer.BeforeContainerInitPropertyArray, func(i, j int) bool {
		return mainContainer.BeforeContainerInitPropertyArray[i].GetPriority() > mainContainer.BeforeContainerInitPropertyArray[j].GetPriority()
	})
	for _, value := range mainContainer.BeforeContainerInitPropertyArray {
		if err := value.BeforeContainerInitPropertyAction(); err != nil {
			panic(err)
		}
	}
}

func (mainContainer *MainContainer) executeAfterContainerInitPropertyArray() {
	sort.SliceStable(mainContainer.AfterContainerInitPropertyArray, func(i, j int) bool {
		return mainContainer.AfterContainerInitPropertyArray[i].GetPriority() > mainContainer.AfterContainerInitPropertyArray[j].GetPriority()
	})
	for _, value := range mainContainer.AfterContainerInitPropertyArray {
		if err := value.AfterContainerInitPropertyAction(mainContainer.PropertiesArray); err != nil {
			panic(err)
		}
	}
}

func (mainContainer *MainContainer) executeAfterContainerInitConverterArray() {
	sort.SliceStable(mainContainer.AfterContainerInitConverterArray, func(i, j int) bool {
		return mainContainer.AfterContainerInitConverterArray[i].GetPriority() > mainContainer.AfterContainerInitConverterArray[j].GetPriority()
	})
	for _, value := range mainContainer.AfterContainerInitConverterArray {
		if err := value.AfterContainerInitConverterAction(mainContainer.TypeConverterMap); err != nil {
			panic(err)
		}
	}
}

func (mainContainer *MainContainer) executeAfterContainerInjectArray() {
	sort.SliceStable(mainContainer.AfterContainerInjectArray, func(i, j int) bool {
		return mainContainer.AfterContainerInjectArray[i].GetPriority() > mainContainer.AfterContainerInjectArray[j].GetPriority()
	})
	for _, value := range mainContainer.AfterContainerInjectArray {
		if err := value.AfterContainerInjectAction(mainContainer.ObjectContainer); err != nil {
			panic(err)
		}
	}
}

func (mainContainer *MainContainer) executeAfterRunArray() {
	sort.SliceStable(mainContainer.AfterRunArray, func(i, j int) bool {
		return mainContainer.AfterRunArray[i].GetPriority() > mainContainer.AfterRunArray[j].GetPriority()
	})
	for _, value := range mainContainer.AfterRunArray {
		if err := value.AfterRunAction(mainContainer.ObjectContainer); err != nil {
			panic(err)
		}
	}
}

func (mainContainer *MainContainer) executeAfterContainerDestroyArray() {
	sort.SliceStable(mainContainer.AfterContainerDestroyArray, func(i, j int) bool {
		return mainContainer.AfterContainerDestroyArray[i].GetPriority() > mainContainer.AfterContainerDestroyArray[j].GetPriority()
	})
	for _, value := range mainContainer.AfterContainerDestroyArray {
		if err := value.AfterContainerDestroyAction(mainContainer.ObjectContainer); err != nil {
			panic(err)
		}
	}
}

func (mainContainer *MainContainer) executeAfterObjectInjectArray() {
	for _, value := range mainContainer.AfterObjectInjectArray {
		if err := value.AfterObjectInjectAction(); err != nil {
			panic(err)
		}
	}
}
func (mainContainer *MainContainer) executeAfterObjectDestroyArray() {
	for _, value := range mainContainer.AfterObjectDestroyArray {
		if err := value.AfterObjectDestroyAction(); err != nil {
			panic(err)
		}
	}
}
