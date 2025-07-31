package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	v1 "kratos_client/api/user/v1"
	"kratos_client/internal/biz"
	"math/rand"
	"time"
)

type UserService struct {
	v1.UnimplementedUserServer
	Rdb *redis.Client
	uc  *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.GreeterUsecase, Rdb *redis.Client) *UserService {
	return &UserService{
		UnimplementedUserServer: v1.UnimplementedUserServer{},
		Rdb:                     Rdb,
		uc:                      uc,
	}
}

// SayHello implements helloworld.GreeterServer.
func (s *UserService) SendSms(ctx context.Context, in *v1.SendSmsRequest) (*v1.SendSmsReply, error) {
	// 生成验证码
	code := rand.Intn(9000000) + 1000000
	// 保存验证码
	s.Rdb.Set(ctx, "sendSms"+in.Mobile+in.Source, code, time.Minute*10)

	return &v1.SendSmsReply{Message: "sendSms "}, nil
}
