package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/memento"
	"fmt"
)

func main(){
	s := memento.NewCaretakerRoleMemory()
	man := memento.NewRole(100)

	// 存档
	s.Save(man.Save())

	// 战斗
	man.Fight()
	fmt.Println(*man)

	// 再存档
	s.Save(man.Save())

	// 在战斗
	man.Fight()

	//	回档
	man.Read(s.GetAndRemoveMemory())
	fmt.Println(*man)

	// 再回档
	man.Read(s.GetAndRemoveMemory())
	fmt.Println(*man)
}