package sse

import (
	"context"

	v1 "gf-sse/api/sse/v1"
	"gf-sse/internal/service"
)

func (c *ControllerV1) BroadcastMessage(ctx context.Context, req *v1.BroadcastMessageReq) (res *v1.BroadcastMessageRes, err error) {
	service.Sse().Broadcast(ctx, req.EventType, req.Data)

	return &v1.BroadcastMessageRes{}, nil
}
