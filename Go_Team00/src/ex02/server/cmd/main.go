package main

import (
	"net"

	"server/logger"
	"server/ports/grpc_server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ListenAndServe(log *zap.Logger, server *grpc.Server) error {
	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		return err
	}

	log.Info("Server Started!!!")

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func main() {
	log := logger.NewLogger()
	server := grpc_server.NewStatisticServer(log)

	if err := ListenAndServe(log, server); err != nil {
		log.Fatal("Error occurred while serving server",
			zap.Error(err),
		)
	}
}
