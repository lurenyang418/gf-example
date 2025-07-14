package sse

import (
	"context"
	"time"
)

// Broadcast 向所有客户端广播消息
func (s *sSse) Broadcast(ctx context.Context, eventType, data string) int {
	count := 0
	s.clients.Iterator(func(k string, v interface{}) bool {
		if s.SendToClient(ctx, k, eventType, data) {
			count++
		}
		return true
	})

	return count
}

func (s *sSse) heartbeatSender() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for range ticker.C {
		s.clients.Iterator(func(k string, v interface{}) bool {
			c := v.(*Client)
			select {
			case c.messageChan <- ": heartbeat\n\n":
			default:
			}
			return true
		})
	}
}
