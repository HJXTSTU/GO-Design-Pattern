package SimpleFactory

import . "projects/DesignPatternsByGo/CreationalPatterns/Factories"

/***
 *	Simple Factory
 */
type FoodFactory struct {
}

func (ff FoodFactory) CreateFood(name string) Food {
	var s Food;
	switch name {
	case "Meat":
		s = new(Meat);
	case "Hamberger":
		s = new(Hamberger)
	}
	return s;
}
