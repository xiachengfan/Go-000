package service

import (
	"Go-000/Week02/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

type Resp struct {
	CODE int
	MSG  interface{}
	ERR  error
}
