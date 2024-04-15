# illusion
## illusion是什么
illusion是用go语言开发的一个开源容器，illusion提供struct的对象托管和对象生命周期管理，同时支持控制反转（IoC）来实现解耦。
## 快速开始
首先定义一个struct，之后将struct创建出来的对象放到illusion容器中，这就实现了一个对象的托管，示例代码如下
type TestStruct struct {
	Hello string
}
func main() {
	illusion.Register(&TestStruct{})
	illusion.Start()
}
这段代码执行完会马上退出，因为这只是一个容器，对象实现托管之后就马上退出了。




