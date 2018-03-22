package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/template"
)

func main() {
	a := template.TmplA{}
	b := template.TmplB{}
	template.Operate(&a)
	template.Operate(&b)

}
