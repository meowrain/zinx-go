package ziface

import "net"

// 定义链接模块的抽象层
type IConnection interface {
	//启动链接,让当前的连接准备开始工作
	Open()

	//关闭连接，结束当前连接的工作
	Close()

	//获取当前连接的绑定的socket conn
	GetTCPConnection() *net.TCPConn

	// 获取当前链接模块的链接id
	GetConnID() uint32

	//获取远程客户端的TCP状态IP和端口
	GetRemoteAddr() net.Addr

	//发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error

//*net.TCPConn 表示要处理的连接, []byte表示要处理的数据, int表示当前连接要处理数据的长度
