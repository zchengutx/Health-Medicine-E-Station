package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type MtEstimate struct {
	Id        int32     `gorm:"column:id;type:int;primaryKey;" json:"id"`
	DrugId    int32     `gorm:"column:drug_id;type:int;comment:药品id;not null;default:0;" json:"drug_id"`      // 药品id
	Content   string    `gorm:"column:content;type:varchar(30);comment:评价内容;" json:"content"`                 // 评价内容
	Image     string    `gorm:"column:image;type:varchar(200);comment:评价图片;not null;default:0;" json:"image"` // 评价图片
	UserId    int32     `gorm:"column:user_id;type:int;not null;default:0;" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);" json:"created_at"`
}

func (m *MtEstimate) TableName() string {
	return "mt_estimate"
}

type EstimateRepo interface {
	CreateEstimate(ctx context.Context, estimate *MtEstimate) error
	DeleteEstimate(ctx context.Context, id int32) error
	GetEstimate(ctx context.Context, id int32, userId int32) (*MtEstimate, *MtUser, error)
	ListDrugEstimate(ctx context.Context, id int32, userId int32) ([]*MtEstimate, *MtDrug, *MtUser, error)
}

type EstimateService struct {
	repo EstimateRepo
	log  *log.Helper
}

// NewContentUsecase new a Content usecase.
func NewEstimateService(repo EstimateRepo, logger log.Logger) *EstimateService {
	return &EstimateService{repo: repo, log: log.NewHelper(logger)}
}

func (uc *EstimateService) CreateEstimate(ctx context.Context, estimate *MtEstimate) error {
	uc.log.WithContext(ctx).Infof("CreateEstimate: %v+v", estimate)
	return uc.repo.CreateEstimate(ctx, estimate)
}
func (uc *EstimateService) DeleteEstimate(ctx context.Context, id int32) error {
	uc.log.WithContext(ctx).Infof("DeleteEstimate: %v+v", id)
	return uc.repo.DeleteEstimate(ctx, id)
}
func (uc *EstimateService) GetEstimate(ctx context.Context, id int32, userId int32) (*MtEstimate, *MtUser, error) {
	uc.log.WithContext(ctx).Infof("GetEstimate: %v+v", id)
	return uc.repo.GetEstimate(ctx, id, userId)
}

func (uc *EstimateService) ListDrugEstimate(ctx context.Context, id int32, userId int32) ([]*MtEstimate, *MtDrug, *MtUser, error) {
	uc.log.WithContext(ctx).Infof("ListDrugEstimate: %v+v", id)
	return uc.repo.ListDrugEstimate(ctx, id, userId)
}
