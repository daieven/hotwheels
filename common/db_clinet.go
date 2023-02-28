package common

// the db init and create a db pool
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"main/config"
)

var DBengine *xorm.Engine

func InitDB() {
	var err error
	conf := config.GetConfigData()
	DBengine, err = xorm.NewEngine("mysql", conf.Db.Mysql)
	if err != nil {
		fmt.Printf("123")
	}
	DBengine.SetMaxOpenConns(100)
	DBengine.SetMaxOpenConns(999)

}
