package grpcclient

import (
	"github.com/sillyhatxu/message-backend/grpc/message"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	address = "localhost:8080"
	timeout = 30 * time.Second
)

func TestClient_SendBirthday(t *testing.T) {
	client := NewGRPCMessageClient(address, Timeout(timeout))
	err := client.SendMessage(&message.AttachmentRequest{
		Fallback:   "Happy Birthday : Shikuan Xu",
		Text:       "生日快乐 阴历 : 五月十四",
		AuthorName: "徐士宽",
	})
	assert.Nil(t, err)
}
