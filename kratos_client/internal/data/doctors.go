package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_client/internal/biz"
)

type doctorsRepo struct {
	data *Data
	log  *log.Helper
}

func NewDoctorsRepo(data *Data, logger log.Logger) biz.DoctorsRepo {
	return &doctorsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (c *doctorsRepo) DoctorsFind(ctx context.Context, m *biz.MtDoctors) (*[]biz.MtDoctors, error) {
	var Doctors []biz.MtDoctors
	err := c.data.Db.Find(&Doctors).Error
	if err != nil {
		return nil, err
	}
	return &Doctors, nil
}

func (c *doctorsRepo) FindByID(ctx context.Context, id int32) (*biz.MtDoctors, error) {
	var doctor biz.MtDoctors
	err := c.data.Db.Where("id = ?", id).First(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}
