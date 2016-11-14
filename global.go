package main

import (
	"fmt"
	"net"
	// "gopkg.in/redis.v5"
	"bytes"
	"encoding/json"
	"encoding/gob"
	"github.com/garyburd/redigo/redis"
)

type Msg struct {
	conn net.Conn
}
var ipsNew map[string]interface{}

func setIps(nickName string, conn net.Conn){
	ipsNew := make(map[string]net.Conn)
	client, errConn := redis.Dial("tcp", "localhost:6379")
    if errConn != nil {
        fmt.Println("Connect to redis error", errConn)
        return
    }
    connRedis := &Conn{conn}
    // ipsNew[nickName] = conn
    fmt.Println("Map connRedis: ", connRedis)
    // serialized, err := json.Marshal(ipsNew)
	// fmt.Println("set---getConndeserialized:", string(serialized))

    // var deserialized interface{}
    // json.Unmarshal(serialized, &deserialized)
	// fmt.Println("set---getConndeserialized:", deserialized)

    err = client.Cmd("HMSET", "ips", nickName, connRedis)
    // _, err = client.Do("HMSET", "ips", nickName, serialized)
    if err != nil {
	    fmt.Println("redis set failed:", err)
	}
}

func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    fmt.Println("redis get failed Buf:", key)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}


// func getIps(nickName string) interface{}{
func getIps(nickName string){
	client, errLink := redis.Dial("tcp", "localhost:6379")
    if errLink != nil {
        fmt.Println("Connect to redis error", errLink)
    }
    serialized, errGet := redis.Values(client.Do("HMGET", "ips", nickName))
    if errGet != nil {
        fmt.Println("redis get error", errGet)
    }
	newSerialized, errByte := GetBytes(serialized)
	if errByte != nil {
        fmt.Println("redis getBytes error", errByte)
	}
    fmt.Println("new Serialized redis get failed:", newSerialized)

    var deserialized Conn
    errDeJson := json.Unmarshal(newSerialized, &deserialized)
    if errDeJson != nil {
    	fmt.Println("redis get failed:", errDeJson)
    }
	fmt.Println("getConndeserialized:", deserialized)


 //    // err = json.Unmarshal(serialized, &deserialized)

	// fmt.Println("getConndeserialized:", deserialized)

 //    if err != nil {
	//     fmt.Println("redis get failed:", err)
	//     return deserialized
	// } else {
	// 	return deserialized
	// }
}

func setMsg(msg string) {
	s := strings.Split(msg, ";")
	length := len(s)
	if length < 2 {
		return ""
	}
	nickName := s[0]
	message  := s[1]

	client, errLink := redis.Dial("tcp", "localhost:6379")
    if errLink != nil {
        fmt.Println("Connect to redis error", errLink)
    }
    _, errSet := client.Do("LPUSH", "ips", nickName, serialized)


}



