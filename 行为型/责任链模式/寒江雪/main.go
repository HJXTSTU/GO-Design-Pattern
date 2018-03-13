package main

import (
	"projects/DesignPatternsByGo/behavioralPatterns/chainofresponsibility"
	"fmt"
)

func Login(request *chainofresponsibility.Request){
	username := request.PostForm["username"][0]
	password := request.PostForm["password"][0]

	fmt.Println(username)
	fmt.Println(password)
}

func main(){
	mux := chainofresponsibility.NewMux()

	mux.Handle("login",Login)


	req := chainofresponsibility.MakeRequest()
	req.Method="POST"
	req.SetValues("username","111")
	req.SetValues("password","222")
	req.Post("login",mux)
}
