package main

import(
    "net"
    "fmt"
    "sync"
)

var ips map[string]net.Conn

// func handleResponse(conn net.Conn, wg *sync.WaitGroup, ips map[string]net.Conn) {
func handleResponse(conn net.Conn, ips map[string]net.Conn) {
    defer conn.Close()
    for {
        fmt.Println("server handleResponse", conn)
        buf := make([]byte, 5000)
        msg, err := conn.Read(buf)
        fmt.Println("conn read msg", msg)
        if err != nil {
          fmt.Println("Error reading:", err)
          continue
        }
        message := parseMsg(string(buf[:msg]))
        fmt.Println(message)
        ips[message.fromNickName] = conn
        fmt.Println("ips :", ips)
        sendMsg      := message.msg
        toUserName   := message.to
        fromNickName := message.fromNickName

        sendReponse(fromNickName, toUserName, sendMsg, ips, conn)
    }
}

func sendReponse(fromNickName string, toUserName string, sendMsg string, ips map[string]net.Conn, fromUserConn net.Conn){
    conn  := ips[toUserName]
    fmt.Println("to User conn", conn, toUserName)
    fmt.Println(sendMsg)
    if conn != nil {
        fmt.Println("conn:", conn)
        newMessage := "【"+fromNickName+"对你说】: "+ sendMsg+"\r\n" 
        fmt.Println(newMessage)
        _, err  := conn.Write([]byte(newMessage))
        if err != nil {
            fmt.Println("error ", err)
        }
    } else {
        fromUserConn.Write([]byte("用户不在线\r\n"))
        fmt.Println("user is not exise")
    }
}

func handleResponseToSelf(conn net.Conn, wg *sync.WaitGroup){
    defer conn.Close()
    
    for {
        fmt.Println("handleResponseToSelf")
        newMessage := "消息收到了"
        fmt.Println(newMessage)
        flag, err  := conn.Write([]byte(newMessage))
        fmt.Println(flag)
        if err != nil {
            fmt.Println("error ", err)
            continue
        }
    }
}


func main() {
    fmt.Println("listener server")

    ips := make(map[string]net.Conn)
    listener, err := net.Listen("tcp", serviceAddr)
    fmt.Println("listener:", listener)
    if err != nil {
        fmt.Println("server_error_msg_1:",err)    
    }
    for  {
        conn, err := listener.Accept()
        // defer conn.Close()
        if err != nil {
            fmt.Println("server_error_msg_2:", err)
            continue
        }
        fmt.Println(ips)
        go handleResponse(conn, ips)
    }
    
}

