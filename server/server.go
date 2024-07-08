package main

import "zinx/znet"

func main() {
	s := znet.NewServer("ubuntu server")
	s.Serve()
}
