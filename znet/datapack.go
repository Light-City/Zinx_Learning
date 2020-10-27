/*
 * @Description: 封包拆包实现:解决TCP粘包
 * @Autor: 光城
 * @Date: 2020-10-27 08:19:03
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-27 11:09:47
 * @FilePath: \Zinx_Learning\znet\datapack.go
 */
package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"light.com/guangcheng/utils"
	"light.com/guangcheng/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4字节) ID uint32 (4字节)
	return 8
}

/**
 * @description: |datalen|msgID|data|
 * @param {*DataPack} dp
 * @return {*}
 */
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	// 将DataLen写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// 将MsgId写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 将Data写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

/**
 * @description: 拆包方法(将包Head信息读出来) 之后再根据head信息里的dataLen再进行一次读
 * @param {*DataPack} dp
 * @return {*}
 */
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断ataLen是否超过了我们允许的最大长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recv!")
	}
	return msg, nil
}
