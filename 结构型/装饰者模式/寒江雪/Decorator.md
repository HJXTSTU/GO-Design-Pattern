## Decorator Pattern

&emsp;&emsp;装饰着模式可以在需要扩展某个类的时候，动态地修改而不需要在内部添加代码，也可以防止类爆炸。<br>

&emsp;&emsp; 装饰者模式可以提供了灵活地扩展方案.<br>

## 实现

&emsp;&emsp;实现一个日志的自定义功能.<br>

```go
package decorator

import (
	"time"
	"fmt"
)

type LogDecorate interface {
	Info() string
}

type LogBody struct {
	Msg string
}

func (this LogBody) Info() string {
	return this.Msg
}

type LogTimeField struct {
	dec LogDecorate
}

func (this *LogTimeField) Info() string {
	return time.Now().Format("[2006-1-2 15:04:05]") + this.dec.Info()
}

func NewLogTimeField(decorate LogDecorate)*LogTimeField{
	return &LogTimeField{decorate}
}

type LogNameField struct {
	dec  LogDecorate
	name string
}

func (this *LogNameField) Info() string {
	return this.name + ":" + this.dec.Info()
}

func NewLogNameField(name string,decorate LogDecorate)*LogNameField{
	return &LogNameField{decorate,name}
}

func Log(msg string,name string){
	var log LogDecorate
	log  = LogBody{msg}
	log  = NewLogTimeField(log)
	if name!=""{
		log = NewLogNameField(name,log)
	}
	fmt.Println(log.Info())
}

```

<br>

### 使用

```go
func main(){
	decorator.Log("Yeah","WWT")
}
```

### 另一种实现形式

&emsp;&emsp;不过这次不是日志了.<br>

```go
package decorator

type DecoratorFunc func(float64)float64

func DecFunc(dec DecoratorFunc)DecoratorFunc{
	return func(f float64) float64 {
		result := dec(f)
		return result
	}
}
```

<br>

```go
package main

func Double(decoratorFunc decorator.DecoratorFunc)decorator.DecoratorFunc{
	return func(f float64) float64 {
		var result float64 = f
		if decoratorFunc!=nil {
			result = decoratorFunc(f)
		}
		return result*2
	}
}

func Sqrt(decoratorFunc decorator.DecoratorFunc)decorator.DecoratorFunc{
	return func(f float64) float64{
		var result float64 = f
		if decoratorFunc!=nil {
			result = decoratorFunc(f)
		}
		return math.Sqrt(result)
	}
}



func main(){
	f := decorator.DecFunc(Double(Sqrt(nil)))
	fmt.Println(f(16.0))
}
```

<br>

## 注意

* 装饰者模式是通过注入的方式来装饰被装饰对象的。
* 装着模式不会修改被装饰对象的接口。<br>