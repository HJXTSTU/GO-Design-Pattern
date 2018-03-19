package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/state"
	"time"
)

func stateMechine(state state.State, ch chan int) {
	for {
		select {
		case i := <-ch:
			if i == 1 {
				state = state.NextState()
			} else if i == 0 {
				return
			}
		default:
			state.Update()
		}
	}
}

func main() {
	st := new(state.GameStartState)
	ch := make(chan int)
	go stateMechine(st, ch)
	time.Sleep(time.Second * 3)
	ch <- 1
	time.Sleep(time.Second * 3)
	ch <- 1
	time.Sleep(time.Second * 3)
	ch <- 0
	time.Sleep(time.Second * 3)
}
