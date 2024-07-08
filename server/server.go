package main

import (
	"zinx/ziface"
	"zinx/znet"

	"github.com/sirupsen/logrus"
)

type PingRouter struct {
	znet.BaseRouter
}

func (prouter *PingRouter) PreHandler(request *ziface.IRequest) {
	logrus.Println("Call Router PreHandler")
	var req ziface.IRequest = (*request)
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	data := req.GetData()
	logrus.Println(string(data))
	if err != nil {
		logrus.Errorln(err)
	}

}
func (prouter *PingRouter) Handler(request *ziface.IRequest) {
	logrus.Println("Call Router Handler")
	var req ziface.IRequest = (*request)
	_, err := req.GetConnection().GetTCPConnection().Write([]byte(" ping\n"))
	if err != nil {
		logrus.Errorln(err)
	}

}
func (prouter *PingRouter) PostHandler(request *ziface.IRequest) {
	logrus.Println("Call Router PostHandler")
	var req ziface.IRequest = (*request)
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		logrus.Errorln(err)
	}
}

func main() {
	s := znet.NewServer("ubuntu server")
	s.AddRoute(&PingRouter{})
	s.Serve()
}
