package service

import (
	"context"
	"strconv"

	drup "kratos_client/api/drug/v1"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
)

type DrugService struct {
	drup.UnimplementedDrugServer
	data *data.Data
	uc   *biz.DrugService
}

// NewAppService new a app service.
func NewDrugService(uc *biz.DrugService, d *data.Data) *DrugService {
	return &DrugService{
		UnimplementedDrugServer: drup.UnimplementedDrugServer{},
		data:                    d,
		uc:                      uc,
	}
}
func (s *DrugService) ListDrug(ctx context.Context, in *drup.ListDrugRequest) (*drup.ListDrugReply, error) {
	drugs, err := s.uc.ListDrug(ctx, int32(in.FristCategoryId), int32(in.SecondCategoryId), in.Keyword)
	if err != nil {
		return nil, err
	}
	var drugInfo []*drup.InfoDrugs
	for _, drug := range drugs {
		drugInfo = append(drugInfo, &drup.InfoDrugs{
			DrugName:      drug.DrugName,
			Specification: drug.Specification,
			Price:         float64(drug.Price),
			SalesVolume:   float64(drug.SalesVolume),
			Inventory:     int64(drug.Inventory),
			ExhibitionUrl: strconv.Itoa(int(drug.ExhibitionId)),
		})
	}
	return &drup.ListDrugReply{
		Code: 0,
		Msg:  "success",
		Drug: drugInfo,
	}, nil
}
func (s *DrugService) GetDrug(ctx context.Context, in *drup.GetDrugRequest) (*drup.GetDrugReply, error) {
	drug, err := s.uc.GetDrug(ctx, int32(in.Id))
	if err != nil {
		return nil, err
	}
	guide, err := s.uc.GetGuide(ctx, int32(in.Id))
	if err != nil {
		return nil, err
	}
	return &drup.GetDrugReply{
		Code: 0,
		Msg:  "success",
		Drug: &drup.InfoDrug{
			DrugName:      drug.DrugName,
			Specification: drug.Specification,
			Price:         float64(drug.Price),
			SalesVolume:   float64(drug.SalesVolume),
			Inventory:     int64(drug.Inventory),
			ExhibitionUrl: strconv.Itoa(int(drug.ExhibitionId)),
			Guide: &drup.GuideInfo{
				MajorFunction:  guide.MajorFunction,
				UsageAndDosage: guide.UsageAndDosage,
				Taboos:         guide.Taboos,
				SpecialCrowd:   guide.SpecialCrowd,
				Condition:      guide.Condition,
			},
		},
	}, nil
}

func (s *DrugService) GetExplain(ctx context.Context, in *drup.GetExplainRequest) (*drup.GetExplainReply, error) {
	explain, err := s.uc.GetExplain(ctx, int32(in.Id))
	if err != nil {
		return nil, err
	}
	return &drup.GetExplainReply{
		Code: 0,
		Msg:  "success",
		Explain: &drup.ExplainInfo{
			CommonName:             explain.CommonName,
			GoodsName:              explain.GoodsName,
			Component:              explain.Component,
			Taboos:                 explain.Taboos,
			Function:               explain.Function,
			UsageAndDosage:         explain.UsageAndDosage,
			Character:              explain.Character,
			PackagingSpecification: explain.PackagingSpecification,
			BadnessReaction:        explain.BadnessReaction,
			Condition:              explain.Condition,
			ValidTime:              explain.ValidTime,
			Notice:                 explain.Notice,
			Interaction:            explain.Interaction,
			RatifyNumber:           explain.RatifyNumber,
			Manufacturer:           explain.Manufacturer,
			StandardNumber:         explain.StandardNumber,
			Possessor:              explain.Possessor,
			Address:                explain.Address,
			Specification:          explain.Specification,
			DosageForm:             explain.DosageForm,
		},
	}, nil
}
func (s *DrugService) GetGuide(ctx context.Context, in *drup.GetGuideRequest) (*drup.GetGuideReply, error) {
	guide, err := s.uc.GetGuide(ctx, int32(in.Id))
	if err != nil {
		return nil, err
	}
	return &drup.GetGuideReply{
		Code: 0,
		Msg:  "success",
		Guide: &drup.GuideInfo{
			MajorFunction:  guide.MajorFunction,
			UsageAndDosage: guide.UsageAndDosage,
			Taboos:         guide.Taboos,
			SpecialCrowd:   guide.SpecialCrowd,
			Condition:      guide.Condition,
		},
	}, nil
}
