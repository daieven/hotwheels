package models

type User struct {
	Id       int64
	Username string `xorm:"varchar(25) notnull unique 'username'"`
	Password string `xorm:"varchar(25) notnull 'password'"`
}
