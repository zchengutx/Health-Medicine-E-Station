# 个人信息页面问题分析和修复方案

## 🔍 问题分析

根据控制台错误信息和代码分析，个人信息页面的问题可能出现在以下几个方面：

### 1. 前端问题

#### 1.1 认证状态问题
- **问题**: `doctorId` 可能未正确初始化
- **位置**: `src/stores/auth.ts` 中的 `loginState.doctorId`
- **症状**: 页面显示"用户信息异常，请重新登录"

#### 1.2 API调用问题
- **问题**: API请求参数格式或认证token问题
- **位置**: `src/api/doctor.ts` 中的 `getProfile` 方法
- **症状**: 404错误或网络错误

### 2. 后端问题

#### 2.1 数据库查询问题
- **问题**: 医生记录不存在或查询条件错误
- **位置**: `doctors/internal/data/doctor.go` 中的 `GetDoctorByID`
- **症状**: 返回"医生不存在"错误

#### 2.2 服务层问题
- **问题**: 业务逻辑处理错误
- **位置**: `doctors/internal/service/doctor.go` 中的 `GetDoctorProfile`
- **症状**: 500内部服务器错误

## 🛠️ 修复方案

### 方案1: 前端修复

#### 1.1 改进认证状态初始化
```typescript
// 确保doctorId正确设置
const initAuth = () => {
  // ... 现有代码
  
  // 确保doctorId从用户信息中正确提取
  if (savedInfo && savedInfo.DId) {
    loginState.value.doctorId = savedInfo.DId
  }
}
```

#### 1.2 添加API请求重试机制
```typescript
// 添加重试逻辑
const fetchProfileWithRetry = async (retries = 3) => {
  for (let i = 0; i < retries; i++) {
    try {
      return await doctorApi.getProfile({ doctor_id: doctorId })
    } catch (error) {
      if (i === retries - 1) throw error
      await new Promise(resolve => setTimeout(resolve, 1000 * (i + 1)))
    }
  }
}
```

### 方案2: 后端修复

#### 2.1 改进错误处理
```go
// 在 GetDoctorProfile 中添加更详细的日志
func (s *DoctorService) GetDoctorProfile(ctx context.Context, req *pb.GetDoctorProfileReq) (*pb.GetDoctorProfileResp, error) {
    s.log.WithContext(ctx).Infof("获取医生信息请求: doctor_id=%d", req.DoctorId)
    
    doctor, err := s.uc.GetDoctorByID(ctx, uint(req.DoctorId))
    if err != nil {
        s.log.WithContext(ctx).Errorf("获取医生信息失败: doctor_id=%d, error=%v", req.DoctorId, err)
        // ... 错误处理
    }
    
    // ... 其余代码
}
```

#### 2.2 验证数据库连接和数据
```go
// 在数据层添加更多调试信息
func (d *DoctorData) GetDoctorByID(ctx context.Context, id uint) (*biz.Doctor, error) {
    d.logger.WithContext(ctx).Infof("查询医生信息: id=%d", id)
    
    var doctorModel model.Doctors
    err := d.data.db.WithContext(ctx).Where("id = ?", id).First(&doctorModel).Error
    
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            d.logger.WithContext(ctx).Warnf("医生记录不存在: id=%d", id)
            return nil, fmt.Errorf("医生不存在")
        }
        d.logger.WithContext(ctx).Errorf("数据库查询失败: id=%d, error=%v", id, err)
        return nil, fmt.Errorf("查询医生失败: %w", err)
    }
    
    d.logger.WithContext(ctx).Infof("成功查询到医生信息: id=%d, name=%s", doctorModel.Id, doctorModel.Name)
    
    doctor := d.modelToEntity(&doctorModel)
    return doctor, nil
}
```

## 🧪 诊断工具

我已经创建了以下诊断工具来帮助定位问题：

### 1. 前端诊断工具
- **文件**: `src/utils/profileDiagnostic.ts`
- **功能**: 检查认证状态、本地存储、API连接
- **使用**: 在个人信息页面点击"运行诊断"按钮

### 2. API测试工具
- **文件**: `src/utils/apiTest.ts`
- **功能**: 直接测试后端API接口
- **使用**: 自动在诊断过程中运行

## 📋 排查步骤

### 步骤1: 运行前端诊断
1. 打开个人信息页面
2. 如果页面显示错误，点击"运行诊断"按钮
3. 查看浏览器控制台的诊断报告

### 步骤2: 检查后端日志
1. 查看后端服务日志
2. 搜索相关的错误信息
3. 确认数据库连接状态

### 步骤3: 验证数据库数据
1. 直接查询数据库中的医生记录
2. 确认医生ID是否存在
3. 检查数据完整性

### 步骤4: 测试API接口
1. 使用Postman或curl直接测试API
2. 验证请求参数和响应格式
3. 检查认证token是否有效

## 🔧 快速修复

如果问题紧急，可以尝试以下快速修复：

### 前端快速修复
```typescript
// 在 ProfileView.vue 中添加备用获取方式
const fetchProfileFallback = async () => {
  try {
    // 尝试从本地存储获取用户信息
    const userInfo = localStorage.getItem('doctor_info')
    if (userInfo) {
      const parsed = JSON.parse(userInfo)
      if (parsed.DId) {
        Object.assign(form, parsed)
        profileLoaded.value = true
        return
      }
    }
  } catch (error) {
    console.error('备用获取失败:', error)
  }
}
```

### 后端快速修复
```go
// 在服务层添加参数验证
func (s *DoctorService) GetDoctorProfile(ctx context.Context, req *pb.GetDoctorProfileReq) (*pb.GetDoctorProfileResp, error) {
    if req.DoctorId <= 0 {
        return &pb.GetDoctorProfileResp{
            Message: "无效的医生ID",
            Code:    400,
        }, nil
    }
    
    // ... 其余代码
}
```

## 📊 监控建议

为了防止类似问题再次发生，建议添加以下监控：

1. **前端错误监控**: 使用Sentry等工具监控前端错误
2. **API响应时间监控**: 监控API接口的响应时间和成功率
3. **数据库查询监控**: 监控数据库查询的性能和错误率
4. **用户行为监控**: 跟踪用户在个人信息页面的操作流程

---

**总结**: 通过系统性的诊断和修复，我们可以快速定位并解决个人信息页面的问题。建议先运行诊断工具确定问题的具体位置，然后采用相应的修复方案。