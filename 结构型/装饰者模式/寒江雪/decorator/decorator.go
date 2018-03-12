package decorator

import (
	"time"
	"fmt"
)

type LogDecorate interface {
	Info() string
}

type LogBody struct {
	Msg string
}

func (this LogBody) Info() string {
	return this.Msg
}

type LogTimeField struct {
	dec LogDecorate
}

func (this *LogTimeField) Info() string {
	return time.Now().Format("[2006-1-2 15:04:05]") + this.dec.Info()
}

func NewLogTimeField(decorate LogDecorate)*LogTimeField{
	return &LogTimeField{decorate}
}

type LogNameField struct {
	dec  LogDecorate
	name string
}

func (this *LogNameField) Info() string {
	return this.name + ":" + this.dec.Info()
}

func NewLogNameField(name string,decorate LogDecorate)*LogNameField{
	return &LogNameField{decorate,name}
}

func Log(msg string,name string){
	var log LogDecorate
	log  = LogBody{msg}
	log  = NewLogTimeField(log)
	if name!=""{
		log = NewLogNameField(name,log)
	}
	fmt.Println(log.Info())
}
