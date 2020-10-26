/*
 * @Author: 光城
 * @Date: 2020-10-22 15:30:56
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-26 20:57:22
 * @Description:
 * @FilePath: \Zinx_Learning\znet\connection.go
 */
package znet

/*
 连接模块
*/
import (
	"fmt"
	"net"

	"light.com/guangcheng/utils"
	"light.com/guangcheng/ziface"
)

// 连接模块
type Connection struct {
	// socket TCP套接字
	Conn *net.TCPConn
	// 连接的ID
	ConnID uint32
	// 当前连接的状态(是否已经关闭)
	isClosed bool
	// 等待连接被动退出的channel
	ExitChan chan bool
	// 该连接处理的方法Router
	Router ziface.IRouter
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is", c.RemoteAddr().String)
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中， 最大MaxPackageSize字节
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", buf)
			continue
		}

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		// 执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	// TODO 启动从当前连接读数据的业务
	go c.StartReader()
	// TODO 启动从当前连接写数据的业务
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = false
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Connection) Send(data []byte) error {
	return nil
}
