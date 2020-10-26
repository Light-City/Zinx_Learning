/*
 * @Author: 光城
 * @Date: 2020-10-26 10:37:19
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-26 10:43:08
 * @Description:
 * @FilePath: /Zinx_Learning/znet/request.go
 */
package znet

import "light.com/guangcheng/ziface"

type Request struct {
	// 已经和客户端建立好的连接
	conn ziface.IConnection
	// 客户端请求的数据
	data []byte
}

// 得到当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 得到请求的数据
func (r *Request) GetData() []byte {
	return r.data
}
