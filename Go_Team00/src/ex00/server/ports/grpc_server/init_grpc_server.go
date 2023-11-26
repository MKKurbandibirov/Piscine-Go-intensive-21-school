package grpc_server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	pb "messages/generated"
)

type StatisticServer struct {
	pb.TransmitterServer

	log *zap.Logger
}

func NewStatisticServer(log *zap.Logger) *grpc.Server {
	statServer := &StatisticServer{
		log: log,
	}

	server := grpc.NewServer()
	pb.RegisterTransmitterServer(server, statServer)

	return server
}
