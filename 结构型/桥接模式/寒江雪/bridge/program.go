package bridge

import (
	"fmt"
)

type IProgram interface{
	GetName()string
	SetSuffix(suffix string)
	CMD()string
	Run()
}

type AbstractProgram struct{
	Name string
	Cmd string
}

func (this *AbstractProgram)GetName()string{
	return this.Name
}

func (this *AbstractProgram)CMD()string{
	return this.Cmd
}

func (this *AbstractProgram)SetSuffix(suffix string){
	this.Name+=suffix
}

func (this *AbstractProgram)Run(){

}

type Sing struct{
	AbstractProgram
	Word string
}

func (this *Sing)Run(){
	fmt.Println("I want to sing:"+this.Word)
}


type Dog struct{
	AbstractProgram
}

func (this *Dog)Run(){
	fmt.Println("Wang!")
}


