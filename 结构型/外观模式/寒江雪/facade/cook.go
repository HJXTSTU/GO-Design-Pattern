package facade

func CookVegtable(vec ...*Vegetable)[]*Vegetable{
	for _,v := range vec{
		v.status=STATUS_COOKED
	}
	return vec
}
