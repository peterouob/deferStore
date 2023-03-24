package redis

import (
	"github.com/go-redis/redis"
	"server/component/config"
	"time"
)

var instance *redis.Client

func RedisInit() {
	instance = redis.NewClient(&redis.Options{
		Addr: config.Config.GetString("redis.addr"),
	})
}

func Set(key string, value interface{}, time time.Duration) *redis.StatusCmd {
	return instance.Set(key, value, time)
}

func Get(key string) *redis.StringCmd {
	return instance.Get(key)
}

func Del(keys ...string) *redis.IntCmd {
	return instance.Del(keys...)
}

func MGet(keys ...string) *redis.SliceCmd {
	return instance.MGet(keys...)
}

func SAdd(key string, member ...interface{}) *redis.IntCmd {
	return instance.SAdd(key, member...)
}

func SMember(key string) *redis.StringSliceCmd {
	return instance.SMembers(key)
}
