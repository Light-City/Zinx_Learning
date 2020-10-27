/*
 * @Author: 光城
 * @Date: 2020-10-22 11:18:14
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 11:07:50
 * @Description:
 * @FilePath: \Zinx_Learning\test\server\Server.go
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
	fmt.Println("Call Router Handle...")

	fmt.Println("recv from client:msgID=", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping error", err)
	}
}

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.AddRouter(&PingRouter{})
	s.Server()
}
