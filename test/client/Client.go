/*
 * @Description:
 * @Autor: 光城
 * @Date: 2020-10-27 09:35:45
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 11:18:12
 * @FilePath: \Zinx_Learning\test\client\Client.go
 */
package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"light.com/guangcheng/znet"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// 1.连接 得到一个conn
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err ", err)
		return
	}
	// 2. 连接调用write写数据
	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.5 client Test Message")))
		if err != nil {
			fmt.Println("pack err ", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write err ", err)
			return
		}
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head err ", err)
			break
		}
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack head err ", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err ", err)
				break
			}
			fmt.Println("-> Recv Server Msg: ID=", msg.Id, "len=", msg.DataLen, "data=", string(msg.Data))
		}

		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
