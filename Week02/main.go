package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

const (
	NOT_FOUND string = "Not Found"
)


func main() {
	Bff()
}




func Bff() {
	id := uint64(110)
	userName, err := GetNameById(id)
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
	fmt.Printf("db select user name:%s", userName)

}


func GetNameById(id uint64) (UserName string, err error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8")
	if err != nil {
		return "default", errors.New("数据库连接失败")
	}
	defer db.Close()
	sqlstr := "SELECT username FROM users WHERE uid=?"
	err = db.QueryRow(sqlstr, id).Scan(&UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "default", errors.New(NOT_FOUND)
		}
		return "default",errors.Wrap(err, "get db error")

	}
	return UserName, nil
}