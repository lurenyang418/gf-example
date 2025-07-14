package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"gf-sse/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

var aiResponses = `
GoFrame是一款模块化、高性能、企业级的Go基础开发框架。
SSE（Server-Sent Events）是一种允许服务器向客户端推送数据的技术。
与WebSocket不同，SSE是单向通信，只能服务器向客户端发送数据。
使用SSE可以实现AI模型的流式输出，提升用户体验。
`

type StreamChatRequest struct {
	g.Meta   `path:"/ai/chat" method:"post"`
	Prompt   string `v:"required#请输入您的问题" json:"prompt"`
	ClientId string `v:"required#请输入您的客户端ID" json:"clientId"`
}

type StreamChatResponse struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func (c *Controller) StreamChat(ctx context.Context, req *StreamChatRequest) (res *StreamChatResponse, err error) {
	if req.ClientId == "" {
		req.ClientId = guid.S()
	}

	go func() {
		words := gstr.SplitAndTrim(aiResponses, "")

		time.Sleep(time.Millisecond * 500)

		for i, char := range words {
			fmt.Println(char)
			resp := StreamChatResponse{
				Content: char,
				Done:    i == len(words)-1,
			}
			d, err := json.Marshal(resp)
			if err != nil {
				fmt.Println(err)
				return
			}

			service.Sse().SendToClient(ctx, req.ClientId, "ai_response", string(d))
		}
	}()

	return &StreamChatResponse{Content: "请求已接收，响应将通过 SSE 流式返回", Done: true}, nil
}
