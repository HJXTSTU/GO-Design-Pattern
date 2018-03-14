package main
import (
	"fmt"
	"time"
	"projects/DesignPatternsByGo/behavioralPatterns/command"
)

func main(){
	sys := command.NewEventSystem()

	sys.Map(command.EVENT_CODE_KEY, func(data command.Event) {
		fmt.Println(data)
	})

	sys.InspectKeyboard()
	time.Sleep(time.Second*100)
}