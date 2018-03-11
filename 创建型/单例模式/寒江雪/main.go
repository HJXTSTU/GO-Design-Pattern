package main

import (
	"projects/DesignPatternsByGo/CreationalPatterns/singleton"
	"fmt"
)

func main(){
	instance_1 := singleton.GetInstance()
	instance_1["this"]="that"

	instance_2 := singleton.GetInstance()
	s := instance_2["this"]
	fmt.Println(s)
}
