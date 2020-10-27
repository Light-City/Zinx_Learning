/*
 * @Author: 光城
 * @Date: 2020-10-22 15:30:49
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 11:01:45
 * @Description: 连接的模块方法
 * @FilePath: \Zinx_Learning\ziface\iconnection.go
 */
package ziface

import "net"

// 定义连接的抽象层
type IConnection interface {
	// 启动连接  让当前连接准备开始工作
	Start()
	// 停止连接
	Stop()
	// 获取当前连接的绑定socket conn
	GetTcpConnection() *net.TCPConn
	// 获取当前连接模块的连接ID
	GetConnID() uint32
	// 获取远程客户端的TCP状态 IP Port
	RemoteAddr() net.Addr
	// 发送数据 将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error
}

/**
 * @description: 定义一个处理连接业务的方法
 * @param {*} net tcp连接套接字
 * @param {*} byte 数据类型
 * @param {*} int 数据长度
 * @return {*} error
 */
type HandleFunc func(*net.TCPConn, []byte, int) error
