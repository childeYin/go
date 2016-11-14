package main

import (
	"strings"
	// "fmt"
	"net"
)
type Msg struct {
	fromNickName, to, msg string
}

func parseMsg(msg string) Msg{
	s := strings.Split(msg, ";")
	// fmt.Println("message_parse:", s)
	message := Msg{s[0], s[1], s[2]}
	// fmt.Println("message_parse_struct", message)
	return message
}

func combinationMsg(nickName string, msg string, conn net.Conn) string{
	s := strings.Split(msg, ";")
	length := len(s)
	if length < 2 {
		return ""
	}
	message := nickName+";"+msg;
	// fmt.Println("message_combination", message)
	return message
}