package main

import (
	"context"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/alicebob/miniredis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	cacheType := g.Cfg().MustGet(context.TODO(), "server.cacheType").String()
	cache := gcache.New()
	switch cacheType {
	case "redis":
		cache.SetAdapter(gcache.NewAdapterRedis(g.Redis()))
	case "miniredis":
		addr := g.Cfg().MustGet(context.TODO(), "redis.default.address").String()
		pass := g.Cfg().MustGet(context.TODO(), "redis.default.pass").String()
		redis := miniredis.NewMiniRedis()
		// 设置redis的认证密码
		redis.RequireAuth(pass)
		if err := redis.StartAddr(addr); err != nil {
			glog.Fatal(context.TODO(), "failed to start mini redis", err)
		}
		cache.SetAdapter(gcache.NewAdapterRedis(g.Redis()))
	case "memory":
	}

	_ = cache.Set(context.TODO(), "name", "luna", time.Minute*5)
	v, _ := cache.Get(context.TODO(), "name")
	println(v.String())

	select {}
}
