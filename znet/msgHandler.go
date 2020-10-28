/*
 * @Author: 光城
 * @Date: 2020-10-27 14:25:06
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 10:06:33
 * @Description: 消息管理+Worker工作池
 * @FilePath: /Zinx_Learning/znet/msgHandler.go
 */
package znet

import (
	"fmt"
	"strconv"

	"light.com/guangcheng/utils"
	"light.com/guangcheng/ziface"
)

type MsgHandler struct {
	// msgID:对应处理方法
	Apis map[uint32]ziface.IRouter
	// 负责worker读取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作worker池的数量
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize), // 每个Worker对应一个TaskQueue
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

// 启动一个Worker工作池 全局开启一次
func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 当前的worker对应的channel消息队列 开辟空间 第0个worker就用第0个channel...
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 启动当前的worker,阻塞等待消息从channel传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("worker ID=", workerID, "is started...")

	// 不断的阻塞等待对应的消息队列的消息
	for {
		select {
		// 如果有消息过来，出列的就是一个客户端的Request,执行当前Request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// 将消息交给TaskQueue,由Worker进行处理
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	// 保证每个worker所收到的request任务是均衡(平均分配),让哪个worker去处理，只需要将这个request请求发送给对应的taskQueue即可
	// 一个连接对应一个request
	// 根据客户端创建的ConnID来进行分配
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(),
		"request MsgID=", request.GetMsgID(),
		"to WorkerID=", workerID)
	// 将消息直接发送给对应的channal
	mh.TaskQueue[workerID] <- request
}
