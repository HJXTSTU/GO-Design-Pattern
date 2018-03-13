package proxy

import (
	"fmt"
	"sync"
)

type IObject interface {
	ObjDo(action string)
}

type Object struct {
	action string
}

func (this *Object) ObjDo(action string) {
	fmt.Println("I can do:" + action)
}

type ProObject struct {
	obj *Object
}

var one  = new(sync.Once)
func (this *ProObject) ObjDo(action string) {
	one.Do(func() {
		if this.obj==nil{
			this.obj=new(Object)
		}
	})
	this.obj.ObjDo(action)
}
