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


