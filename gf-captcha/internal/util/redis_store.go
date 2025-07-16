package util

import (
	"context"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
)

type RedisStore struct {
	Client    gcache.Adapter
	KeyPrefix string
}

// Set sets the digits for the captcha id.
func (r *RedisStore) Set(id string, value string) error {
	expireTime := time.Second * time.Duration(g.Cfg().MustGet(context.TODO(), "captcha.expireTime", 5).Int())
	return r.Client.Set(context.TODO(), r.KeyPrefix+id, value, expireTime)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (r *RedisStore) Get(id string, clear bool) string {
	val, err := r.Client.Get(context.TODO(), r.KeyPrefix+id)
	if err != nil {
		return ""
	}
	if clear {
		r.Client.Remove(context.TODO(), r.KeyPrefix+id)
	}
	return val.String()
}

// Verify captcha's answer directly
func (r *RedisStore) Verify(id string, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}

func NewRedisStore(keyPrefix string) *RedisStore {
	return &RedisStore{
		Client:    gcache.NewAdapterRedis(g.Redis()),
		KeyPrefix: keyPrefix,
	}
}

type MiniRedisStore struct {
	Client    gcache.Adapter
	KeyPrefix string
}

// Set sets the digits for the captcha id.
func (m *MiniRedisStore) Set(id string, value string) error {
	expireTime := time.Second * time.Duration(g.Cfg().MustGet(context.TODO(), "captcha.expireTime", 5).Int())
	return m.Client.Set(context.TODO(), m.KeyPrefix+id, value, expireTime)
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (m *MiniRedisStore) Get(id string, clear bool) string {
	val, err := m.Client.Get(context.TODO(), m.KeyPrefix+id)
	if err != nil {
		return ""
	}
	if clear {
		m.Client.Remove(context.TODO(), m.KeyPrefix+id)
	}
	return val.String()
}

// Verify captcha's answer directly
func (m *MiniRedisStore) Verify(id string, answer string, clear bool) bool {
	v := m.Get(id, clear)
	return v == answer
}

func NewMiniRedisStore(keyPrefix string) *RedisStore {
	addr := g.Cfg().MustGet(context.TODO(), "redis.default.address").String()
	pass := g.Cfg().MustGet(context.TODO(), "redis.default.pass").String()
	redis := miniredis.NewMiniRedis()
	// 设置redis的认证密码
	redis.RequireAuth(pass)
	if err := redis.StartAddr(addr); err != nil {
		glog.Fatal(context.TODO(), "failed to start mini redis", err)
	}
	return &RedisStore{
		Client:    gcache.NewAdapterRedis(g.Redis()),
		KeyPrefix: keyPrefix,
	}
}
