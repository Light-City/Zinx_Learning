/*
 * @Author: 光城
 * @Date: 2020-10-26 10:46:16
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-26 11:10:56
 * @Description: 路由抽象接口 路由里的数据都是IRequest
 * @FilePath: /Zinx_Learning/ziface/irouter.go
 */
package ziface

type IRouter interface {
	// 之前
	PreHandle(request IRequest)
	// 主
	Handle(request IRequest)
	// 之后
	PostHandle(request IRequest)
}
