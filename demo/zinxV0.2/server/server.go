package main

import "github.com/SakBlues/zinx/znet"

func main() {
	s := znet.NewServer("[Zinx V0.2]")

	s.Serve()
}
