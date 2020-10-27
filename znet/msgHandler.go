/*
 * @Author: 光城
 * @Date: 2020-10-27 14:25:06
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 15:16:14
 * @Description:
 * @FilePath: /Zinx_Learning/znet/msgHandler.go
 */
package znet

import (
	"fmt"
	"strconv"

	"light.com/guangcheng/ziface"
)

type MsgHandler struct {
	// msgID:对应处理方法
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID=", request.GetMsgID(), "is NOT FOUND! Need register!")
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}
func (mh *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	// id已经注册
	if _, ok := mh.Apis[msgID]; ok {
		panic("repeat api, msgID=" + strconv.Itoa(int(msgID)))
	}

	mh.Apis[msgID] = router
	fmt.Println("add api msgID=", msgID, "succ!")
}
