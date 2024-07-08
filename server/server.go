package main

import (
	"zinx/ziface"
	"zinx/znet"

	"github.com/sirupsen/logrus"
)

type PingRouter struct {
	znet.BaseRouter
}

func (prouter *PingRouter) PreHandler(request ziface.IRequest) {
	logrus.Println("Call Router PreHandler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	data := request.GetData()
	logrus.Println(string(data))
	if err != nil {
		logrus.Errorln(err)
	}

}
func (prouter *PingRouter) Handler(request ziface.IRequest) {
	logrus.Println("Call Router Handler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte(" ping\n"))
	if err != nil {
		logrus.Errorln(err)
	}

}
func (prouter *PingRouter) PostHandler(request ziface.IRequest) {
	logrus.Println("Call Router PostHandler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		logrus.Errorln(err)
	}
}

func main() {
	s := znet.NewServer("ubuntu server")
	s.AddRoute(&PingRouter{})
	s.Serve()
}
