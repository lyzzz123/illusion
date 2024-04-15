# illusion
## illusion是什么
illusion是用go语言开发的一个开源容器，illusion提供struct的对象托管和对象生命周期管理，同时支持控制反转（IoC）来实现解耦。
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
//单个对象的属性全部注入完成之后
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