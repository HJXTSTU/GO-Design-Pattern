package builder

import (
	"fmt"
	"errors"
)

type Color string
type LampStatus bool
type Brand string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

const (
	OppleBulb Brand = "OPPLE"
	Osram           = "OSRAM"
)

//	Lamp Builder define
type Builder interface {
	Color(Color) LampBuilder
	Brand(Brand) LampBuilder
	Build() LampOperation
}

type LampBuilder struct {
	Lamp
}

func (lb LampBuilder) Color(c Color) LampBuilder {
	lb.color = c
	return lb
}

func (lb LampBuilder) Brand(b Brand) LampBuilder {
	lb.brand = b
	return lb
}

func (lb LampBuilder) Build() LampOperation {
	lamp := Lamp{color: lb.color, brand: lb.brand, status: false}
	return lamp
}

func NewBuilder() Builder {
	return LampBuilder{}
}

type LampOperation interface {
	Open() error
	Close() error
	ProductionIllustrative()
}

type Lamp struct {
	color  Color
	brand  Brand
	status LampStatus
}

func (l Lamp) Open() error {
	if l.status {
		return errors.New("Lamp is opened")
	}
	fmt.Println("Open lamp.")
	l.status = true
	return nil
}

func (l Lamp) Close() error {
	if !l.status {
		return errors.New("Lamp is closed")
	}
	fmt.Println("Close lamp.")
	l.status = true
	return nil;
}

func (l Lamp) ProductionIllustrative() {
	fmt.Println("I'm a lamp.")
	fmt.Println("Color:" + l.color)
	fmt.Println("Brand:" + l.brand)
}
