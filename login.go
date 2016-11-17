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
     "runtime"
    "sync"
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
    fmt.Println(nickName)
    if nickName != "" {
        conn, err := net.Dial("tcp", serviceAddr)
        defer conn.Close()
        var wg sync.WaitGroup
        wg.Add(2)

        if err != nil {
            log.Println("error msg continue! ", err)
        }
        // fmt.Println(conn)
        autoRequest(conn, nickName)
        go requests(conn, &wg, nickName)
        go responses(conn, &wg)
        wg.Wait()
    } else {
        log.Println("程序终止!")
    }
}

func handleLogin(email, pwd string) string{
    var user, ok = userInfo[email]
    if ok == true {
        if strings.EqualFold(pwd, user.pwd) == true {
            fmt.Println(user.nickName)
            return user.nickName
        }
    }
    return ""
}


func handleMessage(nickName, message string) string{
    if message == "auto" {
        message = nickName+";"+"登录成功"
    } else {
        fmt.Println("请输入信息: ")
        reader     := bufio.NewReader(os.Stdin)
        data, _, _ := reader.ReadLine()
        message     = string(data)
    }
    fmt.Println("输入信息为: ", message)
    return message
}

func readConfig(){
    // file, e := ioutil.ReadFile("./config.json")
    // fmt.Printf("%s\n", string(file))
    // fmt.Println("error:", e)
    // var jsontype 
    // json.Unmarshal([]byte(string(file)), &jsontype)
    // fmt.Println("jsontype:", jsontype)
}
func main() {
    runtime.GOMAXPROCS(2)
    login()
     // go requests(conn, nickName)
            // go responses(conn)
    // readConfig()
}
