package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type MtDoctors struct {
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:主键索引;primaryKey;" json:"id"`                           // 主键索引
	DoctorCode    string    `gorm:"column:doctor_code;type:varchar(32);comment:医生编码;not null;" json:"doctor_code"`               // 医生编码
	Name          string    `gorm:"column:name;type:varchar(50);comment:医生姓名;not null;" json:"name"`                             // 医生姓名
	Gender        string    `gorm:"column:gender;type:varchar(20);comment:性别：1-男，2-女;" json:"gender"`                            // 性别：1-男，2-女
	BirthDate     time.Time `gorm:"column:birth_date;type:date;comment:出生日期;" json:"birth_date"`                                 // 出生日期
	Phone         string    `gorm:"column:phone;type:varchar(20);comment:手机号码;not null;" json:"phone"`                           // 手机号码
	Email         string    `gorm:"column:email;type:varchar(100);comment:邮箱地址;" json:"email"`                                   // 邮箱地址
	Avatar        string    `gorm:"column:avatar;type:varchar(255);comment:头像URL;" json:"avatar"`                                  // 头像URL
	LicenseNumber string    `gorm:"column:license_number;type:varchar(50);comment:执业医师资格证号;not null;" json:"license_number"` // 执业医师资格证号
	DepartmentId  int32     `gorm:"column:department_id;type:mediumint;comment:科室ID;" json:"department_id"`                        // 科室ID
	HospitalId    int32     `gorm:"column:hospital_id;type:mediumint;comment:医院ID;" json:"hospital_id"`                            // 医院ID
	Title         string    `gorm:"column:title;type:varchar(50);comment:职称;" json:"title"`                                        // 职称
	Speciality    string    `gorm:"column:speciality;type:varchar(200);comment:专业特长;" json:"speciality"`                         // 专业特长
	PracticeScope string    `gorm:"column:practice_scope;type:varchar(500);comment:执业范围;" json:"practice_scope"`                 // 执业范围
	PasswordHash  string    `gorm:"column:password_hash;type:varchar(255);comment:密码哈希;not null;" json:"password_hash"`          // 密码哈希
	Salt          string    `gorm:"column:salt;type:varchar(32);comment:密码盐值;not null;" json:"salt"`                             // 密码盐值
	Status        string    `gorm:"column:status;type:varchar(20);comment:状态：0-禁用，1-启用，2-待审核;" json:"status"`               // 状态：0-禁用，1-启用，2-待审核
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime(3);comment:最后登录时间;" json:"last_login_time"`            // 最后登录时间
	LastLoginIp   string    `gorm:"column:last_login_ip;type:varchar(45);comment:最后登录IP;" json:"last_login_ip"`                  // 最后登录IP
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime(3);" json:"deleted_at"`
	CreatedBy     uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"` // 创建者
	UpdatedBy     uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"` // 更新者
	DeletedBy     uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"` // 删除者
}

func (m *MtDoctors) TableName() string {
	return "mt_doctors"
}

type DoctorsRepo interface {
	DoctorsFind(ctx context.Context, m *MtDoctors) (*[]MtDoctors, error)
	FindByID(ctx context.Context, id int32) (*MtDoctors, error)
}

type DoctorsService struct {
	repo DoctorsRepo
	log  *log.Helper
}

func NewDoctorsUsecase(repo DoctorsRepo, logger log.Logger) *DoctorsService {
	return &DoctorsService{repo: repo, log: log.NewHelper(logger)}
}

func (m *DoctorsService) DoctorsFind(ctx context.Context, req *MtDoctors) (*[]MtDoctors, error) {
	m.log.WithContext(ctx).Infof("MtCity %+v", req)
	return m.repo.DoctorsFind(ctx, req)
}

func (m *DoctorsService) FindByID(ctx context.Context, id int32) (*MtDoctors, error) {
	m.log.WithContext(ctx).Infof("FindByID doctor %d", id)
	return m.repo.FindByID(ctx, id)
}
