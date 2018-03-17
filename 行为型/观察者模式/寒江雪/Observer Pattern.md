## Observer Pattern

&emsp;&emsp;观察者模式使得一种类型的实例可以发送事件给其他类型，前提是接收事件的实例要根发送者订阅这个事件。

## 实现

&emsp;&emsp;先来定义要使用到的接口

```go
package observer

type(
	Event struct{
		Data int64
	}

	Observer interface{
		OnNotify(Event)
	}

	Notifier interface{
		Register(Observer)

		Degister(Observer)

		Notify(Event)
	}
)
```

&emsp;&emsp;然后写几个简单的类型实现这些接口

```go
package observer

import "fmt"

type EventCenter struct{
	observers []Observer
}

func (this *EventCenter)Notify(event Event){
	for _,v:=range this.observers{
		v.OnNotify(event)
	}
}


func (this *EventCenter)Register(o Observer){
	this.observers=append(this.observers,o)
}

func (this *EventCenter)Degister(o Observer){
	for i:=0;i<len(this.observers);i++{
		if this.observers[i]==o{
			this.observers=append(this.observers[:i],this.observers[i+1:]...)
			break
		}
	}
}

func NewEventCenter()*EventCenter{
	res := EventCenter{}
	res.observers=make([]Observer,0)
	return &res
}


type EventReciver struct{

}

func (this *EventReciver)OnNotify(event Event){
	fmt.Printf("Event receive:%d\n",event.Data)
}
```

&emsp;&emsp;最后是main文件

```go
package main

func main(){
	eventCenter := observer.NewEventCenter()
	r_1 := observer.EventReciver{}
	r_2 := observer.EventReciver{}

	eventCenter.Register(&r_1)
	eventCenter.Register(&r_2)
	eventCenter.Notify(observer.Event{1})
	eventCenter.Degister(&r_1)
	eventCenter.Notify(observer.Event{2})
}
```



