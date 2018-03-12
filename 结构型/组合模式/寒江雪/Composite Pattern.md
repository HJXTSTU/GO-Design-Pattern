## Composite Pattern

&emsp;&emsp;组合模式是以树形结构来组织组件完成某个功能。当根节点的某个方法被调用时，会遍历树结构，调用相应的方法。不过具体如何设计是因人而异的。<br>

&emsp;&emsp; 它突出的是整体和部分的关系。<br>

## 实现

&emsp;&emsp;组合模式的结构中有两种类型，一种是分支节点，它包含了子节点列表，可以往里边添加和删除节点。另一种是叶子节点，叶子节点没有子节点。这两种节点都实现了同一个接口:Component。这个接口只包含了一个方法。（根据实际情况调整)<br>

&emsp;&emsp;分支节点还要实现自己的Add,GetChild,Remove方法。<br>

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

&emsp;&emsp;我这里就实现简单的使用方式吧。

```go
package main

import (
	"projects/DesignPatternsByGo/structuralPatterns/composite"
	"fmt"
)

func main(){
	root := composite.NewComponent(func() {
		fmt.Println("My name is:"+"root")
	},true).(*composite.Composite)
	root.Add(composite.NewComponent(func() {
		fmt.Println("I'm Leaf.")
	},false).(composite.Component))
	composite_2 := composite.NewComponent(func() {
		fmt.Println("I'm Composite2.")
	},true).(*composite.Composite)

	composite_2.Add(composite.NewComponent(func() {
		fmt.Println("I'm Leaf2.")
	},false).(composite.Component))
	composite_2.Add(composite.NewComponent(func() {
		fmt.Println("I'm Leaf3.")
	},false).(composite.Component))

	c2 := root.Add(composite_2)

	root.Operation()
	root.Remove(c2)
	root.Operation()
}
```

<br>