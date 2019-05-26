package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DSN      string
	Idle     int
	Active   int
	IdleTime int
}

var (
	db *gorm.DB
)

func NewMysql(c *Config) {
	var err error
	db, err = gorm.Open("mysql", c.DSN)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.SingularTable(true)

	return
}
