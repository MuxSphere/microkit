package main

import (
	"context"
	"fmt"
	"net"

	"github.com/MuxSphere/microkit/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type grpcServer struct {
	proto.UnimplementedGreeterServiceServer
	logger *zap.Logger
}

func (s *grpcServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	s.logger.Info("Received gRPC request", zap.String("name", req.Name))
	return &proto.HelloReply{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

func startGRPCServer(logger *zap.Logger, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServiceServer(s, &grpcServer{logger: logger})

	logger.Info("Starting gRPC server", zap.String("port", port))
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
