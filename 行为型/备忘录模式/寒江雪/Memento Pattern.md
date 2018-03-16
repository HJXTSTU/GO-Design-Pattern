## Memento Pattern

&emsp;&emsp;备忘录模式保存一个对象的状态，在需要的时候将其恢复。

&emsp;&emsp;该模式在不破坏封装的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，这样可以在以后将对象恢复到原先保存的状态。

&emsp;&emsp;很多时候我们总是需要记录一个对象的内部状态，这样做的目的就是为了允许用户取消不确定或者错误的操作，能够恢复到他原先的状态，使得他有"后悔药"可吃。



## 实现

&emsp;&emsp;我们以游戏中存档作为例子。

&emsp;&emsp;先创建角色。

```go
package memento

type Role struct {
	Hp int
}

func (this *Role) Save() *MemoryObject {
	res := NewMemoryObject()
	res.Save("Hp", this.Hp)
	return res
}

func (this *Role) Read(memory *MemoryObject) {
	this.Hp = memory.Read("Hp").(int)
}

func (this *Role) Fight() {
	this.Hp /= 2;
}

func NewRole(hp int) *Role {
	return &Role{hp}
}



```

&emsp;&emsp;接着创建备忘录对象和保存备忘录的结构。

```go
package memento

type MemoryHash map[string]interface{}
type MemoryObject struct {
	hash MemoryHash
}

func (this *MemoryObject) Init() *MemoryObject {
	this.hash = make(MemoryHash)
	return this
}

func (this *MemoryObject) Save(key string, value interface{}) {
	this.hash[key] = value
}

func (this *MemoryObject) Read(key string) interface{} {
	return this.hash[key]
}

func NewMemoryObject() *MemoryObject {
	return (&MemoryObject{}).Init()
}

type Memoriable interface {
	Save() *MemoryObject
	Read(object *MemoryObject)
}




type CaretakerRoleMemory struct {
	roleMemory []*MemoryObject
}

func (this *CaretakerRoleMemory) Save(memory *MemoryObject) {
	this.roleMemory = append(this.roleMemory, memory)
}

func (this *CaretakerRoleMemory) GetAndRemoveMemory() *MemoryObject {
	l := len(this.roleMemory)
	res := this.roleMemory[l-1]
	this.roleMemory = this.roleMemory[:l-1]
	return res
}

func NewCaretakerRoleMemory() *CaretakerRoleMemory {
	caretakerRoleMemory := CaretakerRoleMemory{}
	caretakerRoleMemory.roleMemory = make([]*MemoryObject, 0)
	return &caretakerRoleMemory
}

```

然后是main文件

```go
package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/memento"
	"fmt"
)

func main(){
	s := memento.NewCaretakerRoleMemory()
	man := memento.NewRole(100)

	// 存档
	s.Save(man.Save())

	// 战斗
	man.Fight()
	fmt.Println(*man)

	// 再存档
	s.Save(man.Save())

	// 在战斗
	man.Fight()

	//	回档
	man.Read(s.GetAndRemoveMemory())
	fmt.Println(*man)

	// 再回档
	man.Read(s.GetAndRemoveMemory())
	fmt.Println(*man)
}
```



## 应用实例

* 后悔药
* 打游戏时的存档
* Windows 里的 ctri + z
* IE 中的后退
* 数据库的事务管理



## 优点

* 给用户提供了一种可以恢复状态的机制，可以使用户能够比较方便地回到某个历史的状态
* 实现了信息的封装，使得用户不需要关心状态的保存细节



## 缺点

* 消耗资源。如果类的成员变量过多，势必会占用比较大的资源，而且每一次保存都会消耗一定的内存。