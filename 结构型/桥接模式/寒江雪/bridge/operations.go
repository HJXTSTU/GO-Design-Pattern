package bridge

import (
	"fmt"
	"strings"
)

type Registry map[string]IProgram

type Operation struct {
	reg         Registry
	validSuffix string
}

type IOperation interface {
	Boot()
	ExecuteProgram(cmd string)
	SetupProgram(cmd string, program IProgram)
	Build(program *IProgram)
}

func (this *Operation)Boot(){
	this.reg = make(Registry)
}

func (this *Operation) ExecuteProgram(cmd string) {
	program, ok := this.reg[cmd]
	if ok {
		program.Run()
	} else {
		fmt.Println(cmd + " is invalid.")
	}
}

func (this *Operation) SetupProgram(cmd string, program IProgram) {
	switch this.validSuffix {
	case "*":
		this.reg[cmd] = program
	default:
		name := program.GetName()
		if strings.HasSuffix(name, this.validSuffix) {
			this.reg[cmd] = program
		} else {
			fmt.Println("Invalid suffix.")
		}
	}
}

func (this *Operation) Build(program IProgram){
	cmd := program.CMD()
	if this.validSuffix!="*"{
		program.SetSuffix(this.validSuffix)
	}
	this.SetupProgram(cmd,program)
}

type Windows struct {
	Operation
}

func (this *Windows) Boot() {
	this.Operation.Boot()
	this.validSuffix = ".exe"
}

type Linux struct {
	Operation
}

func (this *Linux) Boot() {
	this.Operation.Boot()
	this.validSuffix = "*"
}
