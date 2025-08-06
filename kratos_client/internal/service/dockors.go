package service

import (
	"context"
	v1 "kratos_client/api/doctors/v1"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
)

type DoctorsService struct {
	v1.UnimplementedDoctorsServer
	data *data.Data
	uc   *biz.DoctorsService
}

// NewUserService new a user service.
func NewDoctorsService(uc *biz.DoctorsService, d *data.Data) *DoctorsService {
	return &DoctorsService{
		UnimplementedDoctorsServer: v1.UnimplementedDoctorsServer{},
		data:                       d,
		uc:                         uc,
	}
}
func (c *UserService) DoctorsList(ctx context.Context, in *v1.DoctorsListRequest) (*v1.DoctorsListReply, error) {
	find, err := c.city.Find(ctx, &biz.MtCity{})
	if err != nil {
		return &v1.DoctorsListReply{
			Message: "查询失败",
		}, err
	}

	var cityList []*v1.DoctorsList

	for _, city := range *find {
		cityList = append(cityList, &v1.DoctorsList{
			DoctorCode:  string(city.Code),
			DoctorsName: city.Name,
		})
	}

	return &v1.DoctorsListReply{
		Message:     "查询成功",
		DoctorsList: cityList,
	}, nil

}
