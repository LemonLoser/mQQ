package process

import (
	"encoding/json"
	"fmt"
	"gocode/mQQ/common"
	"gocode/mQQ/server/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes common.Message) {
	//取出mes中的内容
	var sms common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), sms)
	if err != nil {
		fmt.Println("unmarshal failed")
		return
	}
	data, err := json.Marshal(sms)
	if err != nil {
		fmt.Println("marshal failed")
	}
	//遍历服务器端的onlineuser
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == sms.User.UserId {
			continue
		}
		this.SendToEachOnlineUser(data, up.Conn)
	}
}
func (this *SmsProcess) SendToEachOnlineUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败", err)
		return
	}
}
