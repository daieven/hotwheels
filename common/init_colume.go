package common

import (
	"fmt"
	"gorm.io/gorm"
)

func InitTable(dbEngine *gorm.DB, tables ...interface{}) {
	for _, table := range tables {
		err := dbEngine.AutoMigrate(table)
		if err != nil {
			fmt.Print(table)
			panic(err)
		}
	}
}
