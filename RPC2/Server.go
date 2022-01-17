package main

import (
	"net/rpc"
	"net"
)

type HelloService struct {


}

func (p *HelloService) Hello(request string,reply *string) error{

	*reply ="hello "+ request

	return nil
}

func main(){


	rpc.RegisterName("HelloService",new(HelloService))

	listener,err := net.Listen("tcp",":1234")

	if err!=nil{
		panic(err.Error())
	}

	conn,err :=listener.Accept()

	if err!=nil{
		panic(err.Error())
	}

	rpc.ServeConn(conn)
}
