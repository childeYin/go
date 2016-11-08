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
    ips[message.from] = conn
    // fmt.Println("map ips:",ips)    
    sendMsg := message.msg
    toUser  := message.to
    sendReponse(toUser, sendMsg, ips )
}

func sendReponse(toUser string, sendMsg string, ips map[string]net.Conn){
    conn  := ips[toUser]
    if conn != nil {
      fmt.Println("conn:", conn)
      conn.Write([]byte(sendMsg))
    } else {
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
    fmt.Println("map ips:",ips)    
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

