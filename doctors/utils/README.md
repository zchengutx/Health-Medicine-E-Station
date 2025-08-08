# 加密工具包

这是一个基于Go语言实现的加密工具包，提供了AES对称加密和密码哈希等安全功能。

## 包含工具

1. **AES加密工具** - 提供AES对称加密和解密功能
2. **哈希加密工具** - 提供密码哈希和各种哈希算法

## 功能特性

### AES加密功能
- 支持AES-128、AES-192、AES-256加密
- 使用GCM模式，提供认证加密
- 自动生成随机nonce，确保每次加密结果不同
- Base64编码输出，便于存储和传输
- 提供密钥生成和转换工具

### 哈希加密功能
- 支持bcrypt密码哈希（推荐用于密码存储）
- 支持MD5、SHA1、SHA256、SHA512哈希算法
- 提供加盐哈希功能
- 密码验证功能
- 可调节bcrypt强度

## 主要函数

### AES加密函数

#### AESEncrypt
```go
func AESEncrypt(key []byte, plaintext string) (string, error)
```
使用AES算法加密数据，返回base64编码的密文。

**参数:**
- `key`: 加密密钥，长度必须是16、24或32字节
- `plaintext`: 要加密的明文

**返回:**
- `string`: base64编码的密文
- `error`: 错误信息

#### AESDecrypt
```go
func AESDecrypt(key []byte, ciphertext string) (string, error)
```
使用AES算法解密数据。

**参数:**
- `key`: 解密密钥，必须与加密时使用的密钥相同
- `ciphertext`: base64编码的密文

**返回:**
- `string`: 解密后的明文
- `error`: 错误信息

#### GenerateAESKey
```go
func GenerateAESKey(keySize int) ([]byte, error)
```
生成指定长度的随机AES密钥。

**参数:**
- `keySize`: 密钥长度，支持16、24、32字节

**返回:**
- `[]byte`: 生成的密钥
- `error`: 错误信息

#### StringToAESKey
```go
func StringToAESKey(keyStr string, keySize int) []byte
```

将字符串转换为指定长度的AES密钥。

**参数:**
- `keyStr`: 密钥字符串
- `keySize`: 目标密钥长度

**返回:**
- `[]byte`: 处理后的密钥

### 哈希加密函数

#### HashPassword
```go
func HashPassword(password string, cost int) (string, error)
```

使用bcrypt算法对密码进行哈希加密。

**参数:**
- `password`: 要加密的密码
- `cost`: 加密强度，建议使用bcrypt.DefaultCost(10)或更高

**返回:**
- `string`: 哈希后的密码
- `error`: 错误信息

#### VerifyPassword
```go
func VerifyPassword(password, hash string) (bool, error)
```

验证密码是否与哈希值匹配。

**参数:**
- `password`: 要验证的明文密码
- `hash`: 存储的哈希值

**返回:**
- `bool`: 是否匹配
- `error`: 错误信息

#### SHA256Hash
```go
func SHA256Hash(data string) string
```

计算字符串的SHA256哈希值。

**参数:**
- `data`: 要哈希的数据

**返回:**
- `string`: SHA256哈希值的十六进制字符串

#### HashWithSalt
```go
func HashWithSalt(data, salt string) string
```

使用盐值进行SHA256哈希。

**参数:**
- `data`: 要哈希的数据
- `salt`: 盐值

**返回:**
- `string`: 加盐后的SHA256哈希值

#### GenerateSalt
```go
func GenerateSalt(length int) (string, error)
```

生成指定长度的随机盐值。

**参数:**
- `length`: 盐值长度

**返回:**
- `string`: 随机盐值
- `error`: 错误信息

## 使用示例

### AES加密解密示例

#### 基本加密解密
```go
package main

import (
    "fmt"
    "log"
    "your-project/utils"
)

func main() {
    // 准备密钥和明文
    keyStr := "mySecretPassword123"
    key := utils.StringToAESKey(keyStr, 32) // 生成32字节密钥
    plaintext := "这是需要加密的敏感信息"
    
    // 加密
    ciphertext, err := utils.AESEncrypt(key, plaintext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("密文: %s\n", ciphertext)
    
    // 解密
    decryptedText, err := utils.AESDecrypt(key, ciphertext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("解密: %s\n", decryptedText)
}
```

#### 生成随机密钥
```go
// 生成32字节密钥（AES-256）
key, err := utils.GenerateAESKey(32)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("生成的密钥: %x\n", key)
```

### 哈希加密示例

#### 密码哈希和验证
```go
package main

import (
    "fmt"
    "log"
    "your-project/utils"
)

func main() {
    password := "mySecretPassword123"
    
    // 哈希密码
    hashedPassword, err := utils.QuickHashPassword(password)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("哈希后的密码: %s\n", hashedPassword)
    
    // 验证密码
    isValid, err := utils.VerifyPassword(password, hashedPassword)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("密码验证结果: %t\n", isValid)
}
```

#### 数据哈希
```go
// SHA256哈希
data := "Hello, World!"
hash := utils.SHA256Hash(data)
fmt.Printf("SHA256: %s\n", hash)

// 加盐哈希
salt, _ := utils.GenerateSalt(16)
saltedHash := utils.HashWithSalt(data, salt)
fmt.Printf("加盐哈希: %s\n", saltedHash)
```

#### 不同哈希算法
```go
data := "test data"
fmt.Printf("MD5: %s\n", utils.MD5Hash(data))
fmt.Printf("SHA1: %s\n", utils.SHA1Hash(data))
fmt.Printf("SHA256: %s\n", utils.SHA256Hash(data))
fmt.Printf("SHA512: %s\n", utils.SHA512Hash(data))
```

## 安全注意事项

### AES加密安全
1. **密钥管理**: 请妥善保管加密密钥，不要将密钥硬编码在代码中
2. **密钥长度**: 推荐使用32字节密钥（AES-256）以获得最高安全性
3. **密钥存储**: 在生产环境中，建议使用专门的密钥管理服务
4. **传输安全**: 在网络传输时，请使用HTTPS等安全协议

### 密码哈希安全
1. **使用bcrypt**: 对于密码存储，强烈推荐使用bcrypt而不是简单的哈希算法
2. **适当的强度**: bcrypt强度建议设置为12或更高（生产环境）
3. **避免MD5/SHA1**: 不要使用MD5或SHA1进行密码哈希，仅用于数据完整性校验
4. **盐值使用**: 对于自定义哈希，务必使用随机盐值防止彩虹表攻击
5. **时间安全**: bcrypt具有时间恒定特性，可防止时序攻击

## 测试

编译检查：
```bash
go build ./utils
```

运行演示：
```go
// AES加密演示（需要先创建测试文件）
utils.DemoAESUsage()

// 哈希加密演示
utils.DemoHashUsage()
```

## 技术细节

### AES加密技术
- **加密模式**: GCM (Galois/Counter Mode)
- **认证**: 提供数据完整性验证
- **随机性**: 每次加密使用不同的nonce
- **编码**: 输出使用Base64编码
- **兼容性**: 符合AES标准，可与其他语言实现互操作

### 哈希技术
- **bcrypt**: 基于Blowfish算法的自适应哈希函数
- **盐值**: 自动生成随机盐值，防止彩虹表攻击
- **工作因子**: 可调节的计算复杂度，随硬件发展可提高安全性
- **时间恒定**: bcrypt验证时间恒定，防止时序攻击
- **标准兼容**: 符合各种哈希算法标准，输出十六进制编码

## 许可证

本项目遵循项目根目录的许可证。