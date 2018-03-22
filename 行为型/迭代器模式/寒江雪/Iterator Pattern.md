## Iterator Pattern

&emsp;&emsp;迭代器模式可以把对象的访问方式给封装出来，只需要给某种数据结构实现自己的迭代器，用户只需要拿到迭代器就可以轻易操作该数据结构，而不需要在乎底层实现。

## 实现

封装一个容器

```go
package iterator

import "container/list"

type Container interface{
	Iterator()Iterator
}

type List struct{
	list list.List
}

func (this *List)Iterator()Iterator{
	return &ListIterator{this.list.Front(),this.list.Back()}
}

func (this *List)Add(value interface{}){
	this.list.PushBack(value)
}

```

为容器实现迭代器

```go
package iterator

import (
	"container/list"
)

type Iterator interface{
	HasNext()bool
	Value()interface{}
	Next()
}

type ListIterator struct {
	cur *list.Element
	end *list.Element
}

func (this *ListIterator)HasNext()bool{
	return this.cur != this.end
}

func (this *ListIterator)Next(){
	this.cur=this.cur.Next()
}

func (this *ListIterator)Value()interface{}{
	return this.cur.Value
}
```

main.go

```go
package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/iterator"
	"fmt"
)

func main() {
	l := iterator.List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	i := l.Iterator()
	for i.HasNext(){
		x := i.Value().(int)
		fmt.Println(x)
		i.Next()
	}
}
```

<h2>有钱的捧个钱场，没钱的捧个人场。</h2>
<h2>出来混不容易。</h2>
<img src="https://raw.githubusercontent.com/lkysyzxz/pictureForMD/master/money.jpg" height="400" width="400">