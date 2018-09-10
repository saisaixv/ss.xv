package main

import (
	"fmt"

	"ss.xv/models"
	"ss.xv/utils/db"
)

func main() {
	fmt.Println("hello")
	user := new(models.User)

	db.CreateEngine("mysql", "root:saisai@@@/db_user?charset=utf8", 10, 10)

	has, err := db.Engine.Get(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if has {
		fmt.Println("存在 =", user)
	} else {
		fmt.Println("不存在")
	}
}
