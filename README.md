# illusion
## illusion是什么
illusion是用go语言开发的一个开源容器，illusion提供struct的对象托管和对象生命周期管理，同时支持控制反转（IoC）来实现解耦。配合illusionmvc可以实现一个web服务。
## 快速开始
首先定义一个struct，之后将struct创建出来的对象放到illusion容器中，这就实现了一个对象的托管，示例代码如下
```
type TestStruct struct {
	Hello string
}
func main() {
	illusion.Register(&TestStruct{})
	illusion.Start()
}
```
这段代码执行完会马上退出，因为这只是一个容器，对象实现托管之后就马上退出了。
## illusion 生命周期
illusion的生命周期分成两种，一种是illusion容器的生命周期，一种是struct对象的生命周期。
### 容器生命周期
1. BeforeContainerInitProperty，容器加载属性之前。
2. AfterContainerInitProperty，容器加载属性之后。
3. AfterContainerInitConverter，容器加载类型转换器之后。
4. AfterContainerInject，容器中的全部托管对象属性注入完之后。
5. AfterRun，容器启动完之后，可以在这个扩展点运行其他的框架，比如illusionmvc，这样就可以启动一个web服务了。
以上五个全部都是接口，自定义的struct只要实现了上述接口，并把struct的对象注册进illusion，那么生命周期函数就会得到执行
```
type ContainerLifeCycleTest struct {

}
//在整个程序启动前执行，例如可以在这个扩展点输出一些banner
func (containerLifeCycleTest *ContainerLifeCycleTest) BeforeContainerInitPropertyAction() error {
	fmt.Println("BeforeContainerInitPropertyAction")
	return nil
}
//属性加载完成之后，可以在这个扩展点对加载完的属性做一些操作
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInitPropertyAction(propertiesArray []map[string]string) error{
	fmt.Println("AfterContainerInitPropertyAction")
	return nil
}
//数据类型转换器加载完成之后，可以在这里加载一些自定义的类型转换器
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInitConverterAction(typeConverterMap map[reflect.Type]converter.Converter) error{
	fmt.Println("AfterContainerInitConverterAction")
	return nil
}
//容器中的所有对象的属性都注入完成之后，可以在这里对托管对象做一些自定义操作
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterContainerInjectAction(objectContainer map[reflect.Type]interface{}) error{
	fmt.Println("AfterContainerInjectAction")
	return nil
}
//容器启动的最后一个扩展点，可以在这里集成一些其他的程序框架，比如illusionmvc
func (containerLifeCycleTest *ContainerLifeCycleTest) AfterRunAction(objectContainer map[reflect.Type]interface{}) error{
	fmt.Println("AfterRunAction")
	return nil
}
func (containerLifeCycleTest *ContainerLifeCycleTest) GetPriority() int{
	return 1
}

func main() {
	illusion.Register(&ContainerLifeCycleTest{})
	illusion.Start()
}
```
### struct对象的生命周期
1. AfterObjectInject，struct对象的所有属性都注入完之后
2. AfterObjectDestroy，struct对象被摧毁之后，就是illusion容器退出的时候
```
type ObjectLifeCycleTest struct {

}
//单个对象的属性全部注入完成之后, 可以在这里执行像消息队列连接之类等初始化操作
func (objectLifeCycleTest *ObjectLifeCycleTest) AfterObjectInjectAction() error {
	fmt.Println("AfterObjectInjectAction")
	return nil
}

func (objectLifeCycleTest *ObjectLifeCycleTest) AfterObjectDestroyAction() error {
	fmt.Println("AfterObjectDestroyAction")
	return nil
}

func main() {
	illusion.Register(&ObjectLifeCycleTest{})
	illusion.Start()
}
```
struct对象的生命周期只对单个对象起作用

## illusion 属性加载
illusion属性加载有三个途径
1. 命令行参数
2. 系统环境变量
3. 属性文件

### 命令行参数加载属性
例如编译完的程序叫illusion，那可以使用illusion #aaa=bbb 命令来启动程序，这样启动完，key为aaa，value为bbb的属性就加载到程序里了
### 系统环境变量加载属性
执行程序的机器的环境变量会加载到程序里
### 属性文件加载属性
illusion强制使用application.property为属性文件，文件内容为key=value的形式。
### 分环境加载属性文件
如果以上三种形式的属性中存在environment.active=xxx的属性，那么程序会继续加载名为application-xxx.property的属性文件。例如environment.active的属性值为dev，那么程序就会继续加载名为application-dev.property的属性文件
### 属性的优先级
命令行参数加载属性 > 系统环境变量加载属性 > application.property属性 > application-xxx.property属性


## 依赖注入
illusion依赖注入分为属性注入和对象注入
### 属性注入
属性注入是指对托管对象注入illusion加载的属性
目前支持注入
1. 所有go的基础类型和基础类型的指针
2. 所有基础类型和基础类型指针的slice
3. 所有key为string，value为所有基础类型和基础类型指针的map
例如
```
application.property文件中的属性

bool.bool=true
bool.ptr=false
float32.float32=3.2
float32.ptr=-3.2
float64.float64=6.4
float64.ptr=-6.4
int8.int8=8
int8.ptr=-8
int16.int16=16
int16.ptr=-16
int32.int32=32
int32.ptr=-32
int64.int64=64
int64.ptr=-64
int.int=128
int.ptr=-128
string.string=wwww
string.ptr=-qqqqqq
uint8.uint8=18
uint8.ptr=18
uint16.uint16=116
uint16.ptr=116
uint32.uint32=132
uint32.ptr=132
uint64.uint64=164
uint64.ptr=164
uint.uint=11212
uint.ptr=11212
int.slice=1,2,3,4,5
int.map.k1=1
int.map.k2=2



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
func main() {
	propertyInjectTest := &PropertyInjectTest{}
	illusion.Register(propertyInjectTest)
	illusion.Start()
}
```
1. 在想要注入的字段后面添加名为property的tag来表示要注入的属性名字，默认为注入不能为空，如果go程序中没有要注入的属性，会报错。
如果想注入字段可以为空，可以在tag中的属性名后面加上", false"，如示例中的Float32float32字段。
2. slice的注入，需要go中的属性值为逗号分隔的字符串，注入slice的时候，会自动把字符串转成slice
3. map的注入，需要go中有多个key有共同前缀的属性，共同前缀就作为map的注入名称，不同的后缀作为map的key，这些key对应的value作为map的value，例如示例中IntMap字段
### 对象注入
除了属性注入，illusion还支持对象注入，对象注入支持对象指针注入和接口注入
#### 对象指针注入
```
type TestA struct {

}

type TestB struct {
	MTestA *TestA `require:"true"`
}

func main() {
	illusion.Register(&TestA{})
	illusion.Register(&TestB{})
	illusion.Start()
}

```
示例中TestA已经被注入到TestB。被注入的字段必须添加名为require的tag，没有require，属性不会被注入。
当require的值为true时，注入不能为空，不然会保存，当require的值为false时，注入可以为空。
illusion注册对象必须是指针，不然会报错。
#### 接口注入
```
type TestInjectInterface interface {
    PrintMessage()
}

type TestA struct {

}

func (testA *TestA) PrintMessage() {
	fmt.Println("this is test a")
}

type TestB struct {
	MTestA TestInjectInterface `require:"true"`
}

func main() {
	illusion.Register(&TestA{})
	illusion.Register(&TestB{})
	illusion.Start()
}

```
TestA实现了接口TestInjectInterface，那么TestA会被以接口的方式注入到TestB




## 静态代理
illusion支持对托管对象的静态代理
illusion的静态代理是基于接口代理的，代理对象和目标对象必须实现同一个接口，同时代理对象还要实现代理接口。例如

```
//代理接口
type Proxy interface {
    SupportInterface() reflect.Type
    SetTarget(target interface{})
}
//--------------------------------------------------------
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
//要代理的接口方法
func (testTargetProxy *TestTargetProxy) PrintMessage() {
	targetProxy, _ := testTargetProxy.Target.(TestTargetInterface)
	fmt.Println("before target run")
	targetProxy.PrintMessage()
	fmt.Println("after target run")
}
//代理接口的实现
func (testTargetProxy *TestTargetProxy) SupportInterface() reflect.Type {
	return reflect.TypeOf(new(TestTargetInterface)).Elem()
}
func (testTargetProxy *TestTargetProxy) SetTarget(target interface{}) {
	testTargetProxy.Target = target
}

type InjectObject struct {
	Target TestTargetInterface `require:"true"`
}

func main() {
	injectObject := &InjectObject{}
	illusion.Register(&TestTarget{})
	illusion.Register(&TestTargetProxy{})
	illusion.Register(injectObject)
	illusion.Start()
}
```
如果TestTargetProxy没有注册，那么injectObject中被注入的是TestTarget实例。注册了TestTargetProxy，那么injectObject中被注入的就是TestTargetProxy对象。

```



```


