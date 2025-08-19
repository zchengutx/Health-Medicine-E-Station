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

func (c *cityRepo) GetAddressList(ctx context.Context, userId int32) (*[]biz.MtAddress, error) {
	var addressList []biz.MtAddress
	err := c.data.Db.Where("user_id = ?", userId).Order("is_default DESC, id DESC").Find(&addressList).Error
	if err != nil {
		return nil, err
	}
	return &addressList, nil
}

func (c *cityRepo) UpdateAddress(ctx context.Context, m *biz.MtAddress) (*biz.MtAddress, error) {
	err := c.data.Db.Where("id = ?", m.Id).Updates(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cityRepo) DeleteAddress(ctx context.Context, id int32) error {
	err := c.data.Db.Where("id = ?", id).Delete(&biz.MtAddress{}).Error
	return err
}

func (c *cityRepo) SetDefaultAddress(ctx context.Context, userId int32, addressId int32) error {
	// 先取消该用户所有地址的默认状态
	err := c.data.Db.Model(&biz.MtAddress{}).Where("user_id = ?", userId).Update("is_default", false).Error
	if err != nil {
		return err
	}
	
	// 如果addressId > 0，设置指定地址为默认地址
	if addressId > 0 {
		err = c.data.Db.Model(&biz.MtAddress{}).Where("id = ? AND user_id = ?", addressId, userId).Update("is_default", true).Error
	}
	
	return err
}
