package facade

const (
	STATUS_BUYED  = iota
	STATUS_STORED
	STATUS_CUTED
	STATUS_COOKED
	STATUS_EATED
)

type VegStatus int

type Vegetable struct {
	name   string
	status VegStatus
}

func BuyVegetable(name string) *Vegetable {
	return &Vegetable{name, STATUS_BUYED}
}

func Eat(veg *Vegetable) {
	veg.status = STATUS_EATED
}
