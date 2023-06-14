package common

// the db init and create a db pool
import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"main/config"
)

var DBengine *gorm.DB

func InitDB() *gorm.DB {
	var err error
	conf := config.GetConfigData()

	DBengine, err := gorm.Open(mysql.Open(conf.Db.Mysql), &gorm.Config{})
	//DBengine, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DBengine
}
