<!--
 * @Author: 光城
 * @Date: 2020-10-22 15:24:14
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 14:24:12
 * @Description:
 * @FilePath: /Zinx_Learning/Readme.md
-->
## 1.V0.1 基础server
- 方法
启动服务器、停止服务器、运行服务器
初始化
- 属性
name名称、监听的IP、监听的端口

## 2.V0.2 简单的连接和业务封装
连接的模块
- 方法
启动连接、停止连接、获取当前连接的conn对象(套接字)、得到连接ID、得到客户端连接的地址和端口、发送数据
连接所绑定的处理业务函数
- 属性
socket TCP套接字、连接的ID、当前连接的状态(是否已经关闭)、与当前连接所绑定的处理业务、等待连接被动退出的channel

## 3.V0.3 基础Route模块

- Request请求封装
  - 将连接与数据和绑定到一起
    - 属性
    连接Connection、请求数据
    - 方法
    得到当前连接、得到当前数据
- Route模块
  - 抽象的Router
    - 处理业务之前的方法
    - 处理业务的主方法
    - 处理业务之后的方法
  - 具体BaseRoute
    - 处理业务之前的方法
    - 处理业务的主方法
    - 处理业务之后的方法
- zinx集成route模块
  - IServer增添路由添加功能
  - Server类增添Route成员
  - Connection类绑定一个Route成员
  - 在Connection调用已经注册的Route处理业务

## 4.V0.4 Zinx的全局配置

  zinx.json用户填写

```json
{
  "Name" : "demo server",
  "Host" : "127.0.0.1",
  "TcpPort" : 7777,
  "MaxConn" : 3
}
```
  - 创建全局参数文件
  init 读取用户配置好的zinx.json文件,保存到全局对象中
  - 硬参数替换与Server初始化参数配置
  将zinx框架种全部的硬代码，用全局对象的参数进行替换

## 5.V0.5 消息封装

- Message
  - 属性
  消息ID、消息长度、数据内容
  - 方法
  Setter、Getter
- 解决TCP粘包
消息TLV序列化 解决TCP粘包问题
head:Len+id 8字节
body:Data

TCP传输是stream,没有数据尾巴,需要在应用层判断这个包在什么时候截止
  - 针对Message进行TLV格式的封装
    - Len+ID+Data
  - 针对Message进行TLV格式的拆包
    - 先读取8字节head，再取偏移拿Data

- 消息封装机制继承在Zinx框架中

## 6.V0.6 多路由模式

- 消息管理模块(支持多路由业务api调度管理)
  - 属性
  集合-消息ID和对应的router的关系map存储
  - 方法
  根据msgID来索引调度路由方法、添加方法到map集合中