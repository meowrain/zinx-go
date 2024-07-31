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
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	// data := request.GetData()
	// logrus.Println(string(data))
	// if err != nil {
	// 	logrus.Errorln(err)
	// }

}
func (prouter *PingRouter) Handler(request ziface.IRequest) {
	// logrus.Println("Call Router Handler")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte(" ping\n"))
	// if err != nil {
	// 	logrus.Errorln(err)
	// }

	logrus.Println("Call Router Handler")

	// 解析接收到的消息
	dp := znet.NewDataPack()
	data := request.GetData()

	for {
		headlen := dp.GetHeadLen()
		headData := data[0:headlen]
		unpackHead, err := dp.UnpackHead(headData)
		if err != nil {
			logrus.Errorln("Failed to unpack head:", err)
			break
		}
		if unpackHead.GetMessageLen() > 0 {
			msg := unpackHead.(*znet.Message)
			msg.SetData(data[headlen : headlen+msg.GetMessageLen()])
			logrus.Printf("MessageId: %d, MessageLen: %d,Data:%v\n", msg.GetMessageId(), msg.GetMessageLen(), string(msg.GetData()))
			data = data[headlen+msg.GetMessageLen():]
		} else {
			break
		}
	}
	// 发送响应
	responseMsg := new(znet.Message)
	responseMsg.SetMessageLen(uint32(len("Pong!")))
	responseMsg.SetMessageId(1)
	responseMsg.SetData([]byte("Pong!"))

	responsePack, err := znet.NewDataPack().Pack(responseMsg)
	if err != nil {
		logrus.Errorln("Failed to pack response:", err)
		return
	}

	_, err = request.GetConnection().GetTCPConnection().Write(responsePack)
	if err != nil {
		logrus.Errorln("Failed to send response:", err)
		return
	}
}
func (prouter *PingRouter) PostHandler(request ziface.IRequest) {
	logrus.Println("Call Router PostHandler")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	// if err != nil {
	// 	logrus.Errorln(err)
	// }
}

func main() {
	s := znet.NewServer("ubuntu server")
	s.AddRoute(&PingRouter{})
	s.Serve()
}
