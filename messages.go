package main

import (
	"strings"
	// "fmt"
)
type Msg struct {
	fromNickName, instruction, to, msg string
}

// var instructionType = "login,quit,msg,search,del,add"
var instructionType = "login,quit,msg,search"

func parseMsg(msg string) Msg{
	s := strings.Split(msg, ";")
	length := len(s)
	message := Msg{}
	switch length {
		case 1:
			message = Msg{s[0], "msg", "", ""}
		case 2:
			message = Msg{s[0], s[1],  "", ""}
		case 3:
			message = Msg{s[0], s[1], s[2], ""}
		case 4:
			message = Msg{s[0], s[1], s[2], s[3]}
	}
	message.fromNickName = strings.TrimSpace(message.fromNickName)
	message.instruction  = strings.TrimSpace(message.instruction)
	message.to 			 = strings.TrimSpace(message.to)
	message.msg 		 = strings.TrimSpace(message.msg)
	if !(strings.Contains(instructionType, message.instruction)) {
		message.instruction = "msg"
	}
	return message
}

func combinationMsg(nickName string, msg string) string{
	newMsg  := strings.Replace(msg, "；", ";", -1)
	newMsgs := strings.Split(newMsg, ";")
	length  := len(newMsgs)
	message := ""
	switch newMsgs[0] {
		case "quit", "msg", "search":
			message = nickName+";"+newMsg;
		default :
			if length < 3 {
				message = nickName+";msg;"+newMsg;
			} else {
				message = nickName+";"+newMsg;
			}
	}
	return message
}