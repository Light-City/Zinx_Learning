/*
 * @Author: 光城
 * @Date: 2020-10-26 10:37:12
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 10:39:29
 * @Description: 客户端请求的连接信息，和请求数据包装到一个Request
 * @FilePath: \Zinx_Learning\ziface\irequest.go
 */
package ziface

type IRequest interface {
	// 得到当前连接
	GetConnection() IConnection
	// 得到请求的消息数据
	GetData() []byte
	GetMsgID() uint32
}
