package grpcclient

import (
	"context"
	"fmt"
	"github.com/sillyhatxu/message-backend/grpc/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

const defaultTimeout = 10 * time.Second

type Client struct {
	address string
	config  *Config
}

func NewGRPCMessageClient(address string, opts ...Option) *Client {
	//default
	config := &Config{
		timeout: defaultTimeout,
	}
	for _, opt := range opts {
		opt(config)
	}
	return &Client{
		address: address,
		config:  config,
	}
}

func (c Client) SendMessage(request *message.AttachmentRequest) error {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := message.NewMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	res, err := client.SendMessage(ctx, request)
	if err != nil {
		return err
	}
	if res.Code != uint32(codes.OK) {
		return fmt.Errorf("response error. %#v", res)
	}
	return nil
}
