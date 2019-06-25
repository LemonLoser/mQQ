package Process

import (
	"encoding/json"
	"fmt"
	"gocode/mQQ/common"
)

func outputGroupMes(mes common.Message) {
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), smsMes)
	if err != nil {
		fmt.Println("unmarshal failded", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.User, smsMes.Content)
	fmt.Println(info)
}
