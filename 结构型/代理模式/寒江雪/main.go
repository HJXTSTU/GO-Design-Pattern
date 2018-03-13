package main

import "projects/DesignPatternsByGo/structuralPatterns/proxy"

func main(){
	proxy := new(proxy.ProObject)
	proxy.ObjDo("Well")
}
