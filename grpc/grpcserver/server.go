package grpcserver

import (
	"context"
	"github.com/sillyhatxu/consul-client"
	"github.com/sillyhatxu/message-backend/grpc/message"
	"github.com/sillyhatxu/message-backend/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("initial grpc server")
	server := grpc.NewServer()

	healthServer := health.NewServer()
	healthServer.SetServingStatus(consul.DefaultHealthCheckGRPCServerName, hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(server, healthServer)

	message.RegisterMessageServiceServer(server, &Message{})
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type Message struct{}

func (Message) SendMessage(ctx context.Context, request *message.AttachmentRequest) (*message.Response, error) {
	err := service.SendMessage(request)
	if err != nil {
		logrus.Errorf("send birthday error. request : %#v; response error : %v", request, err)
		return nil, err
	}
	return &message.Response{
		Code:    uint32(codes.OK),
		Message: "Success",
	}, nil
}
