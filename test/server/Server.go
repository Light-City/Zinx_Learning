/*
 * @Author: 光城
 * @Date: 2020-10-22 11:18:14
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-26 11:34:08
 * @Description:
 * @FilePath: /Zinx_Learning/test/server/Server.go
 */
package main

import (
	"fmt"

	"light.com/guangcheng/ziface"
	"light.com/guangcheng/znet"
)

type PingRouter struct {
	znet.BaseRoute
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping error")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.AddRouter(&PingRouter{})
	s.Server()
}
