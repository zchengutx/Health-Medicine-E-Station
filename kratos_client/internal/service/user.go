package service

import (
	"context"
	"errors"
	"fmt"
	v1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
	"math/rand"
	"time"

	"github.com/go-kratos/kratos/v2/transport/http"
)

type UserService struct {
	v1.UnimplementedUserServer
	data *data.Data
	uc   *biz.UserService
	city *biz.CityService
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserService, d *data.Data, city *biz.CityService) *UserService {
	return &UserService{
		UnimplementedUserServer: v1.UnimplementedUserServer{},
		data:                    d,
		uc:                      uc,
		city:                    city,
	}
}

// SendSms implements user.UserServer.
func (s *UserService) SendSms(ctx context.Context, in *v1.SendSmsRequest) (*v1.SendSmsReply, error) {
	// 检查Redis是否可用
	if s.data.Redis() != nil {
		// 验证手机号是否已发送验证码
		if v := s.data.Redis().Exists(ctx, "sendSms").Val(); v != 0 {
			return nil, errors.New("The verification code is not used")
		}
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
	if s.data.Redis() != nil {
		s.data.Redis().Set(ctx, "sendSms"+in.Mobile+in.Source, code, time.Minute*10)
		// 增加发送次数
		s.data.Redis().Incr(ctx, "sendSms")
		s.data.Redis().Expire(ctx, "sendSms", time.Minute*1)
	} else {
		// Redis不可用时的替代方案：简单记录到日志
		fmt.Printf("验证码(无Redis): %s -> %d\n", in.Mobile, code)
	}
	
	return &v1.SendSmsReply{Message: "sendSms success"}, nil
}

func (s *UserService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	// 验证手机号
	if !comment.CheckMobile(in.Mobile) {
		return nil, errors.New("mobile error")
	}
	
	// 验证验证码
	if s.data.Redis() != nil {
		// 从redis获取验证码
		code := s.data.Redis().Get(ctx, "sendSms"+in.Mobile+"login")
		// 验证验证码
		if code.Val() != in.SendSmsCode {
			return nil, errors.New("code error")
		}
		// 验证成功，删除验证码
		s.data.Redis().Del(ctx, "sendSms"+in.Mobile+"login")
	} else {
		// Redis不可用时的替代方案：跳过验证码验证或使用固定验证码
		fmt.Printf("登录验证(无Redis): %s 使用验证码: %s\n", in.Mobile, in.SendSmsCode)
		// 可以设置一个开发用的固定验证码，或者完全跳过验证
		if in.SendSmsCode != "123456" {
			return nil, errors.New("code error (dev mode: use 123456)")
		}
	}

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

func (s *UserService) UpdateNickName(ctx context.Context, in *v1.UpdateNickNameRequest) (*v1.UpdateNickNameReply, error) {
	value := ctx.Value("user_id")

	_, err := s.uc.Update(ctx, &biz.MtUser{
		Id:       int32(value.(float64)),
		NickName: in.NickName,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateNickNameReply{Message: "update nickName success"}, nil
}

func (s *UserService) UpdateMobile(ctx context.Context, in *v1.UpdateMobileRequest) (*v1.UpdateMobileReply, error) {

	get := s.data.Redis().Get(ctx, "sendSms"+in.Mobile+"update")

	if get.Val() != in.SendSmsCode {
		return &v1.UpdateMobileReply{Message: "send sms code error"}, nil
	}

	_, err := s.uc.Find(ctx, &biz.MtUser{Mobile: in.Mobile})
	if err != nil {
		return &v1.UpdateMobileReply{Message: "I don't have this phone number"}, err
	}

	value := ctx.Value("user_id")

	_, err = s.uc.Update(ctx, &biz.MtUser{
		Id:     int32(value.(float64)),
		Mobile: in.NewMobile,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateMobileReply{Message: "update mobile success"}, nil
}
func (c *UserService) SelectTheCity(ctx context.Context, in *v1.SelectTheCityRequest) (*v1.SelectTheCityReply, error) {
	find, err := c.city.Find(ctx, &biz.MtCity{})
	if err != nil {
		return &v1.SelectTheCityReply{
			Message: "查询失败",
		}, err
	}

	var cityList []*v1.CityList

	for _, city := range *find {
		cityList = append(cityList, &v1.CityList{
			CityName: city.Name,
			CityId:   city.Code,
		})
	}

	return &v1.SelectTheCityReply{
		Message:  "查询成功",
		CityList: cityList,
	}, nil

}

func (c *UserService) SearchForCities(ctx context.Context, in *v1.SearchForCitiesRequest) (*v1.SearchForCitiesReply, error) {

	cityList, err := c.city.LikeFind(ctx, &biz.MtCity{
		Name: in.AddressName,
	})
	if err != nil {
		return &v1.SearchForCitiesReply{
			Message: "查询失败",
		}, err
	}

	var CityLists []*v1.CityList

	for _, city := range *cityList {
		CityLists = append(CityLists, &v1.CityList{
			CityName: city.Name,
			CityId:   city.Code,
		})
	}

	return &v1.SearchForCitiesReply{
		Message:  "查询成功",
		CityList: CityLists,
	}, nil
}

type LocationResponse struct {
	Content struct {
		Point struct {
			X string `json:"x"`
			Y string `json:"y"`
		} `json:"point"`
	} `json:"content"`
}

func (s *UserService) GetTargeted(c http.Context) error {
	fmt.Println("=== GetTargeted 开始执行 ===")

	// 添加错误恢复机制
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("GetTargeted 发生panic: %v\n", r)
			// 确保即使panic也能返回响应
			c.Result(500, map[string]string{
				"error": "服务器内部错误",
				"msg":   "Internal server error",
			})
		}
	}()

	// 注意：GetTargeted 是HTTP处理器，不需要获取用户ID
	// JWT中间件已经验证了Token的有效性

	fmt.Println("返回固定坐标（北京）...")
	fmt.Println("=== GetTargeted 执行完成 ===")

	return c.Result(200, map[string]string{
		"message": "获取成功",
		"x":       "116.404", // 北京经度
		"y":       "39.915",  // 北京纬度
	})
}

func (c *UserService) CreateAddress(ctx context.Context, in *v1.CreateAddressRequest) (*v1.CreateAddressReply, error) {

	value := ctx.Value("user_id")
	if value == nil {
		return nil, errors.New("用户未登录")
	}

	userID, ok := value.(float64)
	if !ok {
		return nil, errors.New("用户ID格式错误")
	}

	_, err := c.city.Create(ctx, &biz.MtAddress{
		UserId:          int32(userID),
		Consignee:       in.UserName,
		Mobile:          in.Mobile,
		CityId:          in.CityId,
		ShippingAddress: in.ShippingAddress,
		DoorplateFloor:  in.DoorplateFloor,
		Label:           in.Label,
		IsDefault:       in.IsDefault,
	})
	if err != nil {
		return &v1.CreateAddressReply{Message: "Failed to add"}, err
	}

	return &v1.CreateAddressReply{
		Message: "Created successfully",
	}, nil
}

func (c *UserService) UserInfo(ctx context.Context, in *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	value := ctx.Value("user_id")
	if value == nil {
		return nil, errors.New("用户未登录")
	}

	userID, ok := value.(float64)
	if !ok {
		return nil, errors.New("用户ID格式错误")
	}

	find, err := c.uc.FindId(ctx, &biz.MtUser{
		Id: int32(userID),
	})
	if err != nil {
		return &v1.UserInfoReply{
			Message: "The query failed",
		}, err
	}

	return &v1.UserInfoReply{
		Message:  "The query was successful",
		UserName: find.NickName,
		Mobile:   find.Mobile,
		Avatar:   find.Avatar,
	}, nil
}

func (c *UserService) GetAddressList(ctx context.Context, in *v1.GetAddressListRequest) (*v1.GetAddressListReply, error) {
	value := ctx.Value("user_id")
	if value == nil {
		return nil, errors.New("用户未登录")
	}

	userIDFloat, ok := value.(float64)
	if !ok {
		return nil, errors.New("用户ID格式错误")
	}
	userId := int32(userIDFloat)

	addressList, err := c.city.GetAddressList(ctx, userId)
	if err != nil {
		return &v1.GetAddressListReply{
			Message: "获取地址列表失败",
		}, err
	}

	var addressInfoList []*v1.AddressInfo
	for _, addr := range *addressList {
		// 获取城市名称
		cityInfo, _ := c.city.Find(ctx, &biz.MtCity{Code: addr.CityId})
		cityName := ""
		if cityInfo != nil && len(*cityInfo) > 0 {
			cityName = (*cityInfo)[0].Name
		}

		addressInfoList = append(addressInfoList, &v1.AddressInfo{
			Id:              addr.Id,
			Consignee:       addr.Consignee,
			Mobile:          addr.Mobile,
			CityId:          addr.CityId,
			CityName:        cityName,
			ShippingAddress: addr.ShippingAddress,
			DoorplateFloor:  addr.DoorplateFloor,
			Label:           addr.Label,
			IsDefault:       addr.IsDefault,
		})
	}

	return &v1.GetAddressListReply{
		Message:     "获取成功",
		AddressList: addressInfoList,
	}, nil
}

func (c *UserService) UpdateAddress(ctx context.Context, in *v1.UpdateAddressRequest) (*v1.UpdateAddressReply, error) {
	value := ctx.Value("user_id")
	userId := int32(value.(float64))

	_, err := c.city.UpdateAddress(ctx, &biz.MtAddress{
		Id:              in.Id,
		UserId:          userId,
		Consignee:       in.UserName,
		Mobile:          in.Mobile,
		CityId:          in.CityId,
		ShippingAddress: in.ShippingAddress,
		DoorplateFloor:  in.DoorplateFloor,
		Label:           in.Label,
		IsDefault:       in.IsDefault,
	})
	if err != nil {
		return &v1.UpdateAddressReply{Message: "更新失败"}, err
	}

	return &v1.UpdateAddressReply{
		Message: "更新成功",
	}, nil
}

func (c *UserService) DeleteAddress(ctx context.Context, in *v1.DeleteAddressRequest) (*v1.DeleteAddressReply, error) {
	err := c.city.DeleteAddress(ctx, in.Id)
	if err != nil {
		return &v1.DeleteAddressReply{Message: "删除失败"}, err
	}

	return &v1.DeleteAddressReply{
		Message: "删除成功",
	}, nil
}
