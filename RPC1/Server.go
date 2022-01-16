package main

import (
	"math"
	"net/rpc"
	"net"
	"net/http"
)

type MathUtil struct {


}

func (mu *MathUtil) CalculateCircleArea(req float32,resp *float32) error{

	*resp=math.Pi*req*req

	return nil
}


func main(){


	  // 1 初始化类型
	  mathUtil:= new(MathUtil)


	  // 注册
	  err:=rpc.Register(mathUtil)


	  if err!=nil{
	  	panic(err.Error())
	  }

	  //3 通过该函数将mathutil 中的服务注册到http协议上,方便调用者可以利用http的方式进行调用
      rpc.HandleHTTP()

	  //4 在特定的端口进行监听

	  listen,err := net.Listen("tcp",":8081")

	  if err!=nil{
	  	panic(err.Error())
	  }

	   http.Serve(listen,nil)


}

