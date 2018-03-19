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
