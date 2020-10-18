package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// 1.连接 得到一个conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err ", err)
		return
	}
	// 2. 连接调用write写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1."))
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err ", err)
			return
		}
		fmt.Printf("server call back: %s, cnt=%d\n", buf, cnt)
		// cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
