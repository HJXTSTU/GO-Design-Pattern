package facade


func SauteVegtable()[]*Vegetable{
	qc := BuyVegetable("青菜")
	suan := BuyVegetable("蒜")
	jiang := BuyVegetable("姜")
	SaveVegetables(qc,suan,jiang)

	vegs := CookVegtable(CutVegtable(storage...)...)
	return vegs
}
