package data

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"kratos_client/internal/conf"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDb, NewUserRepo, NewCityRepo, NewDoctorsRepo, NewDrugRepo, NewEstimateRepo, NewCartRepo, NewOrderRepo, NewDrugInventoryRepo, NewCouponRepo, NewPaymentRepo, NewPrescriptionRepo)

// Data .
type Data struct {
	Db  *gorm.DB
	RDb *redis.Client
	Es  *elasticsearch.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, DB *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	// Redis连接失败不会导致应用启动失败
	var newRedis *redis.Client
	if c.Redis != nil && c.Redis.Addr != "" {
		var err error
		newRedis, err = NewRedis(c)
		if err != nil {
			log.Errorf("Redis connection failed: %v", err)
			log.Warn("Redis unavailable, some features will be limited")
			newRedis = nil // 设置为nil，允许应用继续运行
		}
	} else {
		log.Info("Redis not configured, running without Redis")
		newRedis = nil
	}

	// Elasticsearch连接失败不会导致应用启动失败
	newEs, _ := NewElasticsearch(c)

	return &Data{
		Db:  DB,
		RDb: newRedis,
		Es:  newEs, // 可能为nil，需要在使用时检查
	}, cleanup, nil
}

func NewRedis(c *conf.Data) (Rdb *redis.Client, err error) {
	//初始化redis
	Rdb = redis.NewClient(&redis.Options{
		Network:      c.Redis.Network,
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	ctx := context.Background()
	if err = Rdb.Ping(ctx).Err(); err != nil {
		log.Errorf("failed to ping redis: %v", err)
		return nil, err
	}

	log.Info("redis connected")
	return Rdb, nil
}

func NewDb(c *conf.Data) (Db *gorm.DB, err error) {
	dsn := c.Database.Source
	fmt.Println("dsn:", dsn)

	// 检查是否使用SQLite
	if strings.Contains(dsn, ".db") || strings.Contains(dsn, "file:") {
		log.Info("Using SQLite database")
		Db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		log.Info("Using MySQL database")
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Errorf("Database connection failed: %v", err)
		return nil, err
	}

	log.Info("Database connected successfully")
	return
}

func NewElasticsearch(c *conf.Data) (*elasticsearch.Client, error) {
	// 如果没有配置Elasticsearch地址，返回nil客户端
	if c.Elasticsearch == nil || len(c.Elasticsearch.Addresses) == 0 {
		log.Warn("elasticsearch not configured, search functionality will be limited")
		return nil, nil
	}

	cfg := elasticsearch.Config{
		Addresses: c.Elasticsearch.Addresses,
		Username:  c.Elasticsearch.Username,
		Password:  c.Elasticsearch.Password,
	}

	if c.Elasticsearch.MaxRetries > 0 {
		cfg.MaxRetries = int(c.Elasticsearch.MaxRetries)
	}

	if c.Elasticsearch.Timeout != nil {
		cfg.Transport = &http.Transport{
			ResponseHeaderTimeout: c.Elasticsearch.Timeout.AsDuration(),
		}
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Errorf("failed to create elasticsearch client: %v", err)
		log.Warn("elasticsearch unavailable, search will fallback to database")
		return nil, nil // 返回nil而不是错误，允许应用继续运行
	}

	// 测试连接
	res, err := client.Info()
	if err != nil {
		log.Errorf("failed to get elasticsearch info: %v", err)
		log.Warn("elasticsearch unavailable, search will fallback to database")
		return nil, nil // 返回nil而不是错误，允许应用继续运行
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Errorf("elasticsearch connection error: %s", res.String())
		log.Warn("elasticsearch unavailable, search will fallback to database")
		return nil, nil // 返回nil而不是错误，允许应用继续运行
	}

	log.Info("elasticsearch connected successfully")
	return client, nil
}

func (d *Data) Redis() *redis.Client {
	return d.RDb
}

func (d *Data) Elasticsearch() *elasticsearch.Client {
	return d.Es
}
