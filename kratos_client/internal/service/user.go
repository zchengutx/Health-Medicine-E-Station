package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io/ioutil"
	v1 "kratos_client/api/user/v1"
	"kratos_client/comment"
	"kratos_client/internal/biz"
	"kratos_client/internal/data"
	"math/rand"
	https "net/http"
	"net/url"
	"time"
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

	// 此处填写您在控制台-应用管理-创建应用后获取的AK
	ak := "fbZPlNxNhl49M4bg6zHwLTP0DpJym7eS"

	// 服务地址
	host := "https://api.map.baidu.com"

	// 接口地址
	uri := "/location/ip"

	ip, err := comment.GetPublicIP()
	fmt.Println(ip)
	if err != nil {
		return c.Result(400, map[string]string{
			"error": err.Error(),
			"msg":   "Failed to obtain IP",
		})
	}

	fmt.Println(ip)

	// 设置请求参数
	params := url.Values{
		"ip":   []string{ip},
		"coor": []string{"bd09ll"},
		"ak":   []string{ak},
	}

	// 发起请求
	request, err := url.Parse(host + uri + "?" + params.Encode())
	if nil != err {
		fmt.Printf("host error: %v", err)
		return c.Result(400, map[string]string{
			"error": err.Error(),
			"msg":   "Initiating a request failed",
		})
	}

	resp, err1 := https.Get(request.String())
	fmt.Printf("url: %s\n", request.String())
	defer resp.Body.Close()
	if err1 != nil {
		fmt.Printf("request error: %v", err1)
		return c.Result(400, map[string]string{
			"error": err.Error(),
			"msg":   "Failed to get response",
		})
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Printf("response error: %v", err2)
	}

	var location LocationResponse

	json.Unmarshal(body, &location)

	return c.Result(200, map[string]string{
		"message": "获取成功",
		"x":       location.Content.Point.X,
		"y":       location.Content.Point.Y,
	})
}

func (c *UserService) CreateAddress(ctx context.Context, in *v1.CreateAddressRequest) (*v1.CreateAddressReply, error) {

	value := ctx.Value("user_id")

	_, err := c.city.Create(ctx, &biz.MtAddress{
		UserId:          int32(value.(float64)),
		Consignee:       in.UserName,
		Mobile:          in.Mobile,
		CityId:          in.CityId,
		ShippingAddress: in.ShippingAddress,
		DoorplateFloor:  in.DoorplateFloor,
		Label:           in.Label,
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

	find, err := c.uc.FindId(ctx, &biz.MtUser{
		Id: int32(value.(float64)),
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
