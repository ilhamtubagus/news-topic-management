package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCacher struct {
	rd *redis.Client
}

func NewRedisCacher(rd *redis.Client) *RedisCacher {
	return &RedisCacher{rd: rd}
}
func (n *RedisCacher) Put(key string, val interface{}, ttl uint64) error {
	//expiration in seconds
	expiration, err := time.ParseDuration(fmt.Sprintf("%ds", ttl))
	if err != nil {
		return err
	}
	return n.rd.Set(context.TODO(), key, val, expiration).Err()
}

func (n *RedisCacher) Get(key string) interface{} {
	r, err := n.rd.Get(context.TODO(), key).Result()
	if err != nil {
		return nil
	}
	return r
}
func (n *RedisCacher) IsExist(key string) bool {
	r, err := n.rd.Exists(context.TODO(), key).Result()
	if err != nil {
		return false
	}
	return r > 0
}

func (n *RedisCacher) Delete(key string) error {
	return n.rd.Del(context.TODO(), key).Err()
}
func (n *RedisCacher) Flush() error {
	return n.rd.FlushAll(context.TODO()).Err()
}
