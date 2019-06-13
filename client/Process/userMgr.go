package Process

import (
	"fmt"
	"gocode/mQQ/client/model"
	"gocode/mQQ/common"
)

var onlineUser map[int]common.User
var curUser model.CurUser

func init() {
	onlineUser = make(map[int]common.User, 10)
}
func outputOnlineUser() {
	fmt.Println("当前用户在线列表:")
	for id, _ := range onlineUser {
		fmt.Println("用户id:\t", id)
	}
}
func updataUserStatus(notifymes common.NotifyUserStatusMes) {
	user, ok := onlineUser[notifymes.UserConnId]
	if !ok {
		user = common.User{
			UserId: notifymes.UserConnId,
		}
	}
	user.UserStatus = notifymes.Status
	onlineUser[notifymes.UserConnId] = user
	outputOnlineUser()
}
