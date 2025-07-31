package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int32  `gorm:"column:id;type:int;comment:主键id;primaryKey;" json:"id"`          // 主键id
	NickName string `gorm:"column:nick_name;type:varchar(20);comment:昵称;" json:"nick_name"` // 昵称
	Mobile   string `gorm:"column:mobile;type:char(11);comment:手机号;" json:"mobile"`         // 手机号
	Avatar   string `gorm:"column:avatar;type:varchar(255);comment:头像;" json:"avatar"`      // 头像
}

func (u *User) TableName() string {
	return "user"
}

type UserRepo interface {
	Create(context.Context, *User) (*User, error)
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

func (s *UserService) Create(ctx context.Context, req *User) (*User, error) {
	s.log.WithContext(ctx).Infof("Create user %+v", req.Mobile)
	return s.repo.Create(ctx, req)
}
