package main

import (
	"strings"
	// "fmt"
	"net"
)
type Msg struct {
	from,to,msg,ip string
}

func parseMsg(msg string) Msg{
	s := strings.Split(msg, ";")
	// fmt.Println("message_parse:", s)
	message := Msg{s[0], s[1], s[2], s[3]}
	// fmt.Println("message_parse_struct", message)
	return message
}

func combinationMsg(nickName string, msg string, conn net.Conn) string{
	localAddr := conn.LocalAddr().String()
	// fmt.Println("localAddr:",localAddr)
	message := nickName+";"+msg+";"+localAddr;
	// fmt.Println("message_combination", message)
	return message
}