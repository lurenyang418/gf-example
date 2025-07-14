package sse

import (
	"context"
	"fmt"
	"time"
)

// SendToClient 发送消息到指定客户端
func (s *sSse) SendToClient(ctx context.Context, clientId, eventType, data string) bool {
	if client := s.clients.Get(clientId); client != nil {
		c := client.(*Client)
		msg := fmt.Sprintf("id: %d\nevent: %s\ndata: %s\n\n", time.Now().UnixNano(), eventType, data)
		select {
		case c.messageChan <- msg:
			return true
		default:
			return false
		}
	}

	return false
}
