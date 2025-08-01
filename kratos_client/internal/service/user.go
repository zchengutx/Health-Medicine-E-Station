package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
	"math/rand"
	"time"
)

type UserService struct {
	v1.UnimplementedUserServer
	data *data.Data
	uc   *biz.UserService
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserService, d *data.Data) *UserService {
	return &UserService{
		UnimplementedUserServer: v1.UnimplementedUserServer{},
		data:                    d,
		uc:                      uc,
	}
}

// SendSms implements user.UserServer.
func (s *UserService) SendSms(ctx context.Context, in *v1.SendSmsRequest) (*v1.SendSmsReply, error) {
	// 验证手机号是否已发送验证码
	if v := s.data.Redis().Exists(ctx, "sendSms").Val(); v != 0 {
		return nil, errors.New("The verification code is not used")
	}
	// 验证手机号
	if !comment.CheckMobile(in.Mobile) {
		return nil, errors.New("mobile error")
	}
	//生成验证码
	code := rand.Intn(900000) + 100000
	// 发送短信
	//request := comment.SendSmsRequest{
	//	PhoneNumbers:  in.Mobile,
	//	TemplateParam: fmt.Sprintf(`{"code":"%s"}`, code),
	//}
	//sms, err := comment.SendSms(request)
	//if err != nil {
	//	return nil, errors.New("send sms error")
	//}
	//
	//if sms.Code != "OK" {
	//	return nil, errors.New("send sms error")
	//}
	// 保存验证码
	s.data.Redis().Set(ctx, "sendSms"+in.Mobile+in.Source, code, time.Minute*10)
	// 增加发送次数
	s.data.Redis().Incr(ctx, "sendSms")
	s.data.Redis().Expire(ctx, "sendSms", time.Minute*1)
	return &v1.SendSmsReply{Message: "sendSms success"}, nil
}

func (s *UserService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	// 验证手机号
	if !comment.CheckMobile(in.Mobile) {
		return nil, errors.New("mobile error")
	}
	// 从redis获取验证码
	code := s.data.Redis().Get(ctx, "sendSms"+in.Mobile+"login")
	// 验证验证码
	if code.Val() != in.SendSmsCode {
		return nil, errors.New("code error")
	}
	// 验证成功，删除验证码
	s.data.Redis().Del(ctx, "sendSms"+in.Mobile+"login")

	// 验证成功，查询用户
	find, err := s.uc.Find(ctx, &biz.MtUser{Mobile: in.Mobile})

	if err != nil {
		// 不存在用户，创建用户
		find, err = s.uc.Create(ctx, &biz.MtUser{
			Mobile: in.Mobile,
			Avatar: "5Gm3nxvHbG.jpg",
		})
	}
	// 验证成功，返回token

	token, err := comment.TokenHandler(find.Id)
	fmt.Println(token, err)

	return &v1.LoginReply{
		Message: "login success",
		Token:   token,
	}, nil
}

func (s *UserService) Upload(c http.Context) error {
	req := c.Request()
	err, s2 := comment.Upload(c, req)
	if err != nil {
		return err
	}
	value := req.Context().Value("user_id")
	// 上传成功，更新用户头像

	_, err = s.uc.Update(context.Background(), &biz.MtUser{
		Id:     int32(value.(float64)),
		Avatar: s2,
	})
	if err != nil {
		return err
	}

	return c.Result(200, map[string]interface{}{
		"message": "upload success",
		"url":     s2,
	})
}
