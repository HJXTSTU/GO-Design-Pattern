##  Mediator Pattern

&emsp;&emsp;中介者模式定义一个中介对象来封装系列对象之间的交互。各个对象不需要显示地相互引用，从而使其耦合性松散,而且可以独立地改变它们之间的交互。<br>

## 实现

&emsp;&emsp;说起中介，就想到买房。<br>

```go
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
```

&emsp;&emsp;定义消费者。<br>

```go
package mediator

import "gopkg.in/mgo.v2/bson"

type Consumer struct{
	Id bson.ObjectId
	Name string
}

func NewConsumer(name string)*Consumer{
	return &Consumer{bson.NewObjectId(),name}
}
```

&emsp;&emsp;定义房子。<br>

```go
package mediator

import "gopkg.in/mgo.v2/bson"

type Room struct{
	Id bson.ObjectId
	Name string
}

func NewRoom(name string)*Room{
	return &Room{bson.NewObjectId(),name}
}
```

### 使用

```go
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
```

<br>

## 使用场景

* 一组定义良好的对象，现在要进行复杂的通信。
* 定制一个分布在多个类中的行为，而又不想生成太多的子类。

##  优点 

* 降低了系统对象之间的耦合性，使得对象易于独立的被复用。
* 提高系统的灵活性，使得系统易于扩展和维护。

## 缺点

* 中介者模式的缺点是显而易见的，因为这个“中介“承担了较多的责任，所以一旦这个中介对象出现了问题，那么整个系统就会受到重大的影响。<br>