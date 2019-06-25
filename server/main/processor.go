package main

import (
	"fmt"
	"gocode/mQQ/common"
	process2 "gocode/mQQ/server/process"
	utils2 "gocode/mQQ/server/utils"
	"io"
	"net"
)

//创建一个结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//根据client发送的消息种类不同,决定调用哪个函数
func (this *Processor) ServerProcessMes(mes common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		ups := process2.UserProcess{
			Conn: this.Conn,
		}
		err = ups.ServerProcessLogin(mes)
	case common.RegisterMesType:
		ups := process2.UserProcess{
			Conn: this.Conn,
		}
		err = ups.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println(err)
			return
		}
	case common.SmsType:
		//创建一个smsprocess的实例完成转发群聊消息的任务
		smsProcess := process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("message type is not found...can not handle...")
		return
	}
	return
}

func (this *Processor) Processes() (err error) {
	//循环读取链中的数据包
	for {
		utils := utils2.Transfer{
			Conn: this.Conn,
		}
		mes, err := utils.ReadPack()
		fmt.Println("读包正在进行")
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit,server quite")
				return err
			} else {
				fmt.Println("readPack fail err", err)
				return err
			}
		}
		//fmt.Println("mes=",mes)
		err = this.ServerProcessMes(mes)
		if err != nil {
			return err
		}
	}
}
