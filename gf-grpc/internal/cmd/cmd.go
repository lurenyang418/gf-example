package cmd

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf-grpc/internal/controller/device"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				),
			}...)
			s := grpcx.Server.New(c)
			device.Register(s)

			isDev := g.Cfg().MustGet(ctx, "server.mode", "dev").String() == "dev"
			if isDev {
				glog.Info(ctx, "gRPC server is running in development mode, enabling reflection.")
				reflection.Register(s.Server)
			}

			s.Run()
			return nil
		},
	}
)
