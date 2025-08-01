package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/zhaoxlchn/gin-basic-structure/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Data struct {
	Db    *gorm.DB
	Cache *redis.Client
}

var ProviderSetData = wire.NewSet(NewData, NewDb, NewRedis)

// NewData new data instance
func NewData(db *gorm.DB, redis *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.Println("closing the database/redis resources")
		sqlDb, _ := db.DB()
		if err := sqlDb.Close(); err != nil {
			log.Println("closing database error! ", err.Error())
		}
		if err := redis.Close(); err != nil {
			log.Println("closing redis error! ", err.Error())
		}
	}
	return &Data{Db: db, Cache: redis}, cleanup, nil
}

func NewDb(c *conf.Database) (*gorm.DB, error) {
	config := mysql.Config{
		DSN:                       c.Source, // DSN data source name
		DefaultStringSize:         200,      // string 类型字段的默认长度
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(config), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	setDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	setDb.SetMaxIdleConns(c.MaxIdleConn)
	setDb.SetMaxOpenConns(c.MaxOpenConn)
	setDb.SetConnMaxIdleTime(time.Duration(c.ConnMaxLifeTime) * time.Second)
	return db, nil
}

func NewRedis(c *conf.Redis) *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           c.Db,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Millisecond,
	})
	return cache
}
