package mediator

import "gopkg.in/mgo.v2/bson"

type Consumer struct{
	Id bson.ObjectId
	Name string
}

func NewConsumer(name string)*Consumer{
	return &Consumer{bson.NewObjectId(),name}
}
