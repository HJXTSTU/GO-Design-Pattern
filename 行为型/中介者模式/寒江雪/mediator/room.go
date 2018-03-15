package mediator

import "gopkg.in/mgo.v2/bson"

type Room struct{
	Id bson.ObjectId
	Name string
}

func NewRoom(name string)*Room{
	return &Room{bson.NewObjectId(),name}
}
