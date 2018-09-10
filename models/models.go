package models

type User struct {
	Id   int    `xorm:"id"`
	Name string `xorm:"name"`
	Age  int    `xorm:"age"`
}
