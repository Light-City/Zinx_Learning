/*
 * @Author: 光城
 * @Date: 2020-10-22 11:18:14
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 17:24:53
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

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle...")

	fmt.Println("recv from client:msgID=", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping error", err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRoute
}

// Test Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle...")

	fmt.Println("recv from client:msgID=", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello zinx\n"))
	if err != nil {
		fmt.Println("call back hello zinx error", err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("=>DoConnectionBegin is Called.........")

	if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}
}
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("=>DoConnectionLost is Called.........")
	fmt.Println("conn ID = ", conn.GetConnID(), "is lost.......")
}

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Server()
}
