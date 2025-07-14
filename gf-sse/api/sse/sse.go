// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sse

import (
	"context"

	"gf-sse/api/sse/v1"
)

type ISseV1 interface {
	Sse(ctx context.Context, req *v1.SseReq) (res *v1.SseRes, err error)
	SendMessage(ctx context.Context, req *v1.SendMessageReq) (res *v1.SendMessageRes, err error)
	BroadcastMessage(ctx context.Context, req *v1.BroadcastMessageReq) (res *v1.BroadcastMessageRes, err error)
}
