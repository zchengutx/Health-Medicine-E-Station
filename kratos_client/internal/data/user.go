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

func (r *userRepo) Create(ctx context.Context, u *biz.MtUser) (*biz.MtUser, error) {
	tx := r.data.Db.Create(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}

func (r *userRepo) Find(ctx context.Context, u *biz.MtUser) (*biz.MtUser, error) {
	tx := r.data.Db.Where("mobile = ?", u.Mobile).First(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}

func (r *userRepo) Update(ctx context.Context, u *biz.MtUser) (*biz.MtUser, error) {
	tx := r.data.Db.Where("id = ?", u.Id).Updates(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}
