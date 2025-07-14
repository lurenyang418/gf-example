package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-sse/internal/controller/ai"
	"gf-sse/internal/controller/sse"
)

func SseHandler(r *ghttp.Request) {
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")

	for i := 0; i < 10; i++ {
		r.Response.Writefln("data: %d", i)
		r.Response.Writefln("id: %d", i)
		r.Response.Writefln("event: message\n")
		r.Response.Writefln("data: {\"message\": \"Hello SSE %d\"}\n", i)
		r.Response.Flush()

		time.Sleep(time.Second)
	}
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(sse.NewV1())
				group.Bind(ai.New())
			})
			s.Run()
			return nil
		},
	}
)
