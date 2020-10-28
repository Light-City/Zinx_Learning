/*
 * @Description:
 * @Autor: 光城
 * @Date: 2020-10-26 20:26:00
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 09:50:23
 * @FilePath: /Zinx_Learning/utils/globalobj.go
 */
package utils

import (
	"encoding/json"
	"io/ioutil"

	"light.com/guangcheng/ziface"
)

var FileSwitch string

/* 存储一切有关zinx框架的全局参数，供其他模块使用
 */

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string
	/*
		Zinx
	*/
	Version        string
	MaxConn        int
	MaxPackageSize uint32
	WorkerPoolSize uint32 // 当前业务工作Worker池的Goroutine数量
	// 每个worker对应的消息队列的任务的数量最大值
	MaxWorkerTaskLen uint32 // Zinx框架允许用户最多开辟多少个Worker(限定条件)
}

/*
	定义一个全局的对外Globalobj
*/
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	// 将json文件解析到struct中
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/**
 * @description: 提供给一个init方法，初始化当前的Glob:lObject
 * @param {*}
 * @return {*}
 */
func init() {
	// 如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}
	// 应该从配置文件中加载
	GlobalObject.Reload()
}
