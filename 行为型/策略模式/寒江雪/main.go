package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/strategy"
	"fmt"
)

func main() {
	operation := strategy.Operation{strategy.Addition{}}

	res  := operation.Operate(1,1)

	fmt.Println(res)
}
