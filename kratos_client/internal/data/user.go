package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_client/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	tx := r.data.Db.Create(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return nil, nil
}
