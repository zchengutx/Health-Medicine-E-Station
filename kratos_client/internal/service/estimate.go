package service

import (
	"context"
	estimate "kratos_client/api/estimate/v1"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
)

type EstimateService struct {
	estimate.UnimplementedEstimateServer
	data *data.Data
	uc   *biz.EstimateService
}

func NewEstimateService(uc *biz.EstimateService, d *data.Data) *EstimateService {
	return &EstimateService{
		UnimplementedEstimateServer: estimate.UnimplementedEstimateServer{},
		data:                        d,
		uc:                          uc,
	}
}
func (s *EstimateService) ListDrugEstimate(ctx context.Context, in *estimate.ListDrugEstimateRequest) (*estimate.ListDrugEstimateReply, error) {
	estimates, drug, user, err := s.uc.ListDrugEstimate(ctx, int32(in.DrugId), int32(in.UserId))
	if err != nil {
		return nil, err
	}
	var estimateInfo []*estimate.EstimateInfo
	for _, e := range estimates {
		estimateInfo = append(estimateInfo, &estimate.EstimateInfo{
			Content:    e.Content,
			Image:      e.Image,
			Mobile:     user.Mobile,
			CreateTime: e.CreatedAt.Format("2006-01-02 15:04:05"),
			Avatar:     user.Avatar,
		})
	}
	return &estimate.ListDrugEstimateReply{
		Code:    0,
		Message: "list success",
		Info: &estimate.InfoDrugs{
			DrugName:      drug.DrugName,
			Specification: drug.Specification,
			Price:         float64(drug.Price),
			SalesVolume:   float64(drug.SalesVolume),
			Inventory:     int64(drug.Inventory),
		},
		List: estimateInfo,
	}, nil
}
func (s *EstimateService) GetEstimate(ctx context.Context, in *estimate.GetEstimateRequest) (*estimate.GetEstimateReply, error) {
	es, user, err := s.uc.GetEstimate(ctx, int32(in.Id), int32(in.UserId))
	if err != nil {
		return nil, err
	}
	return &estimate.GetEstimateReply{
		Code:    0,
		Message: "get success",
		Info: &estimate.EstimateInfo{
			Content:    es.Content,
			Image:      es.Image,
			Mobile:     user.Mobile,
			CreateTime: es.CreatedAt.Format("2006-01-02 15:04:05"),
			Avatar:     user.Avatar,
		},
	}, nil
}
func (s *EstimateService) CreateEstimate(ctx context.Context, in *estimate.CreateEstimateRequest) (*estimate.CreateEstimateReply, error) {
	err := s.uc.CreateEstimate(ctx, &biz.MtEstimate{
		DrugId:  int32(in.DrugId),
		Content: in.Content,
		Image:   in.Image,
	})
	if err != nil {
		return nil, err
	}
	return &estimate.CreateEstimateReply{
		Code:    0,
		Message: "create success",
	}, err
}
func (s *EstimateService) DeleteEstimate(ctx context.Context, in *estimate.DeleteEstimateRequest) (*estimate.DeleteEstimateReply, error) {
	err := s.uc.DeleteEstimate(ctx, int32(in.Id))
	if err != nil {
		return nil, err
	}
	return &estimate.DeleteEstimateReply{
		Code:    0,
		Message: "delete success",
	}, nil
}
