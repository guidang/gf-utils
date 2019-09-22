package lredis

import (
	"github.com/go-redis/redis"
	"github.com/gogf/gf/frame/g"
	"time"
)

var (
	Redis      *redis.Client
	expiration time.Duration
)

// Init 初始化 Redis
func Init() error {
	conf := g.Config()
	redisDb := conf.GetInt("cache.redis.db")

	redisHost := conf.GetString("cache.redis.host")
	redisPort := conf.GetString("cache.redis.port")
	redisPassword := conf.GetString("cache.redis.pass")

	Redis = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDb,
	})

	exp := conf.GetInt64("cache.redis.expiration")
	expiration = time.Duration(exp) * time.Second

	return Ping()
}

// Get 获取 Redis
func Get() *redis.Client {
	return Redis
}

// Ping Ping
func Ping() error {
	_, err := Redis.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// SetExpiration SetExpiration
func SetExpiration(redisExpTime int64) time.Duration {
	expiration = time.Duration(redisExpTime) * time.Second
	return expiration
}

// GetExpiration GetExpiration
func GetExpiration() time.Duration {
	return expiration
}
