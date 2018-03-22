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


