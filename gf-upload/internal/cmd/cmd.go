package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-upload/internal/consts"
	"gf-upload/internal/controller/file"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetServerRoot("upload")
			s.SetIndexFolder(true)
			s.AddStaticPath("/static", "upload")
			isDebug := g.Cfg().MustGet(ctx, "server.mode", false).Bool()
			if isDebug {
				s.BindHandler("/api/swagger", func(r *ghttp.Request) {
					r.Response.Write(consts.SwaggerUIPageContent)
				})
				s.SetOpenApiPath("/api/api.json")
			}

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					file.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)
