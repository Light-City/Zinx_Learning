/*
 * @Author: 光城
 * @Date: 2020-10-28 15:55:20
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 16:00:55
 * @Description: 连接管理
 * @FilePath: /Zinx_Learning/ziface/iconnmanager.go
 */
package ziface

type IConnManager interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(connID uint32) (IConnection, error)
	Len() int
	ClearConn()
}
