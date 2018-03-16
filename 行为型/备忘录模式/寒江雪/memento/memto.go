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
