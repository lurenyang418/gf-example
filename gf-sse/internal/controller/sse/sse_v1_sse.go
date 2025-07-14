package sse

import (
	"context"

	v1 "gf-sse/api/sse/v1"
	"gf-sse/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Sse(ctx context.Context, req *v1.SseReq) (res *v1.SseRes, err error) {
	service.Sse().Create(g.RequestFromCtx(ctx))

	return
}
