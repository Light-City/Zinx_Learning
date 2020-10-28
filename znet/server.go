/*
 * @Author: 光城
 * @Date: 2020-10-22 15:34:58
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 10:17:38
 * @Description:
 * @FilePath: /Zinx_Learning/znet/server.go
 */
package znet

import (
	"fmt"
	"net"

	"light.com/guangcheng/utils"
	"light.com/guangcheng/ziface"
)

// IServer的接口实现，定义一个server的服务器模块
type Server struct {
	// 名称
	Name string
	// IP版本
	IPVersion string
	// IP地址
	IP string
	// 端口
	Port int
	// 当前Server的消息管理模块,用来绑定MsgID和对应的处理业务API关系
	MsgHandler ziface.IMsgHandler
}

// 启动
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, listerner at IP : %s, Port: %d is starting\n", utils.GlobalObject.Name,
		utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version :%s, MaxConn:%d, MaxPacketSize:%d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[Start Server Listenner at IP: %s, Port: %d, is starting]\n", s.IP, s.Port)

	// 异步 防止后面read阻塞
	go func() {
		// 0.开启消息队列及Worker工作池
		s.MsgHandler.StartWorkerPool()

		// 1.获取一个TCP的Addr 创建套接字
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}
		// 2.监听服务器的地址 listen
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err", err)
			return
		}
		fmt.Println("start Zinx server succ ", s.Name, "listening...")
		var cid uint32
		cid = 0
		// 3.阻塞等待客户端连接，处理客户端连接业务(读写)
		for {
			// 如果有客户端连接过来,阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 将该处理新连接的业务方法和conn进行绑定 得到我们的连接模块
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()
}

// 停止
func (s *Server) Stop() {
	// TODO 释放服务器资源、状态、连接信息，进行停止或回收
}

// 运行
func (s *Server) Server() {
	// 启动
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

// 路由
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add router succ!")
}

/*
	初始化Server模块的方法
*/
func NewServer(name string) ziface.IServer {

	s := &Server{
		Name:       utils.GlobalObject.Name, // 导包 "light.com/guangcheng/utils" 里面init方法会默认执行
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
	}
	return s
}
