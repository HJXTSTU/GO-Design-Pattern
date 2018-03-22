## Template Pattern

&emsp;&emsp;模版方法设计模式允许把对象不同的部分抽象，在同一段代码中执行相同的逻辑，增加可拓展性。在Go语言中，实现由底层对象实现，而行为由顶层方法控制。



## 实现

```go
package template

import "fmt"

type Tmpl interface{
	DoSomething()
	DoAnything()
}

type TmplA struct{

}

func (this *TmplA)DoSomething(){
	fmt.Println("TmplA.DoSomething")
}

func (this *TmplA)DoAnything(){
	fmt.Println("TmplA.DoAnything")
}

type TmplB struct{

}

func (this *TmplB)DoSomething(){
	fmt.Println("TmplB.DoSomething")
}

func (this *TmplB)DoAnything(){
	fmt.Println("TmplB.DoAnything")
}

func Operate(tmpl Tmpl){
	tmpl.DoSomething()
	tmpl.DoAnything()
}
```

使用

```go
package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/template"
)

func main() {
	a := template.TmplA{}
	b := template.TmplB{}
	template.Operate(&a)
	template.Operate(&b)

}
```

<h2>有钱的捧个钱场，没钱的捧个人场。</h2>
<h2>出来混不容易。</h2>
<img src="https://raw.githubusercontent.com/lkysyzxz/pictureForMD/master/money.jpg" height="400" width="400">