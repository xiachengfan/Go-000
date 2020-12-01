package model

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID     uint32
	Name   string
	Secret string
}

var db *gorm.DB

func Init() (err error) {
	db, err = gorm.Open(mysql.Open("root:rootroot@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	return errors.Wrap(err, "Connect to mysql error")
}

func DB() *gorm.DB {
	return db
}
