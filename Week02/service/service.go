package service

import (
	"Go-000/Week02/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}
