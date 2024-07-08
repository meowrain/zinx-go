package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/ziface"

	"github.com/sirupsen/logrus"
)

// IServer接口实现
type Server struct {
	Name      string //服务器名称
	IP        string //服务器绑定的ip地址
	IPVersion string //服务器绑定的ip版本
	Port      int    //服务器绑定的端口
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
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logrus.Errorln("Accept failed, err:", err)
			continue
		}
		go func() {
			// 关闭连接
			defer conn.Close()
			//TODO: 业务处理
			logrus.Infoln("Accept a new connection from", conn.RemoteAddr())
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						// 连接正常关闭，不是错误
						logrus.Infoln("Connection closed by client")
						break
					}
					// 读取过程中遇到其他错误
					logrus.Errorln("Read failed, err:", err)
					break
				}
				logrus.Infoln("Read", n, "bytes from", conn.RemoteAddr())
				fmt.Println("recv:", string(buf[:n]))
				// 发送数据
				if _, err = conn.Write(buf[:n]); err != nil {
					logrus.Errorln("Write failed, err:", err)
				}
				logrus.Infoln("Write to", conn.RemoteAddr(), "success")
			}
		}()
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
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8080,
	}
	return s

}
