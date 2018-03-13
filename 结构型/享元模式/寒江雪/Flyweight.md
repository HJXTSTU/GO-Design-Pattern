## Flyweight Pattern

&emsp;&emsp;享元模式使用共享物件，来尽可能多地减少内存使用量以及共享给尽可能多的相似组件。<br>

&emsp;&emsp;它适合用于当大量物件只是重复因而导致无法令人接受的使用大量内存。通常物件中的部分状态是可以分享。常见做法是把它们放在外部数据结构，当需要使用时再将它们传递给享元。<br>

## 实现

&emsp;&emsp;我们举个例子:<br>

* 游戏中的角色都有基础属性
* 游戏中的角色都会穿上装备
* 装备是不同的单基础属性都是一样的<br>

&emsp;&emsp;结合装饰者模式来开发一个角色穿上不同装备后，HP上限和MP上限增加的例子.<br>

```go
package flyweight

type IProperty interface{
	GetHPLimit()int
	GetMPLimit()int
}

type PeopleBase struct{
	MAX_HP int
	MAX_MP int
}

func NewPeopleBase()*PeopleBase{
	return &PeopleBase{100,100}
}

func (this *PeopleBase)GetHPLimit()int{
	return this.MAX_HP
}

func (this *PeopleBase)GetMPLimit()int{
	return this.MAX_MP
}

type Helmet struct{
	base IProperty
	HP_ADD int
	MP_ADD int
}

func (this *Helmet)GetHPLimit()int{
	return this.base.GetHPLimit()+this.HP_ADD
}

func (this *Helmet)GetMPLimit()int{
	return this.base.GetMPLimit()+this.MP_ADD
}

func NewHelmet(property IProperty,hp_add,mp_add int)*Helmet{
	return &Helmet{property,hp_add,mp_add}
}
```

<br>

&emsp;&emsp;还需要一个享元工厂来存放享元<br>

```go
package flyweight

type Element struct {
	Value interface{}
}

func newElement(value interface{})*Element{
	return &Element{value}
}

type FlyweightFactory struct {
	pool map[string]*Element
}

func (this *FlyweightFactory) GetElement(key string) interface{} {
	return this.pool[key].Value
}



func (this *FlyweightFactory)SetElement(key string,value interface{}){
	ne := newElement(value)
	this.pool[key]=ne
}

func NewFlyweight()*FlyweightFactory{
	flyweight := FlyweightFactory{}
	flyweight.pool=make(map[string]*Element)
	return &flyweight
}
```

<br>

&emsp;&emsp;最后模拟一个例子<br>

```go
package main

import (
	"projects/DesignPatternsByGo/structuralPatterns/flyweight"
	"fmt"
)

func main(){
    //	创建享元
	fly := flyweight.NewFlyweight()
	base := flyweight.NewPeopleBase()
	fly.SetElement("PeopleBase",base)

    //	生成两个人物
	people_1 := fly.GetElement("PeopleBase").(flyweight.IProperty)
	people_2 := fly.GetElement("PeopleBase").(flyweight.IProperty)
    
    //	捡到装备了——装饰者
	people_1 = flyweight.NewHelmet(people_1,10,10)
	people_2 = flyweight.NewHelmet(people_2,100,100)

    //	获取它们血量上限和魔法上限
	hp_1 := people_1.GetHPLimit()
	mp_1 := people_1.GetMPLimit()

	hp_2 := people_2.GetHPLimit()
	mp_2 := people_2.GetMPLimit()

	fmt.Printf("People_1:\nHP:%d\nMP:%d\n",hp_1,mp_1)
	fmt.Printf("People_2:\nHP:%d\nMP:%d\n",hp_2,mp_2)
}
```

<br>

