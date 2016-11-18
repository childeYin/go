package main

import(
    "net"
    "fmt"
    "sync"
    // "os"
    // "io"
)

var ips       = make(map[string]net.Conn)
var onlineIps = make(map[string]net.Conn)

func handleResponse(conn net.Conn) {
    defer conn.Close()
    for {
        fmt.Println("handleResponse",ips)
        fmt.Println("server handleResponse", conn)
        buf := make([]byte, 5000)
        msg, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error reading:", err)
            return 
        }
        message := parseMsg(string(buf[:msg]))
        fmt.Println(message)
        sendMsg      := message.msg
        toUserName   := message.to
        fromNickName := message.fromNickName
        if fromNickName == toUserName {
            // fmt.Println("自己登录")
            checkOnLine(message.fromNickName)
            onlineIps[message.fromNickName] = conn
        }

        ips[message.fromNickName] = conn
        fmt.Println("handleResponse",ips)

      

        
        sendReponse(fromNickName, toUserName, sendMsg, conn)
    }
}

func sendReponse(fromNickName string, toUserName string, sendMsg string, fromUserConn net.Conn){
    conn  := ips[toUserName]
    fmt.Println("to User conn", conn, toUserName)
    fmt.Println(sendMsg)
    if conn != nil {
        fmt.Println("conn:", conn)
        newMessage := "【"+fromNickName+"对你说】: "+ sendMsg+"\r\n" 
        fmt.Println(newMessage)
        flag, err  := conn.Write([]byte(newMessage))
        fmt.Println(flag)
        if err != nil {
            fmt.Println("error ", err)
            delete(ips, toUserName)
            fromUserConn.Write([]byte("用户不在线\r\n"))
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
        _, err  := conn.Write([]byte(newMessage))
        if err != nil {
            fmt.Println("error ", err)
            continue
        }
    }
}

func checkNet(wg *sync.WaitGroup){
    defer wg.Done()
    for {
        fmt.Println("checkNet", ips)
        for index, conn := range ips {
            fmt.Println("checkNet conn",  index, conn)
        }
    }
}

func sendOnLine(wg *sync.WaitGroup){
    defer wg.Done()
    for {
        // fmt.Println("sentOnLine onlineIps", onlineIps)
        for name, _ := range onlineIps {
            for toUserName, conn := range ips {
                if toUserName != name {
                    newMessage := "【上线通知】"+name + "上线" 
                    conn.Write([]byte(newMessage))
                }
            }
            delete(onlineIps, name)
        }
    }
}

func checkOnLine(fromNickName string){
    userConn := ips[fromNickName]
    if ips[fromNickName] != nil {
        fmt.Println(fromNickName, "已经登录")
        newMessage := "【下线通知】"+fromNickName+",已经在其他端登录"
        fmt.Println(newMessage)
        userConn.Write([]byte(newMessage))
    }
}

func listenResponse(listener net.Listener, wg *sync.WaitGroup){
    defer wg.Done()
    for  {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("server_error_msg_2:", err)
            continue
        }
        go handleResponse(conn)
    }
}

func main() {
    fmt.Println("listener server")

    listener, err := net.Listen("tcp", serviceAddr)
    fmt.Println("listener:", listener)
    if err != nil {
        fmt.Println("server_error_msg_1:",err)    
    }
    var wg sync.WaitGroup
    wg.Add(2)
    go listenResponse(listener, &wg)
    // go checkNet(&wg)
    go sendOnLine(&wg)
    wg.Wait()
}

