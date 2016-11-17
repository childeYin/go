package main

import (
	"strings"
	// "fmt"
)
type Msg struct {
	fromNickName, to, msg string
}

func parseMsg(msg string) Msg{
	s := strings.Split(msg, ";")
	length := len(s)
	message := Msg{}
	switch length {
		case 1:
			message = Msg{s[0], "", ""}
		case 2:
			message = Msg{s[0], s[1],  ""}
		case 3:
			message = Msg{s[0], s[1], s[2]}
	}
	// fmt.Println("message_parse:", s)
	// fmt.Println("message_parse_struct", message)
	return message
}

func combinationMsg(nickName string, msg string) string{
	newMsg  := strings.Replace(msg, "；", ";", -1)
	newMsgs := strings.Split(newMsg, ";")
	length  := len(newMsgs)
	if length < 2 {
		return ""
	}
	message := nickName+";"+newMsg;
	// fmt.Println("message_combination", message)
	return message
}