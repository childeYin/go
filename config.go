package main

type Vertex struct {
	Lat, Long float64
}

type User struct{
	nickName string
	email string
	pwd string
}

var userInfo map[string]User

const (
	serviceAddr = "127.0.0.1:8080"
)

func init(){
	userInfo = make(map[string]User)
	userInfo["zhangjun"] = User{"尹少爷", "zhangjun", "123456"}
    userInfo["shaoye"]   = User{"小少爷", "shaoye", "123456"}
	userInfo["zhangjunshaoye"]   = User{"张君", "zhangjunshaoye", "123456"}
}
