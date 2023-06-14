package common

import (
	"gorm.io/gorm"
)

func InitTable(dbEngine *gorm.DB, tables ...interface{}) {
	for _, table := range tables {
		dbEngine.AutoMigrate(table)
		// TODO: has the error happened, don't use it now
	}
}
