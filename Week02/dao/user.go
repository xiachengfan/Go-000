package dao

import "Go-000/Week02/model"

func (d *Dao) GetUser(name, secret string) (*model.User, error) {
	user := model.User{Name: name, Secret: secret}
	return user.Get(d.engine)
}
