package main

import (
	"fmt"

	"github.com/SakBlues/zinx/ziface"
	"github.com/SakBlues/zinx/znet"
)

var _ ziface.IRouter = (*PingRouter)(nil)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before ping..."))
	if err != nil {
		fmt.Println("call back before ping err:", err)
	}
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Ping..."))
	if err != nil {
		fmt.Println("call back ping err:", err)
	}
}

func (r *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping..."))
	if err != nil {
		fmt.Println("call back after ping err:", err)
	}
}

func main() {
	s := znet.NewServer("[Zinx V0.4]")

	s.AddRouter(&PingRouter{})

	s.Serve()
}
