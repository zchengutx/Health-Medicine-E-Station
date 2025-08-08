package service

import (
	"context"
	pb "doctors/api/patient/v1"
	"doctors/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type PatientService struct {
	pb.UnimplementedPatientServer
	uc  *biz.PatientUsecase
	log *log.Helper
}

func NewPatientService(uc *biz.PatientUsecase, logger log.Logger) *PatientService {
	return &PatientService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

// CreatePatient 创建患者
func (s *PatientService) CreatePatient(ctx context.Context, req *pb.CreatePatientReq) (*pb.CreatePatientResp, error) {
	// 验证必填字段
	if req.Name == "" {
		return &pb.CreatePatientResp{
			Message: "患者姓名不能为空",
			Code:    400,
		}, nil
	}

	// 创建患者对象
	patient := &biz.Patient{
		Name:             req.Name,
		Gender:           req.Gender,
		Phone:            req.Phone,
		IdCard:           req.IdCard,
		Address:          req.Address,
		EmergencyContact: req.EmergencyContact,
		EmergencyPhone:   req.EmergencyPhone,
		MedicalHistory:   req.MedicalHistory,
		Allergies:        req.Allergies,
		Category:         req.Category,
		Remarks:          req.Remarks,
	}

	// 处理生日字段
	if req.BirthDate != "" {
		if birthDate, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			patient.BirthDate = &birthDate
		}
	}

	// 调用业务逻辑创建患者
	err := s.uc.CreatePatient(ctx, patient)
	if err != nil {
		s.log.WithContext(ctx).Errorf("创建患者失败: %v", err)

		if errors.Is(err, biz.ErrPatientAlreadyExists) {
			return &pb.CreatePatientResp{
				Message: "患者已存在",
				Code:    409,
			}, nil
		}

		if errors.Is(err, biz.ErrInvalidPatientData) {
			return &pb.CreatePatientResp{
				Message: "患者数据无效",
				Code:    400,
			}, nil
		}

		return &pb.CreatePatientResp{
			Message: "创建患者失败",
			Code:    500,
		}, nil
	}

	return &pb.CreatePatientResp{
		Message:   "创建成功",
		Code:      200,
		PatientId: int64(patient.ID),
	}, nil
}

// GetPatientProfile 获取患者档案
func (s *PatientService) GetPatientProfile(ctx context.Context, req *pb.GetPatientProfileReq) (*pb.GetPatientProfileResp, error) {
	patient, err := s.uc.GetPatientByID(ctx, uint(req.PatientId))
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取患者信息失败: %v", err)
		return &pb.GetPatientProfileResp{
			Message: "患者不存在",
			Code:    404,
		}, nil
	}

	// 构建响应数据
	profile := &pb.PatientProfile{
		PatientId:        int64(patient.ID),
		PatientCode:      patient.PatientCode,
		Name:             patient.Name,
		Gender:           patient.Gender,
		Age:              int32(patient.Age),
		Phone:            patient.Phone,
		IdCard:           patient.IdCard,
		Address:          patient.Address,
		EmergencyContact: patient.EmergencyContact,
		EmergencyPhone:   patient.EmergencyPhone,
		MedicalHistory:   patient.MedicalHistory,
		Allergies:        patient.Allergies,
		Category:         patient.Category,
		Status:           patient.Status,
		Remarks:          patient.Remarks,
		CreatedAt:        patient.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        patient.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// 处理生日字段
	if patient.BirthDate != nil {
		profile.BirthDate = patient.BirthDate.Format("2006-01-02")
	}

	return &pb.GetPatientProfileResp{
		Message: "获取成功",
		Code:    200,
		Profile: profile,
	}, nil
}

// UpdatePatientProfile 更新患者档案
func (s *PatientService) UpdatePatientProfile(ctx context.Context, req *pb.UpdatePatientProfileReq) (*pb.UpdatePatientProfileResp, error) {
	patient := &biz.Patient{
		ID:               uint(req.PatientId),
		Name:             req.Name,
		Gender:           req.Gender,
		Phone:            req.Phone,
		Address:          req.Address,
		EmergencyContact: req.EmergencyContact,
		EmergencyPhone:   req.EmergencyPhone,
		MedicalHistory:   req.MedicalHistory,
		Allergies:        req.Allergies,
		Remarks:          req.Remarks,
	}

	// 处理生日字段
	if req.BirthDate != "" {
		if birthDate, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			patient.BirthDate = &birthDate
		}
	}

	err := s.uc.UpdatePatientProfile(ctx, patient)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新患者信息失败: %v", err)

		if errors.Is(err, biz.ErrPatientNotFound) {
			return &pb.UpdatePatientProfileResp{
				Message: "患者不存在",
				Code:    404,
			}, nil
		}

		return &pb.UpdatePatientProfileResp{
			Message: "更新失败",
			Code:    500,
		}, nil
	}

	return &pb.UpdatePatientProfileResp{
		Message: "更新成功",
		Code:    200,
	}, nil
}

// GetPatientList 获取患者列表
func (s *PatientService) GetPatientList(ctx context.Context, req *pb.GetPatientListReq) (*pb.GetPatientListResp, error) {
	query := &biz.PatientListQuery{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		Keyword:  req.Keyword,
		Category: req.Category,
		Status:   req.Status,
	}

	result, err := s.uc.GetPatientList(ctx, query)
	if err != nil {
		s.log.WithContext(ctx).Errorf("获取患者列表失败: %v", err)
		return &pb.GetPatientListResp{
			Message: "获取患者列表失败",
			Code:    500,
		}, nil
	}

	// 转换为响应格式
	patients := make([]*pb.PatientProfile, len(result.Patients))
	for i, patient := range result.Patients {
		profile := &pb.PatientProfile{
			PatientId:        int64(patient.ID),
			PatientCode:      patient.PatientCode,
			Name:             patient.Name,
			Gender:           patient.Gender,
			Age:              int32(patient.Age),
			Phone:            patient.Phone,
			IdCard:           patient.IdCard,
			Address:          patient.Address,
			EmergencyContact: patient.EmergencyContact,
			EmergencyPhone:   patient.EmergencyPhone,
			MedicalHistory:   patient.MedicalHistory,
			Allergies:        patient.Allergies,
			Category:         patient.Category,
			Status:           patient.Status,
			Remarks:          patient.Remarks,
			CreatedAt:        patient.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        patient.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		// 处理生日字段
		if patient.BirthDate != nil {
			profile.BirthDate = patient.BirthDate.Format("2006-01-02")
		}

		patients[i] = profile
	}

	return &pb.GetPatientListResp{
		Message:  "获取成功",
		Code:     200,
		Patients: patients,
		Total:    result.Total,
		Page:     int32(result.Page),
		PageSize: int32(result.PageSize),
	}, nil
}

// GetPatientsByCategory 按分类获取患者
func (s *PatientService) GetPatientsByCategory(ctx context.Context, req *pb.GetPatientsByCategoryReq) (*pb.GetPatientsByCategoryResp, error) {
	result, err := s.uc.GetPatientsByCategory(ctx, req.Category, int(req.Page), int(req.PageSize))
	if err != nil {
		s.log.WithContext(ctx).Errorf("按分类获取患者失败: %v", err)
		return &pb.GetPatientsByCategoryResp{
			Message: "获取患者失败",
			Code:    500,
		}, nil
	}

	// 转换为响应格式
	patients := make([]*pb.PatientProfile, len(result.Patients))
	for i, patient := range result.Patients {
		profile := &pb.PatientProfile{
			PatientId:        int64(patient.ID),
			PatientCode:      patient.PatientCode,
			Name:             patient.Name,
			Gender:           patient.Gender,
			Age:              int32(patient.Age),
			Phone:            patient.Phone,
			IdCard:           patient.IdCard,
			Address:          patient.Address,
			EmergencyContact: patient.EmergencyContact,
			EmergencyPhone:   patient.EmergencyPhone,
			MedicalHistory:   patient.MedicalHistory,
			Allergies:        patient.Allergies,
			Category:         patient.Category,
			Status:           patient.Status,
			Remarks:          patient.Remarks,
			CreatedAt:        patient.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        patient.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		// 处理生日字段
		if patient.BirthDate != nil {
			profile.BirthDate = patient.BirthDate.Format("2006-01-02")
		}

		patients[i] = profile
	}

	return &pb.GetPatientsByCategoryResp{
		Message:  "获取成功",
		Code:     200,
		Patients: patients,
		Total:    result.Total,
	}, nil
}

// UpdatePatientCategory 更新患者分类
func (s *PatientService) UpdatePatientCategory(ctx context.Context, req *pb.UpdatePatientCategoryReq) (*pb.UpdatePatientCategoryResp, error) {
	err := s.uc.UpdatePatientCategory(ctx, uint(req.PatientId), req.Category)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新患者分类失败: %v", err)

		if errors.Is(err, biz.ErrPatientNotFound) {
			return &pb.UpdatePatientCategoryResp{
				Message: "患者不存在",
				Code:    404,
			}, nil
		}

		return &pb.UpdatePatientCategoryResp{
			Message: "更新分类失败",
			Code:    500,
		}, nil
	}

	return &pb.UpdatePatientCategoryResp{
		Message: "更新成功",
		Code:    200,
	}, nil
}
