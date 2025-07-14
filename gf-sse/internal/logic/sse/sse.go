package sse

import (
	"gf-sse/internal/service"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Client struct {
	Id          string
	Request     *ghttp.Request
	messageChan chan string
}

type sSse struct {
	clients *gmap.StrAnyMap
}

func New() *sSse {
	return &sSse{
		clients: gmap.NewStrAnyMap(true),
	}
}

func init() {
	service.RegisterSse(New())
}
