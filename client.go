package main

import(
    "net"
    "fmt"
)

func request(msg string, conn net.Conn, nickName string) {
	message := combinationMsg(nickName, msg, conn)
	parseMsg(message)
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
	fmt.Println("serverMsg:",string(buf[:serverMsg]))
}