package facade

var storage []*Vegetable = make([]*Vegetable, 0)

func SaveVegetables(veg ...*Vegetable) {
	storage = append(storage,veg...)
	for _,v := range veg{
		v.status=STATUS_STORED
	}
}

func GetVegetables() *Vegetable {
	l := len(storage)
	res := storage[l-1]
	storage = storage[:l-1]
	return res
}
