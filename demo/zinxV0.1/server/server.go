package main

import "github.com/SakBlues/zinx/znet"

func main() {
	s := znet.NewServer("[Zinx V0.1]")

	s.Serve()
}
