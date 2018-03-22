package template

import "fmt"

type Tmpl interface{
	DoSomething()
	DoAnything()
}

type TmplA struct{

}

func (this *TmplA)DoSomething(){
	fmt.Println("TmplA.DoSomething")
}

func (this *TmplA)DoAnything(){
	fmt.Println("TmplA.DoAnything")
}

type TmplB struct{

}

func (this *TmplB)DoSomething(){
	fmt.Println("TmplB.DoSomething")
}

func (this *TmplB)DoAnything(){
	fmt.Println("TmplB.DoAnything")
}

func Operate(tmpl Tmpl){
	tmpl.DoSomething()
	tmpl.DoAnything()
}