package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
)

func initEngine(driver, url string) *xorm.Engine {
	eng, err := xorm.NewEngine(driver, url)
	if err != nil {
		return nil
	}
	return eng
}

func CreateEngine(driver, url string, maxIdleConns, maxOpenConns int) {
	if Engine != nil {
		fmt.Println("engine is not nil")
		return
	}
	Engine = initEngine(driver, url)
	Engine.SetMaxIdleConns(maxIdleConns)
	Engine.SetMaxOpenConns(maxOpenConns)

}
