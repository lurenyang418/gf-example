// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISse interface {
		// Broadcast 向所有客户端广播消息
		Broadcast(ctx context.Context, eventType string, data string) int
		Create(r *ghttp.Request)
		// SendToClient 发送消息到指定客户端
		SendToClient(ctx context.Context, clientId string, eventType string, data string) bool
	}
)

var (
	localSse ISse
)

func Sse() ISse {
	if localSse == nil {
		panic("implement not found for interface ISse, forgot register?")
	}
	return localSse
}

func RegisterSse(i ISse) {
	localSse = i
}
