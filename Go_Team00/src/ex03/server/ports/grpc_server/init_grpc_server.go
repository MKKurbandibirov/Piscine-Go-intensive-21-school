package grpc_server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	pb "messages/generated"
)

type StatisticData struct {
	Mean float64
	SVD  float64
}

type StatisticServer struct {
	pb.TransmitterServer

	log         *zap.Logger
	clientsInfo map[string]StatisticData
}

func NewStatisticServer(log *zap.Logger) *grpc.Server {
	statServer := &StatisticServer{
		log:         log,
		clientsInfo: make(map[string]StatisticData),
	}

	server := grpc.NewServer()
	pb.RegisterTransmitterServer(server, statServer)

	return server
}
