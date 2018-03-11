package Factories

import "fmt"

type Food interface {
	Eat()
}

type Meat struct {
}

type Hamberger struct {
}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}

func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

type Drink interface{
	Drink()
}

type CoCo struct{

}

func (cc CoCo)Drink(){
	fmt.Println("Drink CoCo")
}

type Tea struct{

}

func (t Tea)Drink(){
	fmt.Println("Drink Tea")
}