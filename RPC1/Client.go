package main

import (
	"net/rpc"
	"fmt"
)

func main(){

	type AddPrama struct {
		Arg1 float32
		Arg2 float32
	}




	client,err:=rpc.DialHTTP("tcp","localhost:8081")

	if err!=nil{
		panic(err.Error())
	}

	var req float32

	req =3

	var resp *float32

	err=client.Call("gogogo.CalculateCircleArea",req,&resp)

	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(*resp)


	add:=new(AddPrama)

	add.Arg2=100
	add.Arg1=233

	err= client.Call("gogogo.Add",add,&resp)
	if err!=nil{
		panic(err.Error())
	}

	fmt.Println(*resp)

}
