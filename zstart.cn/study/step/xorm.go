// xorm
package step

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // 乐观锁
}

var engine *xorm.Engine

func StudyXORM() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:hazel110@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	engine.Ping()
}
