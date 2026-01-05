package impl

import (
	"fmt"
	"frank-zinx-demo/iface"
	"net"
)

/*
  @author lilfrank
  @date   2025/12/23 15:52
*/

// Server 是IServer的实现
type Server struct {
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	IP        string
	Port      int
}

// NewServer 初始化Server
func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8686,
	}
	return s
}

// CallbackToClient 定义当前客户端连接的所绑定的handleAPI（目前先写死，后续由用户自定义handle方法）
func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显数据
	fmt.Println("[conn handle] CallbackToClient")
	fmt.Println("server write:", string(data)+"\n")

	for len(data) > 0 {
		n, err := conn.Write(data)
		if err != nil {
			fmt.Println("write error:", err)
			continue
		}
		data = data[n:] // 剩余部分继续写入
	}

	return nil
}

// Start 启动服务器
func (server *Server) Start() {
	go func() {
		// 1 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		// 2 监听服务器的地址
		listener, err := net.ListenTCP(server.IPVersion, addr)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		// 3 阻塞，等待客户端连接，处理客户端连接的业务（读写）
		var cid uint32
		cid = 1

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			dealConn := NewConnection(conn, cid, CallbackToClient)
			cid++
			dealConn.Start() // 这里可以用协程启动
		}
	}()
}

// Stop 停止服务器
func (server *Server) Stop() {

}

// Serve 运行服务器
func (server *Server) Serve() {
	server.Start()
	select {}
}
