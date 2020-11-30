package service

import (
	"Go-000/Week02/model"
	"github.com/pkg/errors"
)

type GetAuthRequest struct {
	Name   string `json:"name" binding:"required"`
	Secret string `json:"secret" binding:"required,min=6,max=6"`
}

func (svc *Service) CheckAuth(param *GetAuthRequest) (*model.User, error) {
	user, err := svc.dao.GetUser(
		param.Name,
		param.Secret,
	)
	if err != nil {
		return nil, err
	}
	if user.ID > 0 {
		return user, nil
	}
	// 应用新生成的errors
	return nil, errors.New("auth info does not exist or secret is false.")
}
