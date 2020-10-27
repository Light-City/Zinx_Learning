/*
 * @Description: 消息封装
 * @Autor: 光城
 * @Date: 2020-10-27 08:05:42
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 08:09:18
 * @FilePath: \Zinx_Learning\ziface\imessage.go
 */
package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetMsgLen(uint32)
	SetData([]byte)
}
