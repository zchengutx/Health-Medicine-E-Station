package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos_client/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDb, NewUserRepo)

// Data .
type Data struct {
	Db  *gorm.DB
	RDb *redis.Client
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, DB *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	newRedis, err := NewRedis(c)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		Db:  DB,
		RDb: newRedis,
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
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func (d *Data) Redis() *redis.Client {
	return d.RDb
}
