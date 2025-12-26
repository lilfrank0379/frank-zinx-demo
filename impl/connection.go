package impl

import (
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
