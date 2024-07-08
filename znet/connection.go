package znet

import (
	"errors"
	"io"
	"net"
	"zinx/ziface"

	"github.com/sirupsen/logrus"
)

type Connection struct {
	Conn     *net.TCPConn   // TCP套接字
	ConnID   uint32         // 链接的ID
	IsClosed bool           // 当前链接的状态
	ExitChan chan bool      // 等待链接被动退出
	Router   ziface.IRouter //该链接处理的方法
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	return c
}

func (conn *Connection) StartReader() {
	logrus.Infoln("Reader Goroutine is running...")
	defer logrus.Infoln("connID =", conn.GetConnID(), "Reader Goroutine is exiting...Remote Addr is", conn.GetRemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		_, err := conn.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				logrus.Infoln("connID =", conn.GetConnID(), "Read Error: EOF")
				conn.ExitChan <- true
				return
			}
			// logrus.Errorln("connID =", conn.GetConnID(), "Read Error:", err)
			conn.ExitChan <- true
			return
		}
		//从路由中找到注册绑定的connection对应的router
		//得到当前conn数据的Request请求数据
		req := &Request{conn: conn, data: buf}
		go func(req ziface.IRequest) {
			conn.Router.PreHandler(&req)
			conn.Router.Handler(&req)
			conn.Router.PostHandler(&req)
		}(req)
		// if err = conn.handleAPI(conn.Conn, buf, n); err != nil {
		// 	logrus.Errorln("connID =", conn.GetConnID(), "HandleAPI Error:", err)
		// 	conn.ExitChan <- true
		// 	return
		// }
	}
}

func (conn *Connection) Open() {
	logrus.Infoln("Connection Open().. ConnID =", conn.GetConnID())
	go conn.StartReader()

	// 等待 ExitChan 信号
	<-conn.ExitChan
	logrus.Infoln("Connection closed by ExitChan signal, ConnID =", conn.GetConnID())
	conn.Close()
}

func (conn *Connection) Close() {
	logrus.Infoln("Conn Close().. ConnID =", conn.GetConnID())
	if conn.IsClosed {
		return
	}
	conn.IsClosed = true
	conn.GetTCPConnection().Close()
	close(conn.ExitChan)
}

func (conn *Connection) Send(data []byte) error {
	if conn.IsClosed {
		return errors.New("connection is closed")
	}
	_, err := conn.Conn.Write(data)
	return err
}

func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}

func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}

func (conn *Connection) GetRemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}
