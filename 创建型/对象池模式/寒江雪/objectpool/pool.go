package pool

import (
	"fmt"
	"strconv"
)

type Object struct{

}

func (Object)Do(index int){
	fmt.Println("Object Do:"+strconv.Itoa(index))
}


type Pool chan *Object

func NewPool(total int)*Pool{
	p := make(Pool,total)
	for i := 0;i<total;i++{
		p <- new(Object)
	}
	return &p
}