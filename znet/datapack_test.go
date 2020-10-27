/*
 * @Description: 封包拆包实现:解决TCP粘包
 * @Autor: 光城
 * @Date: 2020-10-27 08:19:03
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 10:13:41
 * @FilePath: \Zinx_Learning\znet\datapack_test.go
 */

package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack_Pack(t *testing.T) {
	listerner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Server listerner err:", err)
	}

	go func() {
		for {
			conn, err := listerner.Accept()
			if err != nil {
				fmt.Println("Server accept err:", err)
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					// ReadFull 可以读取任意对象的数据(该对象必须是实现Read方法)
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error", err)
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack error", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data error", err)
							return
						}
						fmt.Println("-> Recv MsgID:", msg.Id, ",datalen=", msg.DataLen, ",data=", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("client dial err: ", err)
		return
	}
	dp := NewDataPack()
	// 模拟粘包过程 封装两个msg一同发送
	// 封装第一个msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte("zinx"),
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}
	// 封装第二个msg2包
	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error", err)
	}
	// 粘包msg1与msg2
	sendData := append(sendData1, sendData2...)
	// 一次性发送给服务端
	conn.Write(sendData)
	select {}
}
