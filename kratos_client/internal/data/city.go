package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_client/internal/biz"
)

type cityRepo struct {
	data *Data
	log  *log.Helper
}

func NewCityRepo(data *Data, logger log.Logger) biz.CityRepo {
	return &cityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cityRepo) Find(ctx context.Context, m *biz.MtCity) (*[]biz.MtCity, error) {
	var City []biz.MtCity
	err := c.data.Db.Where("name like ?", "%市%").Find(&City).Error
	if err != nil {
		return nil, err
	}
	return &City, nil
}

func (c *cityRepo) LikeFind(ctx context.Context, m *biz.MtCity) (*[]biz.MtCity, error) {
	var mList []biz.MtCity
	err := c.data.Db.Where("name like ?", "%"+m.Name+"%").Find(&mList).Error
	if err != nil {
		return nil, err
	}
	return &mList, nil
}

func (c *cityRepo) Create(ctx context.Context, m *biz.MtAddress) (*biz.MtAddress, error) {
	err := c.data.Db.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}
