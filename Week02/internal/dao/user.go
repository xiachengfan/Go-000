package dao

import (
	"Go-000/Week02"
	"Go-000/Week02/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var db *gorm.DB

type Dao struct{}

func (d Dao) Get(name, secret string) (*model.User, error) {
	var user *model.User
	db = db.Where("name= ? AND secret= ? ", name, secret)
	err := db.Find(user).Error
	//为了防止调用者用用errors.Is做unWrap，需要把Sentinel errors的错误在调用层以普通的错误返回。
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(Week02.ErrUserNotFound, " %v user info not exist ", name)
	} else {
		return nil, errors.Wrapf(err, "find user error %v", name)
	}
	return user, nil
}
