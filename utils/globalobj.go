/*
 * @Description:
 * @Autor: 光城
 * @Date: 2020-10-26 20:26:00
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 09:08:14
 * @FilePath: \Zinx_Learning\utils\globalobj.go
 */
package utils

import (
	"encoding/json"
	"io/ioutil"

	"light.com/guangcheng/ziface"
)

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
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	// 应该从配置文件中加载
	// GlobalObject.Reload()
}
