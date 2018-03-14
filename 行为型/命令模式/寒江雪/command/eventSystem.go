package command

/*
	#include "goKey_c.h"
*/
import "C"

const (
	EVENT_CODE_KEY = iota
)

type EventCode rune
type EventData interface{}
type Event struct {
	Code EventCode
	Data EventData
}

type Command interface {
	Execute(Event)
}

type EventFunc func(Event)

type EventFunder struct {
	f EventFunc
}

func (this EventFunder) Execute(data Event) {
	this.f(data)
}

func eventFunder(eventFunc EventFunc) EventFunder {
	return EventFunder{eventFunc}
}

type Hash map[EventCode]Command

type EventSystem struct {
	hash Hash
}

func (this *EventSystem) Init() *EventSystem {
	this.hash = make(Hash)
	return this
}

func (this *EventSystem) Map(code EventCode, eventFunc EventFunc) {
	this.hash[code] = eventFunder(eventFunc)
}

func (this *EventSystem) InspectKeyboard() {
	go func() {
		C.init_keyboard()
		for {
			if C.kbhit() > 0 {
				ch := EventCode(C.readch())
				// TODO::generate key event
				this.generateEvent(Event{EVENT_CODE_KEY, ch})
			}
		}
	}()
}

func (this *EventSystem) generateEvent(e Event) {
	v, ok := this.hash[e.Code]
	if ok {
		v.Execute(e)
	}
}

func NewEventSystem() *EventSystem {
	return (&EventSystem{}).Init()
}
