package bridge

type ICoding interface {
	WriteName(string) ICoding
	WriteCmd(string) ICoding
	Compile() AbstractProgram
}

type CodeUtil struct {
	code AbstractProgram
}

func (this CodeUtil) WriteName(name string) ICoding {
	this.code.Name = name
	return this
}

func (this CodeUtil) WriteCmd(cmd string) ICoding {
	this.code.Cmd = cmd
	return this
}

func (this CodeUtil) Compile() AbstractProgram {
	res := AbstractProgram{Name: this.code.Name, Cmd: this.code.Cmd}
	return res
}

func (this CodeUtil) GetSingProgram(word string) Sing {
	res := Sing{}
	res.AbstractProgram = this.WriteName("sing").WriteCmd("sing").Compile()
	res.Word = word
	return res
}

func (this CodeUtil) GetDogProgram() Dog {
	res := Dog{}
	res.AbstractProgram = this.WriteName("dog").WriteCmd("dog").Compile()
	return res
}
