package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

var db = &sql.DB{}

func init(){
    db,_ = sql.Open("mysql", "root:@/im?charset=utf8")
}

func getUserByAccount(email, password string) string{
	fmt.Println("email:", email)
	fmt.Println("password:", password)
	var login_name string
	err := db.QueryRow("select login_name from user where email=? and password=?", email, password).Scan(&login_name)
	if err != nil {
        fmt.Println("error_message:", err)
        return ""
    }
    return login_name
}