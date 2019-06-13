package model

import (
	"gocode/mQQ/common"
	"net"
)

type CurUser struct {
	Conn net.Conn
	User common.User
}
