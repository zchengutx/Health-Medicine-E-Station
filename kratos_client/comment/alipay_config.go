package comment

// 支付宝沙箱配置说明
// 
// 你需要从支付宝开放平台沙箱获取以下3个配置：
// 
// 1. AppID - 沙箱应用ID
//    获取方式：登录支付宝开放平台 -> 开发者中心 -> 沙箱应用 -> 复制AppID
//    示例：2021000122671234
//
// 2. 应用私钥 - 你生成的RSA私钥
//    获取方式：使用支付宝提供的密钥生成工具生成RSA密钥对
//    上传公钥到支付宝开放平台，保留私钥用于签名
//
// 3. 支付宝公钥 - 支付宝提供的公钥
//    获取方式：在支付宝开放平台沙箱应用中，设置好你的应用公钥后
//    支付宝会生成对应的支付宝公钥供你下载使用

// 配置支付宝参数（请替换为你的实际配置）
func GetAlipayConfig() AlipayConfig {
	return AlipayConfig{
		// TODO: 请替换为你的沙箱AppID
		AppID: "你的沙箱AppID",
		
		// TODO: 请替换为你生成的应用私钥
		PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
你的应用私钥内容
-----END RSA PRIVATE KEY-----`,
		
		// TODO: 请替换为支付宝提供的公钥
		AlipayPublicKey: `-----BEGIN PUBLIC KEY-----
支付宝公钥内容
-----END PUBLIC KEY-----`,
		
		IsProduction: false, // 沙箱环境设为false
		NotifyURL:    "http://localhost:8000/v1/payment/notify",
		ReturnURL:    "http://localhost:8000/v1/payment/return",
	}
}

// 快速测试配置（使用支付宝官方测试密钥）
func GetTestAlipayConfig() AlipayConfig {
	return AlipayConfig{
		// 支付宝官方测试AppID
		AppID: "9021000122671234",
		
		// 支付宝官方测试私钥
		PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4f5wg5l2hKsTeNem/V41fGnJm6gOdrj8ym3rFkEjWT2btpk7
yTMT4qiMseKlOx1Ygkxcls8MABICXzRvK4pBxvvQAJVuE1ML+xbh7zjjHjMpTrC+
8LoEqM4rMwlHYqAyDeHyMtfwxieh6y4hnM9x7+NoIX0fLuktaXE1NpHrRgHQmwvF
pFpVoxXSIUaHQHAHuHUlcCHcHxLtFn4YjHnAHZjWvYykqmhnHvGioDHyobAA5ZlA
SBBm4dTweXjZWCqhqmPl/CLiWLJKyK5QYdpEfcNcFJ6ezNzeNVnVAqlyUNuaacPl
uzMjKJUV6LeCMJN55HcOjmO0Q2oQhAP1050VfwIDAQABAoIBAQDl7sgi9j0Fy0cI
kEiJeqxGm6fUKmpuYVK4xTGGrHiGOjmFXQe1MxMdgPn2plDMjqLDuD/y982ZFhDo
fd2+MGTrQcaKhMsYxiQ2VF3AoAoRqe3T61uEjAkw9E4eDV9HiIJp+Ee5Rz5KVRdq
sBYSjUBKDV5z24kaVgMASE5ITEBfmzTiJU+o/A7IX4FyWRdQw7nNbRDBCn1R2+Hv
aTNjh4EOoltGJukqHtVBaAYRQdcMjwTW5LwHBhRd4F6S2+6p2+5RktagbGlB4CWz
Tt8QZTgNbmHHd+4MAIbOX8b6u5jABfz34+vrdOgO+SoHqEBVBcjhgBqxqWJ1D+JE
Fb/qshLhAoGBAPMwObLPiAJWx+hpJxqsma6aq+Af1DINiWuuuBirce4TtMsBQWrP
5x5Mz8fcr+1IjVk1SjbLiayqHozXzMwe6c4VBcaEKxs5CcbCu6EApBtA3RxCx5Pd
6WwXXuinihI5JGPH0T2+ZpUy0xtavy2wI/JjGhx1FuYUTWhRLHpJAoGBAPFdQcQh
4syGNuH6VYMqeK9+GJ9FtsKrsfqfmI9E6jn8SfEnQqjMxHXGvO+lRfqHqVQub+AZ
RxFwdBjbgHa8Z6glJfVcpLfrUn9eFcSEsXbpWw+5iFn+YqiQnHjLQR5Hl5XrtJcl
+Jg5Z8AnXuBQSObzH7RH48YKo0OVhJEehPiTAoGAE4hhDQAW8g5WT6AwqwbqDrbI
qv7nh1QjbH5pwpBGlEkzFQPiXAyYvhZGo4JRvgjyQSNHggOBGgQ5HKMlirVMoD8g
ggqCpPlk45qXrlOm6lXPx/ENbrqjFwiVGCgaKQx+PxdQiecfxlL/6BqMUZiFDWlg
Qn2vqRICB7keqM9AiPECgYEA7MhbfuEgQ4CE2HA9h6py4g+9v6+Iy3RI3NeWPpZf
4D/TiCxUpkPZpVwfMNhYMjBOb5RFtvMF98dwrxaGd4y7mUu7NkxM+f8C4LfqLYem
26fwlVmvQdWaCYnLiZAHpaiE2BjRiNY8vQDn2g1uEXOBrAmI5/Ni6EyZQd1iYCpT
AoGAYmI73ODjdgEgyCciH6Lrr4HEcRRhQlmhVTjgQWnOHw4BhgLQyXZH+7AqCJiS
3NqIxdAlqDyqNn3UBAlRAV/dhMmXZSMqVSuMlIlO3GQPP7MbFb54OnQ5jbqmzVlh
Pg2fFKxZYdXAQa/VgxtXnaTAN9rADRJ2HdOiKNE3TZcVsUE=
-----END RSA PRIVATE KEY-----`,
		
		// 支付宝官方测试公钥
		AlipayPublicKey: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuWJKrQ6SWvS6niI+4vEV
RiX9hn9VCnN8qrFChR8maIq+6krXMKBrLOuZ4fnDAYaHwIaDOvV5SQxYcGdVzjZC
xmyPWLiMEP4sRx7iakHfU3t6WqkLbzP07kNEHiB0zSUB1c5wjWQxnSR1TVFVkMSV
tMFfn6iFXwUbsGhXGvs1eT5SwoEdUBjyuuFBOJ/SmBmdHksajBiOwrWfGHdPHPGS
4dVAGGGG2v2EOcqWPiMuGfmckIXiHXwa5dWJkFoOp+LHUIrwNh8bDXv4TM/tTzgC
wjlRFcr6PcnwSBSCmp1SSxlDvp1IrGxsbGmPiXiGZjwwEb+c1VCn8b1jb03MbNrH
OQIDAQAB
-----END PUBLIC KEY-----`,
		
		IsProduction: false, // 沙箱环境
		NotifyURL:    "http://localhost:8000/v1/payment/notify",
		ReturnURL:    "http://localhost:8000/v1/payment/return",
	}
}