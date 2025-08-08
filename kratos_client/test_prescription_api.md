# 处方 API 测试文档

## 1. 获取处方列表

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/prescriptions?status=已开具&prescription_type=西药&start_date=2025-01-01&end_date=2025-12-31&page=1&page_size=10"
```

### 响应
```json
{
  "prescriptions": [
    {
      "id": 1,
      "prescription_no": "RX20250107001",
      "doctor_id": 1001,
      "patient_id": 2001,
      "medical_record_id": 3001,
      "prescription_date": "2025-01-07",
      "total_amount": "156.80",
      "prescription_type": "西药",
      "usage_instruction": "按医嘱服用，注意饭后服药",
      "status": "已开具",
      "auditor_id": 1002,
      "audit_time": "2025-01-07T10:30:00Z",
      "audit_notes": "处方审核通过",
      "created_at": "2025-01-07T09:15:00Z",
      "updated_at": "2025-01-07T10:30:00Z",
      "doctor_name": "李医生",
      "patient_name": "张三",
      "auditor_name": "王药师",
      "medicine_count": 3
    }
  ],
  "total": 1
}
```

## 2. 获取患者处方列表

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/patients/2001/prescriptions?status=已开具&page=1&page_size=10"
```

### 响应
```json
{
  "prescriptions": [
    {
      "id": 1,
      "prescription_no": "RX20250107001",
      "doctor_id": 1001,
      "patient_id": 2001,
      "prescription_date": "2025-01-07",
      "total_amount": "156.80",
      "prescription_type": "西药",
      "status": "已开具",
      "doctor_name": "李医生",
      "patient_name": "张三",
      "medicine_count": 3,
      "created_at": "2025-01-07T09:15:00Z"
    }
  ],
  "total": 1
}
```

## 3. 获取医生处方列表

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/doctors/1001/prescriptions?status=已开具&page=1&page_size=10"
```

### 响应
```json
{
  "prescriptions": [
    {
      "id": 1,
      "prescription_no": "RX20250107001",
      "doctor_id": 1001,
      "patient_id": 2001,
      "prescription_date": "2025-01-07",
      "total_amount": "156.80",
      "prescription_type": "西药",
      "status": "已开具",
      "doctor_name": "李医生",
      "patient_name": "张三",
      "medicine_count": 3,
      "created_at": "2025-01-07T09:15:00Z"
    }
  ],
  "total": 1
}
```

## 4. 获取处方详情

### 请求
```bash
curl -X GET "http://localhost:8000/api/v1/prescriptions/1"
```

### 响应
```json
{
  "prescription": {
    "id": 1,
    "prescription_no": "RX20250107001",
    "doctor_id": 1001,
    "patient_id": 2001,
    "medical_record_id": 3001,
    "prescription_date": "2025-01-07",
    "total_amount": "156.80",
    "prescription_type": "西药",
    "usage_instruction": "按医嘱服用，注意饭后服药",
    "status": "已开具",
    "auditor_id": 1002,
    "audit_time": "2025-01-07T10:30:00Z",
    "audit_notes": "处方审核通过",
    "created_at": "2025-01-07T09:15:00Z",
    "updated_at": "2025-01-07T10:30:00Z",
    "doctor_name": "李医生",
    "patient_name": "张三",
    "auditor_name": "王药师",
    "medicine_count": 3
  },
  "medicines": [
    {
      "id": 1,
      "prescription_id": 1,
      "medicine_id": 4001,
      "quantity": "2.00",
      "unit": "盒",
      "unit_price": "28.50",
      "total_price": "57.00",
      "dosage": "1片",
      "frequency": "每日3次",
      "duration": "7天",
      "usage_method": "口服",
      "notes": "饭后服用",
      "created_at": "2025-01-07T09:15:00Z",
      "medicine_name": "阿莫西林胶囊",
      "medicine_spec": "0.25g*24粒",
      "manufacturer": "华北制药"
    },
    {
      "id": 2,
      "prescription_id": 1,
      "medicine_id": 4002,
      "quantity": "1.00",
      "unit": "瓶",
      "unit_price": "45.60",
      "total_price": "45.60",
      "dosage": "10ml",
      "frequency": "每日2次",
      "duration": "5天",
      "usage_method": "口服",
      "notes": "摇匀后服用",
      "created_at": "2025-01-07T09:15:00Z",
      "medicine_name": "止咳糖浆",
      "medicine_spec": "100ml/瓶",
      "manufacturer": "同仁堂"
    },
    {
      "id": 3,
      "prescription_id": 1,
      "medicine_id": 4003,
      "quantity": "1.00",
      "unit": "盒",
      "unit_price": "54.20",
      "total_price": "54.20",
      "dosage": "2片",
      "frequency": "每日1次",
      "duration": "10天",
      "usage_method": "口服",
      "notes": "睡前服用",
      "created_at": "2025-01-07T09:15:00Z",
      "medicine_name": "维生素B复合片",
      "medicine_spec": "30片/盒",
      "manufacturer": "拜耳"
    }
  ]
}
```

## 处方状态说明

- `已取消`: 处方已被取消
- `已开具`: 医生已开具处方，等待审核
- `已审核`: 药师已审核通过，可以发药
- `已发药`: 药品已发放给患者

## 处方类型说明

- `西药`: 西医药品处方
- `中药`: 中医药品处方  
- `中西药`: 中西医结合处方

## 查询参数说明

### 处方列表查询参数
- `status`: 处方状态过滤
- `prescription_type`: 处方类型过滤
- `start_date`: 开始日期 (格式: YYYY-MM-DD)
- `end_date`: 结束日期 (格式: YYYY-MM-DD)
- `page`: 页码 (默认: 1)
- `page_size`: 每页数量 (默认: 10)

### 患者/医生处方列表查询参数
- `status`: 处方状态过滤
- `page`: 页码 (默认: 1)
- `page_size`: 每页数量 (默认: 10)

## 处方药品明细字段说明

- `quantity`: 药品数量
- `unit`: 药品单位 (盒、瓶、袋等)
- `unit_price`: 单价
- `total_price`: 总价
- `dosage`: 单次用量
- `frequency`: 用药频次
- `duration`: 疗程时长
- `usage_method`: 用药方法 (口服、外用等)
- `notes`: 特殊说明

## 扩展信息

处方列表和详情中包含以下扩展信息：
- `doctor_name`: 开方医生姓名
- `patient_name`: 患者姓名
- `auditor_name`: 审核药师姓名
- `medicine_count`: 处方中药品种类数量
- `medicine_name`: 药品名称
- `medicine_spec`: 药品规格
- `manufacturer`: 生产厂家