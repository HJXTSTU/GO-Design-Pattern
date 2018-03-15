package main

import "projects/DesignPatternsByGo/behavioralPatterns/mediator"

func main(){
	room_1 := mediator.NewRoom("凤鸣山水国际")
	room_2 := mediator.NewRoom("1137")
	consumer_1 := mediator.NewConsumer("wwt")
	consumer_2 := mediator.NewConsumer("wpy")

	m := mediator.NewRoomMediator()
	m.SetRoom(room_1)
	m.SetConsumerHash(consumer_1)
	m.SetConsumerHash(consumer_2)
	m.SetRoom(room_2)
}