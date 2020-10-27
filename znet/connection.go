/*
 * @Author: 光城
 * @Date: 2020-10-22 15:30:56
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 17:33:19
 * @Description:
 * @FilePath: /Zinx_Learning/znet/connection.go
 */
package znet

import (
	"errors"
	"fmt"
	"io"
	"net"

	"light.com/guangcheng/ziface"
)

/*
 连接模块
*/

// 连接模块
type Connection struct {
	// socket TCP套接字
	Conn *net.TCPConn
	// 连接的ID
	ConnID uint32
	// 当前连接的状态(是否已经关闭)
	IsClosed bool
	// 等待连接被动退出的channel
	ExitChan chan bool // 由reader告知writer退出
	// 无缓冲 管道 用于读、写goroutine之间的消息通信
	MsgChan chan []byte
	// 该连接处理的方法Router
	MsgHandler ziface.IMsgHandler
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		IsClosed:   false,
		MsgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println("connID = ", c.ConnID, "[Reader is exit, remote addr is", c.RemoteAddr().String()+"]")
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中， 最大MaxPackageSize字节
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())

		// 读端断连 EOF
		if _, err := io.ReadFull(c.GetTcpConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		msg, err := dp.Unpack(headData)

		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetData(data)

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		// 执行注册的路由方法
		go c.MsgHandler.DoMsgHandler(&req)
	}
}

// 写消息goroutine 专门发送给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println("connID = ", c.ConnID, "[Writer is exit, remote addr is", c.RemoteAddr().String()+"]")
	// 不断阻塞等待channel的消息,进行回写给客户端
	for {
		select {
		case data := <-c.MsgChan:
			// 有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error,", err)
				return
			}
		case <-c.ExitChan: // 可读
			// 代表Reader已经退出，此时Writer可以退出
			return
		}
	}
}
func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	// TODO 启动从当前连接读数据的业务
	go c.StartReader()
	go c.StartWriter()
	// TODO 启动从当前连接写数据的业务
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)
	if c.IsClosed == true {
		return
	}
	c.IsClosed = false
	c.Conn.Close()
	c.ExitChan <- true // 告知Writer关闭
	close(c.ExitChan)
	close(c.MsgChan)
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

// 提供一个SendMsg方法 将我们要发送给客户端的数据，先进性封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.IsClosed == true {
		return errors.New("Connection closed when send msg")
	}

	// 将data进行封包 |MsgDataLen|MsgID|Data|))
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id=", msgId)
		return errors.New("Pack error msg")
	}
	c.MsgChan <- binaryMsg // 数据发送给Chan
	return nil
}
