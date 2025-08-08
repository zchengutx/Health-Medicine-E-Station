package data

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// 使用示例：展示如何在repository中使用MySQL和Redis连接

// ExampleUserRepo 示例用户仓库
type ExampleUserRepo struct {
	data *Data
	log  *log.Helper
}

// NewExampleUserRepo 创建用户仓库
func NewExampleUserRepo(data *Data, logger log.Logger) *ExampleUserRepo {
	return &ExampleUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 示例：使用MySQL查询
func (r *ExampleUserRepo) GetUserByID(ctx context.Context, id int64) error {
	// 使用GORM查询数据库
	db := r.data.GetDB()

	// 示例查询（需要根据你的实际表结构调整）
	var user struct {
		ID   int64  `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	err := db.WithContext(ctx).Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		r.log.WithContext(ctx).Errorf("查询用户失败: %v", err)
		return err
	}

	r.log.WithContext(ctx).Infof("查询用户成功: %+v", user)
	return nil
}

// 示例：使用Redis缓存
func (r *ExampleUserRepo) SetUserCache(ctx context.Context, userID int64, data string) error {
	rdb := r.data.GetRedis()

	key := fmt.Sprintf("user:%d", userID)
	err := rdb.Set(ctx, key, data, 30*time.Minute).Err()
	if err != nil {
		r.log.WithContext(ctx).Errorf("设置用户缓存失败: %v", err)
		return err
	}

	r.log.WithContext(ctx).Infof("设置用户缓存成功: key=%s", key)
	return nil
}

// 示例：从Redis获取缓存
func (r *ExampleUserRepo) GetUserCache(ctx context.Context, userID int64) (string, error) {
	rdb := r.data.GetRedis()

	key := fmt.Sprintf("user:%d", userID)
	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		r.log.WithContext(ctx).Errorf("获取用户缓存失败: %v", err)
		return "", err
	}

	r.log.WithContext(ctx).Infof("获取用户缓存成功: key=%s", key)
	return result, nil
}

// 示例：事务操作
func (r *ExampleUserRepo) CreateUserWithTransaction(ctx context.Context, name string) error {
	db := r.data.GetDB()

	// 开始事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 执行数据库操作
	if err := tx.WithContext(ctx).Exec("INSERT INTO users (name) VALUES (?)", name).Error; err != nil {
		tx.Rollback()
		r.log.WithContext(ctx).Errorf("创建用户失败: %v", err)
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		r.log.WithContext(ctx).Errorf("提交事务失败: %v", err)
		return err
	}

	r.log.WithContext(ctx).Infof("创建用户成功: name=%s", name)
	return nil
}
