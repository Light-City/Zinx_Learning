/*
 * @Description: 封包拆包实现:解决TCP粘包
 * @Autor: 光城
 * @Date: 2020-10-27 08:18:56
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 08:24:20
 * @FilePath: \Zinx_Learning\ziface\idatapack.go
 */
package ziface

type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
