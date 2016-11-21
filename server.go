package main

import(
    "net"
    "fmt"
    "sync"
    // "os"
    // "io"
)

var ips = make(map[string]net.Conn)

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
        instruction  := message.instruction
        switch instruction {
            case "login":
                checkOnLine(fromNickName)
                sendOnLineMsg(fromNickName)
                ips[fromNickName] = conn
                sendOnLineMsgToUser(conn)
            case "quit" :
                quit(fromNickName)
            case "msg":
                sendReponse(fromNickName, toUserName, sendMsg, conn)
            case "search":
                search(fromNickName, toUserName, conn)
        }
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

func search(fromNickName, toUserName string,  fromUserConn net.Conn) {
    searchUser := ips[toUserName] 
    message    := ""
    if searchUser != nil {
        message = "【好友搜索】"+toUserName+"在线"
    } else {
        message = "【好友搜索】"+toUserName+"不在线"
    }
    fromUserConn.Write([]byte(message))
}

func sendOnLineMsg(loginUserName string){
    for toUserName, conn := range ips {
        if toUserName != loginUserName {
            newMessage := "【上线通知】"+loginUserName + "上线" 
            conn.Write([]byte(newMessage))
        }
    }
}
func sendOnLineMsgToUser(fromUserConn net.Conn){
    newMessage := "【上线通知】您已经上线,可以通过【发送信息 msg , 退出登录 quit, 查找好友search】 格式为【关键字;昵称;消息内容】" 
    fromUserConn.Write([]byte(newMessage))
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

func sendQuitMsg(logoutUserName string){
    fmt.Println("sendQuitMsg")
    for userName, conn := range ips {
        if userName != logoutUserName {
            newMessage := "【退出通知】"+logoutUserName + "退出登录" 
            conn.Write([]byte(newMessage))
        }
    }
}

func quit(fromNickName string){
    fmt.Println("quit", fromNickName)
    conn  := ips[fromNickName]
    defer conn.Close()
    newMessage := "【下线通知】"+fromNickName+"退出登录"
    // fmt.Println("del before", ips)
    delete(ips, fromNickName)
    // fmt.Println("del after", ips)
    sendQuitMsg(fromNickName)
    conn.Write([]byte(newMessage))
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
    wg.Add(1)
    go listenResponse(listener, &wg)
    wg.Wait()
}

