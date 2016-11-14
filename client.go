package main

import(
    "net"
    "fmt"
)

func request(msg string, conn net.Conn, nickName string) {
	message := combinationMsg(nickName, msg, conn)
	if message == "" {
		fmt.Println("请按照【接收人;信息】格式,正确填写:")
		return 
	}
	_, err := conn.Write([]byte(""+message+"\r\n\r\n"))
	if err != nil {
		fmt.Println("error_msg_1:",err)
		return 
	}
}

func response(conn net.Conn) {
	buf    := make([]byte, 1024)
	serverMsg, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error_msg",err)
		return 
	}
	fmt.Println(string(buf[:serverMsg]))
}