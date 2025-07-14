package sse

import (
	"context"

	v1 "gf-sse/api/sse/v1"
	"gf-sse/internal/service"
)

func (c *ControllerV1) SendMessage(ctx context.Context, req *v1.SendMessageReq) (res *v1.SendMessageRes, err error) {
	ok := service.Sse().SendToClient(ctx, req.ClientId, req.EventType, req.Data)

	return &v1.SendMessageRes{Ok: ok}, nil
}
