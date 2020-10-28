<!--
 * @Author: 光城
 * @Date: 2020-10-22 15:24:14
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-28 19:32:02
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

## 7.V0.7 读写协程分离
- 添加一个Reader和Writer之间通信的channel
- 添加一个Writer Goroutine
- Reader由之前直接发送给客户端改为发送给通信Channel
- 启动Reader和Writer一同工作

## 8.V0.8 消息队列及多任务

10w client连接 会有10w个reader(阻塞)、10w个writer(阻塞)
阻塞并不会占用CPU

10个(固定值)处理业务的goroutine,不管客户端请求数量有多大，CPU在调度go之间，只需要在10个go之间进行切换。

- 创建一个消息队列
- 创建多任务work的工作池并启动
- 将之前的发送消息，全部改成把消息发送给消息队列和worker工作池来处理

## 9.V0.9 连接管理

- 创建一个连接管理模块
  - 属性
  已经创建的Connection集合(map)、针对map的互斥锁
  - 方法
  添加、删除、根据ID查找对应的连接、总连接个数、清理全部连接
- 集成连接管理模块
  - 将ConnManager加入Server模块中
  - 每次成功与客户端建立连接后添加连接
  - 判断当前连接数量是否已经超出最大值MaxConn
  - 每次成功与客户端建立连接断开后删除连接
- hook方法
玩家一上线发送广播

- 属性
  该Server创建连接之后自动调用Hook函数
  该Server销毁连接之前自动调用Hook函数
- 方法
  注册调用OnConnStart(conn)、OnConnStop(conn)方法
- 在Conn创建之后调用OnConnStart、在Conn销毁之前调用OnConnStop

## 10.V1.0 连接属性

- 给Connection模块添加可以配置属性的功能
  - 属性
    连接属性集合map
    保护连接属性的锁
  - 方法
    设置连接属性、获取连接属性、移除连接属性