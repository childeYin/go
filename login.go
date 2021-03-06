package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
)

var conn net.Conn

func login() string {
	log.Println("请输入邮箱: ")
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	email := string(data)
	log.Println("请输入密码: ")
	reader = bufio.NewReader(os.Stdin)
	data, _, _ = reader.ReadLine()
	password := string(data)
	fmt.Println("email=" + email + "; pwd=" + password)
	nickName := handleLogin(email, password)
	fmt.Println(nickName)
	if nickName != "" {
		conn, _ = net.Dial("tcp", serviceAddr)
		if conn == nil {
			fmt.Println("服务器繁忙")
			return ""
		}
		return email
	} else {
		log.Println("程序终止!")
		return ""
	}
}

func handleLogin(email, pwd string) string {
	var user, ok = userInfo[email]
	fmt.Println(user.nickName)
	fmt.Println(user.pwd)
	if ok == true {
		if strings.EqualFold(pwd, user.pwd) == true {
			fmt.Println(user.nickName)
			return user.nickName
		}
	}
	return ""
}

func handleRequest(email string, wg *sync.WaitGroup) {
	defer conn.Close()
	autoLoginRequest(conn, email)
	requests(conn, wg, email)
}

func handleResponse(wg *sync.WaitGroup) {
	defer conn.Close()
	responses(conn, wg)
}

func handleMessage(nickName, message string) string {
	if message == "auto" {
		message = nickName + ";" + "登录成功"
	} else {
		fmt.Println("请输入信息: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		message = string(data)
	}
	fmt.Println("输入信息为: ", message)
	return message
}

func readConfig() {
	// file, e := ioutil.ReadFile("./config.json")
	// fmt.Printf("%s\n", string(file))
	// fmt.Println("error:", e)
	// var jsontype
	// json.Unmarshal([]byte(string(file)), &jsontype)
	// fmt.Println("jsontype:", jsontype)
}

func main() {
	runtime.GOMAXPROCS(2)
	email := login()
	var wg sync.WaitGroup
	wg.Add(2)
	if email != "" {
		go handleRequest(email, &wg)
		go handleResponse(&wg)
	} else {
		// die()
		os.Exit(1)
	}
	wg.Wait()
}
