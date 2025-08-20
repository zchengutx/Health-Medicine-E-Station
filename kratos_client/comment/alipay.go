package comment

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/smartwalle/alipay/v3"
)

// 支付宝配置结构
type AlipayConfig struct {
	AppID           string // 应用ID
	PrivateKey      string // 应用私钥
	AlipayPublicKey string // 支付宝公钥
	IsProduction    bool   // 是否生产环境
	NotifyURL       string // 异步通知地址
	ReturnURL       string // 同步返回地址
}

// 支付宝客户端
var alipayClient *alipay.Client

// 使用你提供的密钥配置
var defaultConfig = AlipayConfig{
	AppID: "9021000140671234", // 请替换为你的实际沙箱AppID
	// 你的应用私钥
	PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQChuX8kFszTtwxS
gD0y/sM6D7p9zpvDChlFpAujWOvEisUSrm6IIBvtf/0xHWayXcxMLmh5mgE7Cyjr
W6YID9B0hrEQuSaC5PQ99ddHpR4qg69m+KNQ5duagBqz16wBRGkF4FLRka1InDPL
36qMsYclavRdMpPdeoDQLgi3Sdl/sXMNHKYVLUhhRIjon1MHwD3dGf3eHLmns580
azxLyFQxgB74kfOFpiJ1gMLcdfSpeuXujvxhyY+f4noms5RQBh4NwlFMJkpyyvcR
HPgyVN+lUVkIgYyxLJSAV6gAdnMx+27Rou1ZFHSKxo+nNUVKbry7hkMxH9BvuZ5g
XlxWoPTbAgMBAAECggEAdP9xj3Y/MFsYuwazP5U3P2XpkOJLpUpFBjCrirzltAaA
lAdFR42TJrqVPVb72MYq6mIYiwBzK3fjXoGrF+H4+JQIvQR1a/SfDcQwvlAiBrfF
yUTPQdNIj/llV/4LHc+T+wBSafJt3j3C6xcglzBHiTZbGqFgf7YEQpdLDu6KPunF
z6O9YK3ZsYsyYQcY7CNJJotus49qz6q54fk2hIwj+4MnQSblwcGixj4IB59a8Y5E
JKCk6LhwdjOuxHrEwRXQZfXCYCwz+ox55YXZDGqGSHBLC+7Mpc42xqSMRpp6SHO0
QhFFobyX5FtzjoGbsyHB3EgCoYUnaifecYSAgXDt4QKBgQDddcVS3c0UOWeXmiGm
5N9l6yxzT5Q6Lo44ryGcQRxqH2+iMukh5XwaQYjVnoawgaZd2wkXadFcdn4SCs7z
dKQUB76yt/1I7AMKmv4vs4q4uzT4EX7ruk2/fizfXMhabmsKxG0q/Nz0O44Xsi1I
JjyPOKbRQhkKJA7WpgYoLvwH9QKBgQC68quYOChRL8NZeNun7cSWJd4inInupIGj
AzrZe4yoyC7AsiUZXs/sNngvNvCzDlEIyL8vc9OFLSC8NdhwwlvVIaXW6tGTqifn
AghqW/9rYq2Vk3+tcdVeeXpFF1JuMH4+p/Nj9gvzduC9BEXl623IW3JREzRBNobR
Aez7FSuXjwKBgFEJIDxTVxCYdNSfnMLCKxDTPj+vlfC4SmhphSX1GV2nxSSX9oDl
xUSiSFzKlkSOHH9pf+kmWmq4HSei9tlVDBkcQGaLNs5xNieyUWLJEvDH5/kCBexI
DsMMe4T8IYAduWOGPuAlCQEBrdvz4eftvek2dKxLwHfae+eFdulLUAPlAoGBAJLJ
a9ZvcaideiNMdBwc4xiJzysaAmtwm6FlLdYJ3l3AIIWI2vxap6Nu+VsJJmFRQmtF
RGh7539P+b4OAU44LWbhrpdbdQceuYn23Ki2Z4znxCgH0l5bXQ97DnglYcHHLbkA
omjAjo2xr7B6JG/tNRv2QYJLM+Job964Rly3OK0dAoGBAKg0VKBd/hU22eBx+0TC
ZzjJ60f1TCw6xwILtSvxS//RjYDy2ixV85u8Q/3rQNwGrfzLMb5x9ZoOgtBx01JN
isepcZJ4vd/ov6rQ0rqt3Y4ZiUeGP0QJCS4OkV9OdUxO90ZZ5lfPhRPYNAcTE70Z
jVNzcP9ZsGeGA6Tiktk/+Cit
-----END RSA PRIVATE KEY-----`,
	// 你的支付宝公钥
	AlipayPublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkqrv74Mlgh0f6cZNxyqa
m6HWraKW5OKauIfTN4jtY3oqysw2ESsxAjV38y5eDGQtf2Q2ZgKprxbZx1JowzrQ
YquCdF2hU5yHfM12Q2Tqal2W3DsmdTBQ9r8WdKLjC3L+yO2fuP2IwcHsX/cesz+k
CrvDvOPB2UNm6PI824QfFz7FeZaDv5JDCRb8T4qXrVjLCCFdVJiWeyBevJ+qdVSA
uR04BWx/ESEkn78zLsTw3YXxLJMDQo07IZNIAjd8vj+2nxC1sjOARcDGdlj60Dqy
Ja+jHlt93/QfXcGOoBSB7GYkB3oRi50g24YnL0Uq9o9V8WKjkGooIgn9ifLiLoD
JbwIDAQAB
-----END PUBLIC KEY-----`,
	IsProduction: false, // 沙箱环境
	NotifyURL:    "http://localhost:8888/v1/payment/notify",
	ReturnURL:    "http://localhost:8888/v1/payment/return",
}

// 初始化支付宝客户端
func InitAlipayClient(config AlipayConfig) error {
	var err error
	alipayClient, err = alipay.New(config.AppID, config.PrivateKey, config.IsProduction)
	if err != nil {
		return fmt.Errorf("初始化支付宝客户端失败: %v", err)
	}

	// 加载支付宝公钥
	err = alipayClient.LoadAliPayPublicKey(config.AlipayPublicKey)
	if err != nil {
		return fmt.Errorf("加载支付宝公钥失败: %v", err)
	}

	log.Printf("支付宝客户端初始化成功，环境: %s", map[bool]string{true: "生产", false: "沙箱"}[config.IsProduction])
	return nil
}

// 创建支付订单
func CreateAlipayOrder(orderID, subject, totalAmount, userID string) (string, error) {
	if alipayClient == nil {
		// 如果客户端未初始化，使用默认配置初始化
		if err := InitAlipayClient(defaultConfig); err != nil {
			return "", err
		}
	}

	// 创建支付参数
	var p = alipay.TradePagePay{}
	p.NotifyURL = defaultConfig.NotifyURL
	p.ReturnURL = defaultConfig.ReturnURL
	p.Subject = subject
	p.OutTradeNo = orderID
	p.TotalAmount = totalAmount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	// 可选参数
	p.PassbackParams = userID // 回传参数，可以传递用户ID

	// 生成支付URL
	url, err := alipayClient.TradePagePay(p)
	if err != nil {
		return "", fmt.Errorf("创建支付订单失败: %v", err)
	}

	log.Printf("创建支付订单成功: orderID=%s, amount=%s, url=%s", orderID, totalAmount, url.String())
	return url.String(), nil
}

// 验证支付回调
func VerifyAlipayCallback(params map[string]string) (bool, error) {
	if alipayClient == nil {
		return false, fmt.Errorf("支付宝客户端未初始化")
	}

	// 转换参数格式
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	// 验证签名
	err := alipayClient.VerifySign(values)
	if err != nil {
		return false, fmt.Errorf("验证签名失败: %v", err)
	}

	return true, nil
}

// 查询支付状态
func QueryAlipayOrder(orderID string) (*alipay.TradeQueryRsp, error) {
	if alipayClient == nil {
		return nil, fmt.Errorf("支付宝客户端未初始化")
	}

	var p = alipay.TradeQuery{}
	p.OutTradeNo = orderID

	rsp, err := alipayClient.TradeQuery(context.Background(), p)
	if err != nil {
		return nil, fmt.Errorf("查询订单失败: %v", err)
	}

	return rsp, nil
}

// 支付结果结构
type PaymentResult struct {
	Success     bool   `json:"success"`
	OrderID     string `json:"order_id"`
	TradeNo     string `json:"trade_no"`
	TotalAmount string `json:"total_amount"`
	TradeStatus string `json:"trade_status"`
	Message     string `json:"message"`
}

// 解析支付回调结果
func ParseAlipayCallback(params map[string]string) *PaymentResult {
	result := &PaymentResult{
		OrderID:     params["out_trade_no"],
		TradeNo:     params["trade_no"],
		TotalAmount: params["total_amount"],
		TradeStatus: params["trade_status"],
	}

	// 判断支付是否成功
	if params["trade_status"] == "TRADE_SUCCESS" || params["trade_status"] == "TRADE_FINISHED" {
		result.Success = true
		result.Message = "支付成功"
	} else {
		result.Success = false
		result.Message = "支付失败或未完成"
	}

	return result
}
