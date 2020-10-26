/*
 * @Author: 光城
 * @Date: 2020-10-22 16:51:06
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-26 11:22:19
 * @Description:
 * @FilePath: /Zinx_Learning/ziface/iserver.go
 */
package ziface

// 定义一个服务器接口
type IServer interface {
	// 启动
	Start()
	// 停止
	Stop()
	// 运行
	Server()
	// 路由功能：给当前服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(router IRouter)
}
