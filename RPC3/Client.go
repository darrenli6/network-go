package main

import (
	"fmt"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}




func main(){


	client,err :=rpc.Dial("tcp","localhost:1234")

	if err!=nil{
		panic(err.Error())
	}

	var reply string

	err=client.Call(HelloServiceName+".Hello","golang",&reply)

	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(reply)
}