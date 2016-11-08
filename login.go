package main

import(
    "bufio"
    "log"
    "os"
    "net"
    "fmt"
)
type User struct {
    addr  string
    nickName string
}

func login() {
    log.Println("请输入邮箱: ")
    reader     := bufio.NewReader(os.Stdin)
    data, _, _ := reader.ReadLine()
    email      := string(data)
    log.Println("请输入密码: ")
    reader      = bufio.NewReader(os.Stdin)
    data, _, _  = reader.ReadLine()
    password   := string(data)
    nickName   := getUserByAccount(email, password)
    if nickName != "" {
        log.Println("登录成功", nickName, "欢迎回来!")
        tcpAddr   := "127.0.0.1:8080"
        for {
            conn, err := net.Dial("tcp", tcpAddr)
            if err != nil {
                log.Println("error msg continue!")
            }
            fmt.Print("请输入信息: ")
            reader     := bufio.NewReader(os.Stdin)
            data, _, _ := reader.ReadLine()
            message    := string(data)
            // log.Println("输入信息为: ", message)
            request(message, conn, email)
            response(conn)
        }
    }
    log.Println("程序终止!")

}

// func getIpArrress(nickName string) string{
//     addrs, err := net.InterfaceAddrs()
//     // log.Println("addrs:",addrs)
//     // log.Println("err:",err)
//     var ip string
//     for _, address := range addrs {
//         // 检查ip地址判断是否回环地址
//         if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//             if ipnet.IP.To4() != nil {
//                 log.Println(ipnet.IP.String())
//                 ip   := ipnet.IP.String()
//                 return ip
//             }
//         }
//     }
//     return ip
// }
func main() {
    login()
}
