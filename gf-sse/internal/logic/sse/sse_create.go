package sse

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/guid"
)

func (s *sSse) Create(r *ghttp.Request) {
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")

	clientId := r.Get("client_id", guid.S()).String()
	client := &Client{
		Id:          clientId,
		Request:     r,
		messageChan: make(chan string, 100),
	}
	// 注册客户端
	s.clients.Set(clientId, client)
	// 客户端断开连接时，移除客户端
	defer func() {
		s.clients.Remove(clientId)
		close(client.messageChan)
	}()

	// 发送连接成功消息
	r.Response.Writefln("id: %s", clientId)
	// r.Response.Writefln("event: connected")
	r.Response.Writefln("data: {\"status\": \"connected\", \"client_id\": \"%s\"}\n", clientId)
	r.Response.Flush()

	for {
		select {
		case msg, ok := <-client.messageChan:
			if !ok {
				return
			}
			r.Response.Writefln(msg)
			r.Response.Flush()
		case <-r.Context().Done():
			glog.Infof(context.Background(), "client %s disconnected: ", clientId)
			// 客户端断开连接
			return
		}
	}
}
