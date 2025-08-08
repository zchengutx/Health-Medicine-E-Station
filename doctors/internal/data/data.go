package data

import (
	"context"
	"doctors/internal/conf"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDoctorData, NewPatientData, NewConsultationData)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, l log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(l)

	// 初始化MySQL连接
	db, err := NewDB(c, l)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 初始化Redis连接
	rdb, err := NewRedis(c, l)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	d := &Data{
		db:  db,
		rdb: rdb,
	}

	cleanup := func() {
		helper.Info("closing the data resources")

		// 关闭MySQL连接
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}

		// 关闭Redis连接
		rdb.Close()
	}

	return d, cleanup, nil
}

// NewDB 创建数据库连接
func NewDB(c *conf.Data, l log.Logger) (*gorm.DB, error) {
	helper := log.NewHelper(l)

	// 配置GORM日志 - 使用默认日志配置
	gormLogger := logger.Default.LogMode(logger.Silent)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mysql: %w", err)
	}

	// 获取底层sql.DB以配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	helper.Info("database connected successfully")
	return db, nil
}

// NewRedis 创建Redis连接
func NewRedis(c *conf.Data, l log.Logger) (*redis.Client, error) {
	helper := log.NewHelper(l)

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  5 * time.Second,  // 连接超时时间
		PoolTimeout:  10 * time.Second, // 连接池超时时间
		PoolSize:     10,               // 连接池大小
		MinIdleConns: 5,                // 最小空闲连接数
		MaxRetries:   2,                // 最大重试次数
	})

	// 测试Redis连接，但不因为连接失败而终止应用
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		helper.Warnf("Redis连接失败，将使用内存存储作为备用方案: %v", err)
	} else {
		helper.Info("Redis连接成功")
	}

	return rdb, nil
}

// GetDB 获取数据库连接
func (d *Data) GetDB() *gorm.DB {
	return d.db
}

// GetRedis 获取Redis连接
func (d *Data) GetRedis() *redis.Client {
	return d.rdb
}
