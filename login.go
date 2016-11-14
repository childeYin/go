package main

import(
    "bufio"
    // "encoding/json"
    "log"
    "os"
    "net"
    "fmt"
    // "io/ioutil"
    // "reflect"
    "strings"
)

func login(){
    log.Println("请输入邮箱: ")
    reader     := bufio.NewReader(os.Stdin)
    data, _, _ := reader.ReadLine()
    email      := string(data)
    log.Println("请输入密码: ")
    reader      = bufio.NewReader(os.Stdin)
    data, _, _  = reader.ReadLine()
    password   := string(data)
    // nickName   := getUserByAccount(email, password)
    nickName := handleLogin(email, password)
    if nickName != "" {
        log.Println("登录成功", nickName, "欢迎回来!")
        handle("auto", nickName)
        for {
            handle("", nickName)
            continue
        }
    } else {
        log.Println("程序终止!")
    }
}

func handleLogin(email, pwd string) string{
    var user, ok = userInfo[email]
    if ok == true {
        fmt.Println(user.pwd)
        fmt.Println(pwd)
        if strings.EqualFold(pwd, user.pwd) == true {
            fmt.Println(user.nickName)
            return user.nickName
        }
    }
    return ""
}
func readConfig(){
    // file, e := ioutil.ReadFile("./config.json")
    // fmt.Printf("%s\n", string(file))
    // fmt.Println("error:", e)
    // var jsontype 
    // json.Unmarshal([]byte(string(file)), &jsontype)
    // fmt.Println("jsontype:", jsontype)
}

func handle(message string, nickName string){
    descAddr := "127.0.0.1:8080"
    conn, err := net.Dial("tcp", descAddr)
    // setIps(nickName, conn)
    // getIps(nickName)
    if err != nil {
        log.Println("error msg continue!")
    }
    if message == "auto" {
        message = nickName+";"+"登录成功"
        go request(message, conn, nickName)
        go response(conn)
    } else {
        fmt.Println("请输入信息: ")
        reader     := bufio.NewReader(os.Stdin)
        data, _, _ := reader.ReadLine()
        message     = string(data)
        // log.Println("输入信息为: ", message)
        go request(message, conn, nickName)
        go response(conn)
    }
}

func main() {
    login()
    // readConfig()
}
