package main

import (
	"net/rpc"
	"fmt"
)

func main(){


	client,err :=rpc.Dial("tcp","localhost:1234")

	if err!=nil{
		panic(err.Error())
	}

	var reply string

	err=client.Call("HelloService.Hello","hello",&reply)

	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(reply)
}
