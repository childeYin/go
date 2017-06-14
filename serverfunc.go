package main

import(
    "net"
    "fmt"
    "sync"
    "os"
    "io"
    "strings"
    "bufio"
)
const addFriendFileName = "./file/addfriend.log"
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
        sendMsg     := message.msg
        toUserEmail := message.to
        email       := message.email
        instruction := message.instruction
        fmt.Println("sendMsg=="+sendMsg)
        switch instruction {
            case "login":
                checkOnLine(email)
                sendOnLineMsg(email)
                ips[email] = conn
                sendOnLineMsgToUser(conn)
            case "quit" :
                quit(email)
            case "msg":
                sendResponse(email, toUserEmail, sendMsg, conn)
            case "all":
                sendAllResponse(email, sendMsg, conn)
            case "search":
                search(toUserEmail, conn, "search")
            case "add" :
            	add(email, toUserEmail , conn)
        }
    }
}
func sendAllResponse(email string, sendMsg string, fromUserConn net.Conn){
    fmt.Println("fromEmail:", email)
    fromNickName := getFromUserNickName(email)
    fromConn     := ips[email]
    for userEmail, conn := range ips {
        fmt.Printf("userEmail:%s", userEmail)
        if conn != fromConn {
            fmt.Println("conn:", conn)
            newMessage := "【"+fromNickName+"说】: "+ sendMsg+"\r\n"
            fmt.Println(newMessage)
            flag, err  := conn.Write([]byte(newMessage))
            fmt.Println(flag)
            fmt.Println(err)
        }
    }
}
func sendResponse(email string, toUserName string, sendMsg string, fromUserConn net.Conn){
    conn  := ips[toUserName]
    fmt.Println("to User conn", conn, toUserName)
    fmt.Println(sendMsg)
    fromNickName := getFromUserNickName(email)
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

func getFromUserNickName(email string) string{
    var user, ok = userInfo[email]
    fmt.Println("getFromUserNickName", )
    if ok == true {
        return user.nickName
    } else {
       return "匿名"
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

func search(toUserName string, fromUserConn net.Conn, flag string) string{
    message    := ""
    var _, ok = userInfo[toUserName]
    fmt.Println("search", userInfo)
    if flag == "search" {
		if ok == true {
	        message = "【好友搜索】找到用户"+toUserName
	    } else {
	        message = "【好友搜索】"+toUserName+"用户不存在"
	    }
	    fromUserConn.Write([]byte(message))
	    return ""
    } else {
		if ok == true {
	        return "exise"
	    } else {
	    	return ""
	    }
    }

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
    newMessage := "【上线通知】您已经上线,可以通过【发送信息 msg , 退出登录 quit, 查找好友search, 全部发送all】 格式为【关键字;昵称;消息内容】"
    fromUserConn.Write([]byte(newMessage))
}

func checkOnLine(email string){
    userConn     := ips[email]
    fromNickName := getFromUserNickName(email)
    if userConn  != nil {
        fmt.Println(fromNickName, "已经登录")
        newMessage := "【下线通知】"+fromNickName+",已经在其他端登录"
        fmt.Println(newMessage)
        userConn.Write([]byte(newMessage))
    }
}

func sendQuitMsg(logoutUserEmail, fromNickName string){
    fmt.Println("sendQuitMsg")
    for email, conn := range ips {
        if email != logoutUserEmail {
            newMessage := "【退出通知】"+ fromNickName + "退出登录"
            conn.Write([]byte(newMessage))
        }
    }
}

func quit(email string){
    fmt.Println("quit", email)
    conn  := ips[email]
    if conn != nil {
    	fromNickName := getFromUserNickName(email)
	    defer conn.Close()
	    newMessage := "【下线通知】"+fromNickName+"退出登录"
	    // fmt.Println("del before", ips)
	    delete(ips, fromNickName)
	    // fmt.Println("del after", ips)
	    sendQuitMsg(email, fromNickName)
	    conn.Write([]byte(newMessage))
    }
}
func addFriendRequests(email, addUserEmail string, fromUserConn, addUserConn net.Conn){

	addUserNickName  := getFromUserNickName(addUserEmail)
	fromUserNickName := getFromUserNickName(email)

	flag := fileExistsAndWrite(addFriendFileName, "")

	if flag {
		checkFlag := checkRequest(email, addUserEmail)
        fmt.Println("addFriendRequests", checkFlag)
        if (!checkFlag) {
            message := "【添加好友请求】添加"+addUserNickName+"用户为好友的请求已经发送"
            fromUserConn.Write([]byte(message))
            message = "【添加好友请求】"+fromUserNickName+"用户,请求添加您为好友,同意请回复【 add;"+email+"】"
            addUserConn.Write([]byte(message))
            return
        }
        message := "【添加好友请求】添加"+addUserNickName+"用户为好友的请求已经发送过，请勿重复发送"
        fromUserConn.Write([]byte(message))
		return
	}
	message := "【添加好友】添加好友失败,稍后重试"
	fromUserConn.Write([]byte(message))
	return
}

func checkRequest(email, addUserEmail string) bool{
	reqMsg  := email+","+addUserEmail+"\r\n"
	flag    := checkFileContent(addFriendFileName, reqMsg)
	fmt.Println("checkRequest", flag)
	if (!flag) {
		fileExistsAndWrite(addFriendFileName, reqMsg)
	}
    return flag
}

func generateFriendFile(email, addUserEmail string) bool{
	fromFileName := "./file/"+email
	fromMsg 	 := addUserEmail+"\r\n"
	fromFlag 	 := fileExistsAndWrite(fromFileName, fromMsg)

	addFileName := "./file/"+addUserEmail
	addMsg := email+"\r\n"
	addFlag 	:= fileExistsAndWrite(addFileName, addMsg)

	if addFlag && fromFlag {
		return true
	} else {
		return false
	}
}


func checkResponse(email, addUserEmail string, fromUserConn, addUserConn net.Conn) bool {
	fmt.Println("checkResponse")
	respMsg := addUserEmail+","+email
	flag    := checkFileContent(addFriendFileName, respMsg)
	fmt.Println("checkResponse flag",flag)
	addUserNickName := getFromUserNickName(addUserEmail)
	fromUserNickName := getFromUserNickName(email)
	if flag {
		fileFlag := generateFriendFile(email, addUserEmail)
		if fileFlag {
			message := "【添加好友】" + fromUserNickName+"用户同意添加您为好友"
			addUserConn.Write([]byte(message))
			message = "【添加好友】同意了来自于"+addUserNickName+"用户的好友请求"
			fromUserConn.Write([]byte(message))
		} else {
			message := "【添加好友】同意了来自于"+addUserNickName+"用户的好友请求,失败,稍后重试"
			fromUserConn.Write([]byte(message))
		}
	}
	return flag
}

func checkFileContent(fileName, msg string) bool{
	inputFile, inputError := os.Open(fileName)
    if inputError != nil {
		fmt.Println("checkFileContent inputError", inputError)
        return false
    }
    defer inputFile.Close()
    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        inputString = strings.TrimSpace(inputString)
        msg = strings.TrimSpace(msg)
        if readerError == io.EOF {
			fmt.Println("checkFileContent end", )
            break
        }
        fmt.Println(inputString," == ", msg)
        if inputString == msg {
        	return true;
        }
    }
    return false
}

func fileExistsAndWrite(fileName, msg string) bool{
	_, err := os.Stat(fileName)
	if err != nil {
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println("create file ", fileName, " fail")
			return false
		}
		if msg != "" {
			f.WriteString(msg)
		}
		return true
	}
	f, _:= os.OpenFile(fileName, os.O_WRONLY, os.ModeAppend)
	if msg != "" {
		_, err = f.WriteString(msg)
	}
	return true
}
func checkFriendFile(email, addUserEmail string) bool{
	fileName := "./file/"+email
	flag := checkFileContent(fileName, addUserEmail)
	return flag
}
func add(email string, addUserEmail string , fromUserConn net.Conn) {
	addUserConn := ips[addUserEmail]
	addUserNickName := getFromUserNickName(addUserEmail)

	friendFlag := checkFriendFile(email, addUserEmail)
	if friendFlag {
		message := "【添加好友】"+addUserNickName+"用户,已经是您的好友,请勿重复添加"
		fromUserConn.Write([]byte(message))
		return
	}
	flag   := search(addUserEmail, fromUserConn, "add")
	if flag == "exise" {
		if addUserConn != nil {
			flag := checkResponse(email, addUserEmail, fromUserConn, addUserConn)
			if !flag {
				addFriendRequests(email, addUserEmail, fromUserConn, addUserConn)
				return
			}
			return
		}
		message := "【添加好友】添加"+addUserNickName+"用户为好友,该用户不在线,请稍后发送"
		fromUserConn.Write([]byte(message))
		return
	}
	message := "【添加好友】未找到"+addUserNickName+"用户"
	fromUserConn.Write([]byte(message))
	return
}
