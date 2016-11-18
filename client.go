package main

import(
    "net"
    "fmt"
    "os"
    "bufio"
    "sync"
    "strings"

)

func request(msg string, conn net.Conn, nickName string) {
	message := combinationMsg(nickName, msg)
	if message == "" {
		fmt.Println("请按照【接收人;信息】或者【接收人；信息】格式,正确填写:")
		return 
	}
	fmt.Println("conn is ", conn)
	length, err := conn.Write([]byte(message+"\r\n"))
	fmt.Println("conn write", length)
	if err != nil {
		fmt.Println("request server error msg:",err)
		return 
	}
}
func autoRequest(conn net.Conn, nickName string) {
	// fmt.Println(nickName,"----")
	message := nickName+";欢迎回来"
	message = combinationMsg(nickName, message)
	conn.Write([]byte(message+"\r\n"))
}
func requests(conn net.Conn, wg *sync.WaitGroup, nickName string) {
	defer wg.Done()
	for {
		fmt.Println("请输入信息: ")
        reader     := bufio.NewReader(os.Stdin)
        data, _, _ := reader.ReadLine()
        message    := string(data)

		message = combinationMsg(nickName, message)
		if message == "" {
			fmt.Println("请按照【接收人;信息】或者【接收人；信息】格式,正确填写:")
			continue 
		}
		// fmt.Println("conn is ", conn)
		_, err := conn.Write([]byte(message+"\r\n"))
		// fmt.Println("conn write", length)
		if err != nil {
			fmt.Println("request server error msg:",err)
			continue 
		}
	}
}

func responses(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		buf := make([]byte, 5000)
		serverMsg, err := conn.Read(buf)
		// fmt.Println(serverMsg)
		if err != nil {
			fmt.Println("server response error msg",err)
			os.Exit(1)
		}
		receiveMsg := string(buf[:serverMsg])
		flag 	   := strings.Contains(receiveMsg, "下线通知")
		if flag {
			conn.Close()
			fmt.Println(receiveMsg)
			os.Exit(1)
		} else {
			fmt.Println(receiveMsg)
		}
		// fmt.Println(string(buf[:serverMsg]))
	}
}

func response(conn net.Conn) {
	buf    := make([]byte, 5000)
    // io.Copy(os.Stdout, conn)
    // fmt.Println(os.Stdout)
	serverMsg, err := conn.Read(buf)
	if err != nil {
		fmt.Println("server response error msg",err)
		return 
	}
	fmt.Println(string(buf[:serverMsg]))
}