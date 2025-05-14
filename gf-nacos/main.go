package main

import (
	"time"

	"github.com/gogf/gf/contrib/config/nacos/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	ctx := gctx.GetInitCtx()
	serverConfig := constant.ServerConfig{
		IpAddr: "localhost",
		Port:   8848,
	}
	clientConfig := constant.ClientConfig{
		LogDir:   "tmp/",
		CacheDir: "tmp/",
		// NamespaceId: "public",
		LogLevel:            "debug",
		NotLoadCacheAtStart: true,
		Username:            "admin",
		Password:            "admin",
	}
	configParam := vo.ConfigParam{
		DataId: "config.yaml",
		Group:  "user",
	}

	adapter, _ := nacos.New(ctx, nacos.Config{
		ServerConfigs: []constant.ServerConfig{serverConfig},
		ClientConfig:  clientConfig,
		ConfigParam:   configParam,
		Watch:         true,
	})

	g.Cfg().SetAdapter(adapter)
	if !adapter.Available(ctx) {
		panic("adapter not available")
	}

	for {
		g.Dump(g.Cfg().Data(ctx))
		g.Dump(g.Cfg().MustGet(ctx, "server.address"))
		time.Sleep(time.Second * 3)
	}

}
