package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
	"zinx/zutils"

	"github.com/sirupsen/logrus"
)

// IServer接口实现
type Server struct {
	Name      string         //服务器名称
	IP        string         //服务器绑定的ip地址
	IPVersion string         //服务器绑定的ip版本
	Port      int            //服务器绑定的端口
	Router    ziface.IRouter // 该链接处理的方法Router
}

func init() {
	// 配置logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,                  // 强制彩色输出
		FullTimestamp:   true,                  // 显示完整时间戳
		TimestampFormat: "2006-01-02 15:04:05", // 设置时间戳格式
	})
	// 设置logrus输出到标准输出
	logrus.SetOutput(logrus.StandardLogger().Out)
	// 设置logrus日志级别
	logrus.SetLevel(logrus.InfoLevel)
}

func (server *Server) Start() {
	//TODO: 启动服务
	tcpAddr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		logrus.Errorln("ResolveTCPAddr failed, err:", err)
		return
	}
	listener, err := net.ListenTCP(server.IPVersion, tcpAddr)
	if err != nil {
		logrus.Errorln("ListenTCP failed, err:", err)
		return
	}
	logrus.Infof("Start zinx successfully!,server is running at %s:%d", server.IP, server.Port)
	var cid uint32 = 0

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logrus.Errorln("Accept failed, err:", err)
			return
		}
		dealConn := NewConnection(conn, cid, server.Router) //*Connection
		go dealConn.Open()
		cid++
	}
}
func (server *Server) Stop() {
	//TODO: 停止服务
}
func (server *Server) Serve() {
	//TODO: 监听服务
	go server.Start()

	//中间执行一些其它的业务
	select {}

}
func (server *Server) AddRoute(router ziface.IRouter) {
	server.Router = router
	logrus.Infoln("AddRoute successfully!")
}
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8080,
		Router:    nil,
	}
	zutils.GlobalObject.TcpServer = s
	logrus.Infoln("Loaded GlobalObject:", zutils.GlobalObject)
	return s

}
