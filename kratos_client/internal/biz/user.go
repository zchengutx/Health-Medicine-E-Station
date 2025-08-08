package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MtUser struct {
	Id       int32  `gorm:"column:id;type:int;comment:主键id;primaryKey;" json:"id"`          // 主键id
	NickName string `gorm:"column:nick_name;type:varchar(20);comment:昵称;" json:"nick_name"` // 昵称
	Mobile   string `gorm:"column:mobile;type:char(11);comment:手机号;" json:"mobile"`        // 手机号
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:头像;" json:"avatar"`      // 头像
}

func (u *MtUser) TableName() string {
	return "mt_user"
}

type UserRepo interface {
	Create(context.Context, *MtUser) (*MtUser, error)
	Find(context.Context, *MtUser) (*MtUser, error)
	Update(context.Context, *MtUser) (*MtUser, error)
	FindId(context.Context, *MtUser) (*MtUser, error)
}

type UserService struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *UserService) Create(ctx context.Context, req *MtUser) (*MtUser, error) {
	s.log.WithContext(ctx).Infof("Create user %+v", req.Mobile)
	return s.repo.Create(ctx, req)
}

func (s *UserService) Find(ctx context.Context, req *MtUser) (*MtUser, error) {
	s.log.WithContext(ctx).Infof("Find user %+v", req)
	return s.repo.Find(ctx, req)
}

func (s *UserService) FindId(ctx context.Context, req *MtUser) (*MtUser, error) {
	s.log.WithContext(ctx).Infof("Find id user %+v", req)
	return s.repo.FindId(ctx, req)
}

func (s *UserService) Update(ctx context.Context, req *MtUser) (*MtUser, error) {
	s.log.WithContext(ctx).Infof("Update user %+v", req)
	return s.repo.Update(ctx, req)
}
