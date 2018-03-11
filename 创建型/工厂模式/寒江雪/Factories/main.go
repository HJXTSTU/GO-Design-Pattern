package main

func main(){
	// Abstract Factory
	fa := FactoryA{}
	fa.CreateFood().Eat()
	fa.CreateDrink().Drink()

	fb := FactoryB{}
	fb.CreateFood().Eat()
	fb.CreateDrink().Drink()
}
