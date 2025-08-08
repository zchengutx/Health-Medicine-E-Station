package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos_client/internal/biz"
)

type drugRepo struct {
	data *Data
	log  *log.Helper
}

func NewDrugRepo(data *Data, logger log.Logger) biz.DrugRepo {
	return &drugRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *drugRepo) ListDrug(ctx context.Context, fristCategoryId int32, secondCategoryId int32, keyword string) ([]*biz.MtDrug, error) {
	var drugs []*biz.MtDrug
	var err error
	db := r.data.Db.Model(&biz.MtDrug{})
	if keyword != "" {
		db = db.Where("drug_name like ?", "%"+keyword+"%")
	}
	if secondCategoryId > 0 {
		db = db.Where("second_category_id = ?", secondCategoryId)
	}
	if fristCategoryId > 0 {
		db = db.Where("frist_category_id = ?", fristCategoryId)
	}
	err = db.Find(&drugs).Error
	if err != nil {
		return nil, err
	}
	return drugs, nil
}

func (r *drugRepo) GetDrug(ctx context.Context, id int32) (*biz.MtDrug, error) {
	var drug biz.MtDrug
	err := r.data.Db.Where("id = ?", id).First(&drug).Error
	if err != nil {
		return nil, err
	}
	return &drug, nil
}

func (r *drugRepo) GetExplain(ctx context.Context, id int32) (*biz.MtExplain, error) {
	var explain biz.MtExplain
	err := r.data.Db.Where("id = ?", id).First(&explain).Error
	if err != nil {
		return nil, err
	}
	return &explain, nil
}

func (r *drugRepo) GetGuide(ctx context.Context, id int32) (*biz.MtGuide, error) {
	var guide biz.MtGuide
	err := r.data.Db.Where("id = ?", id).Find(&guide).Limit(1).Error
	if err != nil {
		return nil, err
	}
	return &guide, nil
}