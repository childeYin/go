package main

import(
    "net"
    "fmt"
)

var ips map[string]net.Conn

func handleResponse(conn net.Conn, ips map[string]net.Conn) {
    fmt.Println("server handleResponse", conn)
    buf := make([]byte, 1024)
    msg, err := conn.Read(buf)
    if err != nil {
      fmt.Println("Error reading:", err)
    }
    message := parseMsg(string(buf[:msg]))
    ips[message.fromNickName] = conn
    // fmt.Println("map ips:",ips)    
    sendMsg      := message.msg
    toUserName   := message.to
    fromNickName := message.fromNickName

    sendReponse(fromNickName, toUserName, sendMsg, ips, conn)
}

func sendReponse(fromNickName string, toUserName string, sendMsg string, ips map[string]net.Conn, fromUserConn net.Conn){
    // conn := getIps(toUserName)
    // fmt.Println("conn:", conn)
    // switch conn {
        // case "":
        //     fromUserConn.Write([]byte("用户不在线"))
        //     fmt.Println(toUserName, "用户不在线")
        // case nil:
        //     fromUserConn.Write([]byte("用户不在线"))
        //     fmt.Println(toUserName, "用户不在线")
        // default :
        //     fmt.Println("conn:", conn)
        //     newMessage := "【"+fromNickName+"对你说】: "+ sendMsg  
        //     conn.Write([]byte(newMessage))
    // }
    conn  := ips[toUserName]
    if conn != nil {
      fmt.Println("conn:", conn)
      newMessage := "【"+fromNickName+"对你说】: "+ sendMsg  
      conn.Write([]byte(newMessage))
    } else {
       fromUserConn.Write([]byte("用户不在线"))
       fmt.Println("user is not exise")
    }
}

func main() {
    ips := make(map[string]net.Conn)
    service := "127.0.0.1:8080"
    listener, err := net.Listen("tcp", service)
    fmt.Println("listener:", listener)
    if err != nil {
        fmt.Println("server_error_msg_1:",err)    
    }
    for  {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("server_error_msg_2:", err)
            continue
        }
        // fmt.Println("msg:", conn)
        go handleResponse(conn, ips)
    }
    
}

