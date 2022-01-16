package main

import (
	"net/rpc"
	"fmt"
)

func main(){

	client,err:=rpc.DialHTTP("tcp","localhost:8081")

	if err!=nil{
		panic(err.Error())
	}

	var req float32

	req =3

	var resp *float32

	err=client.Call("MathUtil.CalculateCircleArea",req,&resp)

	if err!=nil{
		panic(err.Error())
	}

	fmt.Print(*resp)

}
