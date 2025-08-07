package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_client/internal/biz"
)

type estimateRepo struct {
	data *Data
	log  *log.Helper
}

func NewEstimateRepo(data *Data, logger log.Logger) biz.EstimateRepo {
	return &estimateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *estimateRepo) ListDrugEstimate(ctx context.Context, drugId int32, userId int32) ([]*biz.MtEstimate, *biz.MtDrug, *biz.MtUser, error) {
	var estimates []*biz.MtEstimate
	var err error
	var drug *biz.MtDrug
	err = r.data.Db.Where("id = ?", drugId).Find(&drug).Limit(1).Error
	if err != nil {
		return nil, nil, nil, err
	}
	err = r.data.Db.Where("drug_id = ?", drugId).Find(&estimates).Error
	if err != nil {
		return nil, nil, nil, err
	}
	var user *biz.MtUser
	err = r.data.Db.Where("id = ?", userId).Find(&user).Limit(1).Error
	if err != nil {
		return nil, nil, nil, err
	}
	return estimates, drug, user, nil
}
func (r *estimateRepo) GetEstimate(ctx context.Context, id int32, userId int32) (*biz.MtEstimate, *biz.MtUser, error) {
	var estimate *biz.MtEstimate
	var user *biz.MtUser
	err := r.data.Db.Where("id = ?", id).Find(&estimate).Limit(1).Error
	err = r.data.Db.Where("user_id = ?", userId).Find(&user).Limit(1).Error
	if err != nil {
		return nil, user, err
	}
	return estimate, user, nil
}
func (r *estimateRepo) CreateEstimate(ctx context.Context, estimate *biz.MtEstimate) error {
	err := r.data.Db.Create(estimate).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *estimateRepo) DeleteEstimate(ctx context.Context, id int32) error {
	err := r.data.Db.Where("id = ?", id).Delete(&biz.MtEstimate{}).Error
	if err != nil {
		return err
	}
	return nil
}
