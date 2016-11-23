package main

import(
    "net"
    "fmt"
    "sync"
    // "os"
    // "io"
)

var ips = make(map[string]net.Conn)

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

