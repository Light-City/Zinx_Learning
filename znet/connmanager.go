/*
 * @Author: 光城
 * @Date: 2020-10-28 16:02:06
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 16:42:30
 * @Description: 连接管理
 * @FilePath: /Zinx_Learning/znet/connmanager.go
 */
package znet

import (
	"errors"
	"fmt"
	"sync"

	"light.com/guangcheng/ziface"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection // 管理的连接集合
	connLock    sync.RWMutex                  // 保护连接集合的读写锁
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	// 保护共享资源map  加写锁  写操作加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	// 将conn加入到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("connID=", conn.GetConnID(), "connection add to ConnManager succ")
}
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("connID=", conn.GetConnID(), "connection delete to ConnManager succ conn num=", connMgr.Len())
}
func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	// 获取加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.Unlock()
	if conn, ok := connMgr.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not FOUND!")
	}
}
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}
func (connMgr *ConnManager) ClearConn() {
	// 写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	// 删除并停止
	for connID, conn := range connMgr.connections {
		// 停止
		conn.Stop()
		// 删除
		delete(connMgr.connections, connID)
	}
	fmt.Println("Clear All conections succ! conn num=", connMgr.Len())
}
