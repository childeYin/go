package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func autoLoginRequest(conn net.Conn, email string) {
	// fmt.Println(email,"----")
	message := "login;" + email + ";欢迎回来"
	message = combinationMsg(email, message)
	conn.Write([]byte(message + "\r\n"))
}

func requests(conn net.Conn, wg *sync.WaitGroup, email string) {
	defer wg.Done()
	for {
		fmt.Println("请输入信息: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		message := string(data)

		message = combinationMsg(email, message)
		if message == "" {
			fmt.Println("请按照【接收人;信息】或者【接收人；信息】格式,正确填写:")
			continue
		}
		// fmt.Println("conn is ", conn)
		_, err := conn.Write([]byte(message + "\r\n"))
		// fmt.Println("conn write", length)
		if err != nil {
			fmt.Println("request server error msg:", err)
			continue
		}
	}
}

func quit(conn net.Conn, email string) {
	message := "【下线通知】" + email + " 退出登录"
	_, err := conn.Write([]byte(message + "\r\n"))
	if err != nil {
		fmt.Println("退出登录失败:", err)
	}
}

func responses(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		buf := make([]byte, 5000)
		serverMsg, err := conn.Read(buf)
		// fmt.Println(serverMsg)
		if err != nil {
			fmt.Println("server response error msg", err)
			os.Exit(1)
		}
		receiveMsg := string(buf[:serverMsg])
		flag := strings.Contains(receiveMsg, "下线通知")
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
