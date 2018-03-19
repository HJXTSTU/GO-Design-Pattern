## State Pattern

&emsp;&emsp;状态模式把对象每一个状态的行为封装在对象内部。避免大量状态逻辑杂糅。

## 实现

```go
package state

import "fmt"

type State interface{
	NextState()State
	Update()
}

type GameStartState struct{

}

type GameRunState struct{

}

type GameEndState struct{

}

func (this *GameStartState)NextState()State{
	fmt.Println("Start Next")
	return new(GameRunState)
}

func (this *GameStartState) Update(){
	fmt.Println("Game Start")

}

func (this *GameRunState)NextState()State{
	fmt.Println("Run Next")
	return new(GameEndState)
}

func (this *GameRunState) Update(){
	fmt.Println("Run")
}

func (this *GameEndState)NextState()State{
	return new(GameStartState)
}

func (this *GameEndState) Update(){
	fmt.Println("End")
}

```



```go
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
```

