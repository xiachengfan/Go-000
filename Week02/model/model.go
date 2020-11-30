package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine() (*gorm.DB, error) {
	db, err := gorm.Open(fmt.Sprintf("root:root@tcp(172.0.0.1:3306)/blog"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
