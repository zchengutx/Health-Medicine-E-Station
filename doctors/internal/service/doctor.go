package service

import (
	"context"
	pb "doctors/api/doctor/v1"
	"doctors/internal/biz"
	"doctors/utils"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type DoctorService struct {
	pb.UnimplementedDoctorServer
	uc  *biz.DoctorUsecase
	log *log.Helper
}

func NewDoctorService(uc *biz.DoctorUsecase, logger log.Logger) *DoctorService {
	return &DoctorService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

// SendSms 发送短信验证码
func (s *DoctorService) SendSms(ctx context.Context, req *pb.SendSmsReq) (*pb.SendSmsResp, error) {
	if req.Phone == "" {
		return &pb.SendSmsResp{
			Message: "手机号不能为空",
			Code:    400,
		}, nil
	}

	err := s.uc.SendSmsCode(ctx, req.Phone)
	if err != nil {
		s.log.WithContext(ctx).Errorf("发送验证码失败: %v", err)
		return &pb.SendSmsResp{
			Message: "发送验证码失败",
			Code:    500,
		}, nil
	}

	return &pb.SendSmsResp{
		Message: "验证码发送成功",
		Code:    200,
	}, nil
}

// RegisterDoctor 医生注册
func (s *DoctorService) RegisterDoctor(ctx context.Context, req *pb.RegisterDoctorReq) (*pb.RegisterDoctorResp, error) {
	// 验证短信验证码
	valid, err := s.uc.VerifySmsCode(ctx, req.Phone, req.SendSmsCode)
	if err != nil || !valid {
		return &pb.RegisterDoctorResp{
			Message: "验证码错误或已过期",
			Code:    400,
		}, nil
	}

	// 创建医生对象
	doctor := &biz.Doctor{
		Phone:    req.Phone,
		Password: req.Password, // 密码将在业务层进行加密处理
		Status:   "启用",         // 启用状态
	}

	// 调用业务逻辑注册医生
	err = s.uc.RegisterDoctor(ctx, doctor)
	if err != nil {
		s.log.WithContext(ctx).Errorf("注册医生失败: %v", err)

		if errors.Is(err, biz.ErrPhoneAlreadyExists) {
			return &pb.RegisterDoctorResp{
				Message: "手机号已注册",
				Code:    409,
			}, nil
		}

		return &pb.RegisterDoctorResp{
			Message: "注册失败",
			Code:    500,
		}, nil
	}

	return &pb.RegisterDoctorResp{
		Message: "注册成功",
		Code:    200,
	}, nil
}

// LoginDoctor 医生登录
func (s *DoctorService) LoginDoctor(ctx context.Context, req *pb.LoginDoctorReq) (*pb.LoginDoctorResp, error) {
	// 如果提供了验证码，则进行验证码登录
	if req.SendSmsCode != "" {
		// 验证短信验证码
		valid, err := s.uc.VerifySmsCode(ctx, req.Phone, req.SendSmsCode)
		if err != nil || !valid {
			return &pb.LoginDoctorResp{
				Message: "验证码错误或已过期",
				Code:    400,
			}, nil
		}
		// 验证码登录：只需验证手机号存在
		doctor, err := s.uc.GetDoctorByPhone(ctx, req.Phone)
		if err != nil {
			s.log.WithContext(ctx).Errorf("验证码登录失败: %v", err)
			return &pb.LoginDoctorResp{
				Message: "医生不存在",
				Code:    404,
			}, nil
		}
		return &pb.LoginDoctorResp{
			Message: "登录成功",
			Code:    200,
			DId:     int64(doctor.ID),
		}, nil
	}

	// 密码登录
	doctor, err := s.uc.LoginDoctor(ctx, req.Phone, req.Password)
	if err != nil {
		s.log.WithContext(ctx).Errorf("登录失败: %v", err)

		if errors.Is(err, biz.ErrDoctorNotFound) {
			return &pb.LoginDoctorResp{
				Message: "医生不存在",
				Code:    404,
			}, nil
		}

		if errors.Is(err, biz.ErrInvalidPassword) {
			return &pb.LoginDoctorResp{
				Message: "密码错误",
				Code:    401,
			}, nil
		}

		return &pb.LoginDoctorResp{
			Message: "登录失败",
			Code:    500,
		}, nil
	}

	// 登录成功后生成 JWT token
	token, err := utils.TokenHandler(int64(doctor.ID))
	if err != nil {
		s.log.WithContext(ctx).Errorf("生成JWT失败: %v", err)
		return &pb.LoginDoctorResp{
			Message: "登录失败",
			Code:    500,
		}, nil
	}

	return &pb.LoginDoctorResp{
		Message: "登录成功",
		Code:    200,
		DId:     int64(doctor.ID),
		Token:   token,
	}, nil
}

// Authentication 医生认证
func (s *DoctorService) Authentication(ctx context.Context, req *pb.AuthenticationReq) (*pb.AuthenticationResp, error) {
	// 获取医生信息
	doctor, err := s.uc.GetDoctorByID(ctx, uint(req.DId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取医生信息失败: %v", err)
		return &pb.AuthenticationResp{
			Message: "医生不存在",
			Code:    404,
		}, nil
	}

	// 更新医生信息
	doctor.Name = req.Name
	doctor.Gender = req.Gender
	doctor.Email = req.Email
	doctor.Avatar = req.Avatar
	doctor.LicenseNumber = req.LicenseNumber

	// 处理可选字段
	if req.DepartmentId > 0 {
		departmentId := uint(req.DepartmentId)
		doctor.DepartmentID = &departmentId
	} else {
		doctor.DepartmentID = nil
	}
	if req.HospitalId > 0 {
		hospitalId := uint(req.HospitalId)
		doctor.HospitalID = &hospitalId
	} else {
		doctor.HospitalID = nil
	}

	doctor.Title = req.Title
	doctor.Speciality = req.Speciality
	doctor.PracticeScope = req.PracticeScope

	// 调用业务逻辑更新医生信息
	err = s.uc.UpdateDoctorInfo(ctx, doctor)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新医生信息失败: %v", err)
		return &pb.AuthenticationResp{
			Message: "认证失败",
			Code:    500,
		}, nil
	}

	return &pb.AuthenticationResp{
		Message: "认证成功",
		Code:    200,
	}, nil
}

// GetDoctorProfile 获取医生个人信息
func (s *DoctorService) GetDoctorProfile(ctx context.Context, req *pb.GetDoctorProfileReq) (*pb.GetDoctorProfileResp, error) {
	doctor, err := s.uc.GetDoctorByID(ctx, uint(req.DoctorId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取医生信息失败: %v", err)
		if err.Error() == "医生不存在" || err.Error() == "查询医生失败: 医生不存在" {
			return &pb.GetDoctorProfileResp{
				Message: "医生不存在",
				Code:    404,
			}, nil
		}
		return &pb.GetDoctorProfileResp{
			Message: "查询医生信息失败，请稍后重试",
			Code:    500,
		}, nil
	}

	// 构建响应数据
	profile := &pb.DoctorProfile{
		DId:           int64(doctor.ID),
		DoctorCode:    doctor.DoctorCode,
		Name:          doctor.Name,
		Gender:        doctor.Gender,
		Phone:         doctor.Phone,
		Email:         doctor.Email,
		Avatar:        doctor.Avatar,
		LicenseNumber: doctor.LicenseNumber,
		Title:         doctor.Title,
		Speciality:    doctor.Speciality,
		PracticeScope: doctor.PracticeScope,
		Status:        doctor.Status,
		CreatedAt:     doctor.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     doctor.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 处理可选字段
	if doctor.BirthDate != "" {
		profile.BirthDate = doctor.BirthDate
	}
	if doctor.DepartmentID != nil {
		profile.DepartmentId = int64(*doctor.DepartmentID)
	}
	if doctor.HospitalID != nil {
		profile.HospitalId = int64(*doctor.HospitalID)
	}

	return &pb.GetDoctorProfileResp{
		Message: "获取成功",
		Code:    200,
		Profile: profile,
	}, nil
}

// UpdateDoctorProfile 更新医生个人信息
func (s *DoctorService) UpdateDoctorProfile(ctx context.Context, req *pb.UpdateDoctorProfileReq) (*pb.UpdateDoctorProfileResp, error) {
	doctor, err := s.uc.GetDoctorByID(ctx, uint(req.DId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取医生信息失败: %v", err)
		return &pb.UpdateDoctorProfileResp{
			Message: "医生不存在",
			Code:    404,
		}, nil
	}

	// 合并新参数
	if req.Name != "" {
		doctor.Name = req.Name
	}
	if req.Gender != "" {
		doctor.Gender = req.Gender
	}
	if req.Email != "" {
		doctor.Email = req.Email
	}
	if req.Avatar != "" {
		doctor.Avatar = req.Avatar
	}
	if req.Title != "" {
		doctor.Title = req.Title
	}
	if req.Speciality != "" {
		doctor.Speciality = req.Speciality
	}
	if req.PracticeScope != "" {
		doctor.PracticeScope = req.PracticeScope
	}
	if req.BirthDate != "" {
		if req.BirthDate == "0000-00-00" || req.BirthDate == "0001/01/01" || req.BirthDate == "0001-01-01" {
			doctor.BirthDate = ""
		} else if _, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			doctor.BirthDate = req.BirthDate
		} else {
			// 解析失败时也设置为空字符串
			doctor.BirthDate = ""
		}
	} else {
		// 空字符串时保持为空字符串
		doctor.BirthDate = ""
	}

	err = s.uc.UpdateDoctorProfile(ctx, doctor)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新医生信息失败: %v", err)

		if errors.Is(err, biz.ErrDoctorNotFound) {
			return &pb.UpdateDoctorProfileResp{
				Message: "医生不存在",
				Code:    404,
			}, nil
		}

		return &pb.UpdateDoctorProfileResp{
			Message: "更新失败",
			Code:    500,
		}, nil
	}

	return &pb.UpdateDoctorProfileResp{
		Message: "更新成功",
		Code:    200,
	}, nil
}

// ChangePassword 修改密码
func (s *DoctorService) ChangePassword(ctx context.Context, req *pb.ChangePasswordReq) (*pb.ChangePasswordResp, error) {
	// 验证新密码和确认密码是否一致
	if req.NewPassword != req.ConfirmPassword {
		return &pb.ChangePasswordResp{
			Message: "新密码和确认密码不一致",
			Code:    400,
		}, nil
	}

	// 验证密码长度
	if len(req.NewPassword) < 6 {
		return &pb.ChangePasswordResp{
			Message: "密码长度不能少于6位",
			Code:    400,
		}, nil
	}

	err := s.uc.ChangePassword(ctx, uint(req.DId), req.OldPassword, req.NewPassword)
	if err != nil {
		s.log.WithContext(ctx).Errorf("修改密码失败: %v", err)

		if errors.Is(err, biz.ErrDoctorNotFound) {
			return &pb.ChangePasswordResp{
				Message: "医生不存在",
				Code:    404,
			}, nil
		}

		if errors.Is(err, biz.ErrInvalidPassword) {
			return &pb.ChangePasswordResp{
				Message: "原密码错误",
				Code:    401,
			}, nil
		}

		return &pb.ChangePasswordResp{
			Message: "修改密码失败",
			Code:    500,
		}, nil
	}

	return &pb.ChangePasswordResp{
		Message: "密码修改成功",
		Code:    200,
	}, nil
}

// DeleteAccount 注销医生账号
func (s *DoctorService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountReq) (*pb.DeleteAccountResp, error) {
	err := s.uc.DeleteDoctorByID(ctx, uint(req.DId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("注销账号失败: %v", err)
		return &pb.DeleteAccountResp{
			Message: "注销失败",
			Code:    500,
		}, nil
	}
	return &pb.DeleteAccountResp{
		Message: "注销成功",
		Code:    200,
	}, nil
}
