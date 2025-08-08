package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	pb "kratos_client/api/prescription/v1"
	"kratos_client/internal/biz"
)

// PrescriptionService 处方服务
type PrescriptionService struct {
	pb.UnimplementedPrescriptionServiceServer

	prescriptionUc *biz.PrescriptionUsecase
	log            *log.Helper
}

// NewPrescriptionService 创建处方服务
func NewPrescriptionService(prescriptionUc *biz.PrescriptionUsecase, logger log.Logger) *PrescriptionService {
	return &PrescriptionService{
		prescriptionUc: prescriptionUc,
		log:            log.NewHelper(logger),
	}
}

// ListPrescriptions 获取处方列表
func (s *PrescriptionService) ListPrescriptions(ctx context.Context, req *pb.ListPrescriptionsRequest) (*pb.ListPrescriptionsReply, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	// 解析日期
	var startDate, endDate *time.Time
	if req.StartDate != "" {
		if t, err := time.Parse("2006-01-02", req.StartDate); err == nil {
			startDate = &t
		}
	}
	if req.EndDate != "" {
		if t, err := time.Parse("2006-01-02", req.EndDate); err == nil {
			endDate = &t
		}
	}

	bizReq := &biz.ListPrescriptionsRequest{
		Status:           req.Status,
		PrescriptionType: req.PrescriptionType,
		StartDate:        startDate,
		EndDate:          endDate,
		Page:             page,
		PageSize:         pageSize,
	}

	prescriptions, total, err := s.prescriptionUc.ListPrescriptions(ctx, bizReq)
	if err != nil {
		s.log.Errorf("获取处方列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	pbPrescriptions := make([]*pb.Prescription, len(prescriptions))
	for i, prescription := range prescriptions {
		pbPrescriptions[i] = s.convertToPbPrescription(prescription)
	}

	return &pb.ListPrescriptionsReply{
		Prescriptions: pbPrescriptions,
		Total:         total,
	}, nil
}

// ListPatientPrescriptions 获取患者处方列表
func (s *PrescriptionService) ListPatientPrescriptions(ctx context.Context, req *pb.ListPatientPrescriptionsRequest) (*pb.ListPatientPrescriptionsReply, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	bizReq := &biz.ListPatientPrescriptionsRequest{
		PatientID: req.PatientId,
		Status:    req.Status,
		Page:      page,
		PageSize:  pageSize,
	}

	prescriptions, total, err := s.prescriptionUc.ListPatientPrescriptions(ctx, bizReq)
	if err != nil {
		s.log.Errorf("获取患者处方列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	pbPrescriptions := make([]*pb.Prescription, len(prescriptions))
	for i, prescription := range prescriptions {
		pbPrescriptions[i] = s.convertToPbPrescription(prescription)
	}

	return &pb.ListPatientPrescriptionsReply{
		Prescriptions: pbPrescriptions,
		Total:         total,
	}, nil
}

// ListDoctorPrescriptions 获取医生处方列表
func (s *PrescriptionService) ListDoctorPrescriptions(ctx context.Context, req *pb.ListDoctorPrescriptionsRequest) (*pb.ListDoctorPrescriptionsReply, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	bizReq := &biz.ListDoctorPrescriptionsRequest{
		DoctorID: req.DoctorId,
		Status:   req.Status,
		Page:     page,
		PageSize: pageSize,
	}

	prescriptions, total, err := s.prescriptionUc.ListDoctorPrescriptions(ctx, bizReq)
	if err != nil {
		s.log.Errorf("获取医生处方列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	pbPrescriptions := make([]*pb.Prescription, len(prescriptions))
	for i, prescription := range prescriptions {
		pbPrescriptions[i] = s.convertToPbPrescription(prescription)
	}

	return &pb.ListDoctorPrescriptionsReply{
		Prescriptions: pbPrescriptions,
		Total:         total,
	}, nil
}

// GetPrescriptionDetail 获取处方详情
func (s *PrescriptionService) GetPrescriptionDetail(ctx context.Context, req *pb.GetPrescriptionDetailRequest) (*pb.GetPrescriptionDetailReply, error) {
	detail, err := s.prescriptionUc.GetPrescriptionDetail(ctx, req.PrescriptionId)
	if err != nil {
		s.log.Errorf("获取处方详情失败: %v", err)
		return nil, err
	}

	// 转换处方信息
	pbPrescription := s.convertToPbPrescription(detail.Prescription)

	// 转换药品明细
	pbMedicines := make([]*pb.PrescriptionMedicine, len(detail.Medicines))
	for i, medicine := range detail.Medicines {
		pbMedicines[i] = s.convertToPbPrescriptionMedicine(medicine)
	}

	return &pb.GetPrescriptionDetailReply{
		Prescription: pbPrescription,
		Medicines:    pbMedicines,
	}, nil
}

// 转换处方为protobuf格式
func (s *PrescriptionService) convertToPbPrescription(prescription *biz.MtPrescription) *pb.Prescription {
	pbPrescription := &pb.Prescription{
		Id:               prescription.ID,
		PrescriptionNo:   prescription.PrescriptionNo,
		DoctorId:         prescription.DoctorID,
		PatientId:        prescription.PatientID,
		PrescriptionDate: prescription.PrescriptionDate.Format("2006-01-02"),
		TotalAmount:      prescription.TotalAmount.String(),
		PrescriptionType: prescription.PrescriptionType,
		UsageInstruction: prescription.UsageInstruction,
		Status:           prescription.Status,
		AuditNotes:       prescription.AuditNotes,
		CreatedAt:        prescription.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        prescription.UpdatedAt.Format(time.RFC3339),
		DoctorName:       prescription.DoctorName,
		PatientName:      prescription.PatientName,
		AuditorName:      prescription.AuditorName,
		MedicineCount:    prescription.MedicineCount,
	}

	if prescription.MedicalRecordID != nil {
		pbPrescription.MedicalRecordId = *prescription.MedicalRecordID
	}

	if prescription.AuditorID != nil {
		pbPrescription.AuditorId = *prescription.AuditorID
	}

	if prescription.AuditTime != nil {
		pbPrescription.AuditTime = prescription.AuditTime.Format(time.RFC3339)
	}

	return pbPrescription
}

// 转换处方药品明细为protobuf格式
func (s *PrescriptionService) convertToPbPrescriptionMedicine(medicine *biz.MtPrescriptionMedicine) *pb.PrescriptionMedicine {
	return &pb.PrescriptionMedicine{
		Id:             medicine.ID,
		PrescriptionId: medicine.PrescriptionID,
		MedicineId:     medicine.MedicineID,
		Quantity:       medicine.Quantity.String(),
		Unit:           medicine.Unit,
		UnitPrice:      medicine.UnitPrice.String(),
		TotalPrice:     medicine.TotalPrice.String(),
		Dosage:         medicine.Dosage,
		Frequency:      medicine.Frequency,
		Duration:       medicine.Duration,
		UsageMethod:    medicine.UsageMethod,
		Notes:          medicine.Notes,
		CreatedAt:      medicine.CreatedAt.Format(time.RFC3339),
		MedicineName:   medicine.MedicineName,
		MedicineSpec:   medicine.MedicineSpec,
		Manufacturer:   medicine.Manufacturer,
	}
}