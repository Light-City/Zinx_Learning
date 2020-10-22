<!--
 * @Author: 光城
 * @Date: 2020-10-22 15:24:14
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-22 16:05:26
 * @Description:
 * @FilePath: /Zinx_Learning/Readme.md
-->
V0.1 基础server
- 方法
启动服务器、停止服务器、运行服务器
初始化
- 属性
name名称、监听的IP、监听的端口
V0.2 简单的连接和业务封装
连接的模块
- 方法
启动连接、停止连接、获取当前连接的conn对象(套接字)、得到连接ID、得到客户端连接的地址和端口、发送数据
连接所绑定的处理业务函数
- 属性
socket TCP套接字、连接的ID、当前连接的状态(是否已经关闭)、与当前连接所绑定的处理业务、等待连接被动退出的channel