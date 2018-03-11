## Bridge Pattern

&emsp;&emsp;Bridge Pattern说的是，当一个系统中，包含一个组件，该组件是可变的，该系统是可变的。这个时候就需要一个桥连接抽象的系统和抽象的组件。<br>

&emsp;&emsp;如果该系统包含多个其他的组件，这些组件都是可变的时候，也需要在该系统和这些组件之间架桥。<br>

&emsp;&emsp;如果该系统包含的某个组件所包含的组件，是可变的，那么递归地重复上述过程。<br>

&emsp;&emsp;也就是说，在桥接模式中存在两个抽象，这两个抽象之间存在组合关系，为了防止类的数量膨胀，将它们分离出来分别实现，但由于存在包含关系，在上级中包含对下级抽象的引用。<br>

&emsp;&emsp;看了很多资料都说的是将抽象与实现进行解耦，通俗地说就是，大系统包含小组件，大系统的某个方法需要依赖小组件的方法，不类型的小组件的实现不一样，但是都能满足大系统的需求，大系统无需关心小组件怎么实现。大系统又分为很多种类，每个种类的大系统都会需要不同的小组件来实现自己的功能，根据需求的变化而变化。

&emsp;&emsp;为了应对这种变化，就需要一个抽象类，表示大系统，其中包含了小组件的抽象类或接口，当我们需要开发新的组件的时候，就直接继承小组件，并实现抽象方法就行了。当我们需要开发新的大系统的时候，就继承大系统，实现相应的抽象方法就行了，桥接模式的优点在这里就体现出来了，我们这时候就可以为不同的大系统与小组件之间做自由组合了。<br>

## UML

![](https://raw.githubusercontent.com/lkysyzxz/pictureForMD/csdn_blog_golang/Bridge%20Pattern%20uml.png)

## 例子

&emsp;&emsp;举个例子，这个例子优点复杂，其中还用了Builder Pattern来生成AbstractProgram。<br>

&emsp;&emsp;我们设想一下，操作系统和应用程序。操作系统执行应用程序并不需要知道应用程序具体是如何实现的。但是操作系统是由一个个应用程序组成的。<br>

&emsp;&emsp;因此，我们先来设计操作系统的接口，并实现它，同时写两个具体的系统:Windows和Linux。<br>

```go
package bridge

import (
	"fmt"
	"strings"
)

// 命令行和应用程序的映射表
type Registry map[string]IProgram

// 操作系统的通用接口
type IOperation interface {
	Boot()
	ExecuteProgram(cmd string)
	SetupProgram(cmd string, program IProgram)
	Build(program *IProgram)
}

// 抽象的操作系统
type Operation struct {
	reg         Registry
	validSuffix string
}

// 启动
func (this *Operation)Boot(){
	this.reg = make(Registry)
}

// 执行
func (this *Operation) ExecuteProgram(cmd string) {
	program, ok := this.reg[cmd]
	if ok {
		program.Run()
	} else {
		fmt.Println(cmd + " is invalid.")
	}
}

// 安装
func (this *Operation) SetupProgram(cmd string, program IProgram) {
	switch this.validSuffix {
	case "*":
		this.reg[cmd] = program
	default:
		name := program.GetName()
		if strings.HasSuffix(name, this.validSuffix) {
			this.reg[cmd] = program
		} else {
			fmt.Println("Invalid suffix.")
		}
	}
}

// 构建
func (this *Operation) Build(program IProgram){
	cmd := program.CMD()
	if this.validSuffix!="*"{
		program.SetSuffix(this.validSuffix)
	}
	this.SetupProgram(cmd,program)
}

// Windows
type Windows struct {
	Operation
}

// Windows 的启动过程
func (this *Windows) Boot() {
	this.Operation.Boot()
	this.validSuffix = ".exe"
}

// Linux
type Linux struct {
	Operation
}

//	Linux的启动过程
func (this *Linux) Boot() {
	this.Operation.Boot()
	this.validSuffix = "*"
}
```

<br>

&emsp;&emsp;然后我们来设计和实现应用程序<br>

```go
package bridge

import (
	"fmt"
)

//	程序通用接口
type IProgram interface{
	GetName()string
	SetSuffix(suffix string)
	CMD()string
	Run()
}

//	程序的抽象
type AbstractProgram struct{
	Name string
	Cmd string
}

//	获取程序名字
func (this *AbstractProgram)GetName()string{
	return this.Name
}

//	获取程序命令行
func (this *AbstractProgram)CMD()string{
	return this.Cmd
}

//	设置程序后缀名
func (this *AbstractProgram)SetSuffix(suffix string){
	this.Name+=suffix
}

//	执行程序
//	AbstractProgram并不具体实现
//	为了让AbstractProgram看起来实现了IProgram接口
//	如果不写这段就需要子结构强制实现Run接口
func (this *AbstractProgram)Run(){

}

type Sing struct{
	AbstractProgram
	Word string
}

func (this *Sing)Run(){
	fmt.Println("I want to sing:"+this.Word)
}


type Dog struct{
	AbstractProgram
}

func (this *Dog)Run(){
	fmt.Println("Wang!")
}
```

<br>

&emsp;&emsp;最后，多余的一步，希望更逼近现实情况，我们需要构建应用程序，就需要一些工具了，这里使用了Builder Pattern来实现这个工具。<br>

```go
package bridge

type ICoding interface {
	WriteName(string) ICoding
	WriteCmd(string) ICoding
	Compile() AbstractProgram
}

type CodeUtil struct {
	code AbstractProgram
}

func (this CodeUtil) WriteName(name string) ICoding {
	this.code.Name = name
	return this
}

func (this CodeUtil) WriteCmd(cmd string) ICoding {
	this.code.Cmd = cmd
	return this
}

func (this CodeUtil) Compile() AbstractProgram {
	res := AbstractProgram{Name: this.code.Name, Cmd: this.code.Cmd}
	return res
}

func (this CodeUtil) GetSingProgram(word string) Sing {
	res := Sing{}
	res.AbstractProgram = this.WriteName("sing").WriteCmd("sing").Compile()
	res.Word = word
	return res
}

func (this CodeUtil) GetDogProgram() Dog {
	res := Dog{}
	res.AbstractProgram = this.WriteName("dog").WriteCmd("dog").Compile()
	return res
}
```

<br>

### 使用

```go
func main(){
	// craete windows
	windows := new(bridge.Windows)
	windows.Boot()

	// create linux
	linux := new(bridge.Linux)
	linux.Boot()

	// Code util
	ide := bridge.CodeUtil{}

	// write programs
	sing := ide.GetSingProgram("Hello World")
	dog := ide.GetDogProgram()

	windows.Build(&sing)
	windows.Build(&dog)
	linux.Build(&sing)
	linux.Build(&dog)

	windows.ExecuteProgram("dog")
	windows.ExecuteProgram("sing")
	linux.ExecuteProgram("sing")
	linux.ExecuteProgram("dog")
}
```

<br>

