package mediator

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Mediator interface{

}

type RoomHash map[bson.ObjectId]*Room
type ConsumerHash map[bson.ObjectId]*Consumer


type RoomMediator struct {
	roomHash RoomHash
	consumerHash ConsumerHash
}

func NewRoomMediator()*RoomMediator{
	res := RoomMediator{}
	res.roomHash=make(RoomHash)
	res.consumerHash=make(ConsumerHash)
	return &res
}

func (this *RoomMediator)SetRoom(room *Room){
	if len(this.consumerHash)>0{
		var c *Consumer
		for _,v := range this.consumerHash{
			fmt.Println(v.Name +" buy " + room.Name)
			c = v
			break
		}
		delete(this.consumerHash,c.Id)
		return
	}
	this.roomHash[room.Id]=room
}

func (this *RoomMediator)SetConsumerHash(consumer *Consumer){
	if len(this.roomHash)>0{
		var r *Room
		for _,v := range this.roomHash{
			fmt.Println(consumer.Name +" buy "+ v.Name)
			r = v
			break
		}
		delete(this.roomHash,r.Id)
		return
	}
	this.consumerHash[consumer.Id]=consumer
}


