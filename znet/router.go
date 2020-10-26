/*
 * @Author: 光城
 * @Date: 2020-10-26 10:49:19
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-26 11:06:32
 * @Description:
 * @FilePath: /Zinx_Learning/znet/route.go
 */
package znet

import "light.com/guangcheng/ziface"

// 实现route时，先嵌入BaseRoute基类，然后根据需要对基类的方法进行重写就好了
type BaseRoute struct {
}

// 这里之所以BaseRequest的方法都为空是因为有的Route不希望有PreHandle、PostHandle这两个业务
// 所以Route全部继承BaseRoute的好处就是不需要实现PreHandle、PostHandle

// 之前
func (br *BaseRoute) PreHandle(request ziface.IRequest) {}

// 主
func (br *BaseRoute) Handle(request ziface.IRequest) {}

// 之后
func (br *BaseRoute) PostHandle(request ziface.IRequest) {}
