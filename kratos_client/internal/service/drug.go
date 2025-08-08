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

// 搜索药品
func (s *DrugService) SearchDrugs(ctx context.Context, in *drup.SearchDrugsRequest) (*drup.SearchDrugsReply, error) {
	// 设置默认值
	page := int(in.Page)
	size := int(in.Size)
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	// 执行搜索
	drugs, total, err := s.uc.SearchDrugs(ctx, in.Keyword, in.CategoryId, page, size)
	if err != nil {
		return &drup.SearchDrugsReply{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	// 转换响应
	var drugInfos []*drup.SearchDrugInfo
	for _, drug := range drugs {
		drugInfos = append(drugInfos, &drup.SearchDrugInfo{
			Id:             int64(drug.Id),
			DrugName:       drug.DrugName,
			Specification:  drug.Specification,
			Price:          float64(drug.Price),
			Inventory:      int64(drug.Inventory),
			Manufacturer:   drug.Manufacturer,
			IsPrescription: false, // 暂时设为false，因为原表没有这个字段
			ExhibitionUrl:  strconv.Itoa(int(drug.ExhibitionId)),
		})
	}

	return &drup.SearchDrugsReply{
		Code:  0,
		Msg:   "success",
		Drugs: drugInfos,
		Total: total,
	}, nil
}

// 获取热门搜索
func (s *DrugService) GetHotSearch(ctx context.Context, in *drup.GetHotSearchRequest) (*drup.GetHotSearchReply, error) {
	limit := int(in.Limit)
	if limit <= 0 {
		limit = 10
	}

	// 获取热门关键词
	keywords, err := s.uc.GetHotKeywords(ctx, limit)
	if err != nil {
		return &drup.GetHotSearchReply{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	// 转换响应
	var hotKeywords []*drup.HotItem
	for i, keyword := range keywords {
		hotKeywords = append(hotKeywords, &drup.HotItem{
			Content: keyword,
			Count:   int64(100 - i*5), // 模拟计数
			Score:   float64(100 - i*5),
		})
	}

	// 模拟热门症状
	symptoms := []string{"感冒", "发烧", "咳嗽", "头痛", "胃痛"}
	var hotSymptoms []*drup.HotItem
	for i, symptom := range symptoms {
		if i >= limit {
			break
		}
		hotSymptoms = append(hotSymptoms, &drup.HotItem{
			Content: symptom,
			Count:   int64(80 - i*3),
			Score:   float64(80 - i*3),
		})
	}

	// 模拟热门问题
	questions := []string{"感冒药怎么选？", "发烧了吃什么药？", "咳嗽药有哪些？"}
	var hotQuestions []*drup.HotItem
	for i, question := range questions {
		if i >= limit {
			break
		}
		hotQuestions = append(hotQuestions, &drup.HotItem{
			Content: question,
			Count:   int64(60 - i*2),
			Score:   float64(60 - i*2),
		})
	}

	return &drup.GetHotSearchReply{
		Code:         0,
		Msg:          "success",
		HotKeywords:  hotKeywords,
		HotSymptoms:  hotSymptoms,
		HotQuestions: hotQuestions,
	}, nil
}


