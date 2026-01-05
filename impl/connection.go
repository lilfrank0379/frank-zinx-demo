package impl

import (
	"fmt"
	"frank-zinx-demo/iface"
	"net"
)

/*
  @author lilfrank
  @date   2025/12/25 12:20
*/

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	// 当前连接绑定的处理业务的方法API
	handleAPI iface.HandleFunc
	// 告知当前连接已经退出（停止）的管道
	ExitChan chan bool
}

// NewConnection 初始化Connection
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI iface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

// StartReader 连接的读数据的业务
func (c *Connection) StartReader() {
	fmt.Println("connID = ", c.ConnID, " reader goroutine is running")
	defer fmt.Println("connID = ", c.ConnID, " reader goroutine is closed", ",remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf error:", err)
			continue
		}

		// 调用当前连接所绑定的handleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID = ", c.ConnID, " handle api error:", err)
			break
		}
	}
}

// Start 启动连接，让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("conn start... connID = ", c.ConnID)

	// 启动从当前连接的读数据的业务
	go c.StartReader()

	// todo 启动从当前连接的写数据的业务
}

// Stop 停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("conn stop... connID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	// 关闭socket连接
	c.Conn.Close()
	// 回收资源，关闭管道
	close(c.ExitChan)
}

// GetTCPConnection 获取当前连接的绑定socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端的TCP状态，IP，端口
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
