package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	ID     uint32
	Name   string
	Secret string
}

func (u *User) Get(db *gorm.DB) (*User, error) {
	var user User
	db = db.Where("name= ? AND secret= ? and is_del= ?", u.Name, u.Secret, 0)
	err := db.Find(&user).Error
	//为了防止调用者用用errors.Is做unWrap，需要把Sentinel errors的错误在调用层以普通的错误返回。
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "dao: user not found")
	} else {
		return nil, err
	}
	return &user, nil
}
