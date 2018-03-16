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


