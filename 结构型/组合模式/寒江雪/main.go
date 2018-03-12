package main

import (
	"projects/DesignPatternsByGo/structuralPatterns/composite"
	"fmt"
)

func main(){
	root := composite.NewComponent(func() {
		fmt.Println("My name is:"+"root")
	},true).(*composite.Composite)
	root.Add(composite.NewComponent(func() {
		fmt.Println("I'm Leaf.")
	},false).(composite.Component))
	composite_2 := composite.NewComponent(func() {
		fmt.Println("I'm Composite2.")
	},true).(*composite.Composite)

	composite_2.Add(composite.NewComponent(func() {
		fmt.Println("I'm Leaf2.")
	},false).(composite.Component))
	composite_2.Add(composite.NewComponent(func() {
		fmt.Println("I'm Leaf3.")
	},false).(composite.Component))

	c2 := root.Add(composite_2)

	root.Operation()
	root.Remove(c2)
	root.Operation()
}
