package v1

import "github.com/gogf/gf/v2/frame/g"

type SseReq struct {
	g.Meta `path:"/sse/create" method:"get"`
}

type SseRes struct{}

type SendMessageReq struct {
	g.Meta    `path:"/sse/send" method:"post"`
	ClientId  string
	EventType string
	Data      string
}

type SendMessageRes struct {
	Ok bool
}

type BroadcastMessageReq struct {
	g.Meta    `path:"/sse/broadcast" method:"post"`
	EventType string
	Data      string
}

type BroadcastMessageRes struct{}
