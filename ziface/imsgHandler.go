/*
 * @Author: 光城
 * @Date: 2020-10-27 14:24:56
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 10:17:25
 * @Description: 消息管理抽象层
 * @FilePath: /Zinx_Learning/ziface/imsgHandler.go
 */
package ziface

type IMsgHandler interface {
	// 调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)
	// 为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)
	StartWorkerPool()
	StartOneWorker(workerID int, taskQueue chan IRequest)
	SendMsgToTaskQueue(request IRequest)
}
