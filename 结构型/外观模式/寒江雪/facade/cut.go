package facade

func CutVegtable(veg ...*Vegetable)[]*Vegetable{
	for _,v := range veg{
		v.status=STATUS_CUTED
	}
	return veg
}
