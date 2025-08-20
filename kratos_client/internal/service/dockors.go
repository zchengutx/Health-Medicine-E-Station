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
func (c *DoctorsService) DoctorsList(ctx context.Context, in *v1.DoctorsListRequest) (*v1.DoctorsListReply, error) {
	// 查询所有医生
	find, err := c.uc.DoctorsFind(ctx, &biz.MtDoctors{})
	if err != nil {
		return &v1.DoctorsListReply{
			Message: "查询失败",
		}, err
	}

	var doctorsList []*v1.DoctorsList

	for _, doctor := range *find {
		doctorsList = append(doctorsList, &v1.DoctorsList{
			DoctorCode:  doctor.DoctorCode,
			DoctorsName: doctor.Name,
		})
	}

	return &v1.DoctorsListReply{
		Message:     "查询成功",
		DoctorsList: doctorsList,
	}, nil
}
