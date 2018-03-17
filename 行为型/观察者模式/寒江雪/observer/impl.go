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
