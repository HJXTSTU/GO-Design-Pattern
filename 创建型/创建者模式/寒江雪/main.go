package main

import "projects/DesignPatternsByGo/CreationalPatterns/builder"

func main(){
	b := builder.NewBuilder()
	lamp_1 := b.Color(builder.BlueColor).Brand(builder.Osram).Build()
	lamp_1.Open()
	lamp_1.ProductionIllustrative()

	lamp_2 := b.Color(builder.GreenColor).Brand(builder.OppleBulb).Build()
	lamp_2.ProductionIllustrative()
}
