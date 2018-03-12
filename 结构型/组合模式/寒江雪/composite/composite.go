package composite

type Component interface {
	Operation()
}



type Child struct {
	Value interface{}
	list  *Children
}

type Children []*Child

func (this *Children) appendValue(e *Child) *Child {
	(*this) = append(*this, e)
	e.list = this
	return e
}

func (this *Children) remove(c *Child) {
	l := len(*this)
	for i := 0; i < l; i++ {
		if (*this)[i] == c {
			(*this) = append((*this)[:i], (*this)[i+i:]...)
			break
		}
	}
}

func (this *Children) get(i int) *Child {
	return (*this)[i]
}

func newChildren() *Children {
	children := make(Children, 0)
	return &children
}

type Composite struct {
	Leaf
	children *Children
}

func (this *Composite) Operation() {
	this.Op()
	for _, v := range *this.children {
		v.Value.(Component).Operation()
	}
}

func (this *Composite) Add(component Component) *Child {
	if this.children == nil {
		this.children = newChildren()
	}
	return this.children.appendValue(&Child{Value: component})
}

func (this *Composite) GetChild(i int) Component {
	return this.children.get(i).Value.(Component)
}

func (this *Composite) Remove(c *Child) Component {
	if c.list == this.children {
		this.children.remove(c)
	}
	return c.Value.(Component)
}

type Leaf struct {
	Op func()
}

func (this *Leaf) Operation() {
	this.Op()
}

func NewComponent(op func(), isComposite bool) interface{} {
	leaf := Leaf{Op: op}
	if isComposite {
		return &Composite{Leaf: leaf}
	}
	return &leaf
}
