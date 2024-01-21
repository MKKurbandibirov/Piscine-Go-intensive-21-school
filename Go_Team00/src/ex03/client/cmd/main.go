package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	adapters "client/internal/adapters/grpc"
	"client/internal/usecase"
	"client/internal/usecase/anomaly_detect"
	"client/internal/usecase/predict"
	"server/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	log := logger.NewLogger()
	client := adapters.NewClient(log)
	err := client.Connect("0.0.0.0:7777")
	if err != nil {
		log.Fatal("Error occurred while connection",
			zap.Error(err),
		)
	}

	id, err := client.Hello(ctx)
	if err != nil {
		log.Fatal("Error occurred while greeting",
			zap.Error(err),
		)
	}

	if err := client.GetStatistics(ctx, id); err != nil {
		log.Fatal("Error occurred while getting statistics",
			zap.Error(err),
		)
	}

	uc := GetUseCase(log, client)
	go uc.Detect(uc.Predict())

	GracefulShutdown(log, cancel, signals, client)
}

func GetUseCase(log *zap.Logger, client *adapters.Client) *usecase.UseCase {
	predictor := predict.NewPredictor(log, client, 100)
	detector := anomaly_detect.NewDetector(log, client, 3)

	return usecase.NewUseCase(predictor, detector)
}

func GracefulShutdown(
	log *zap.Logger,
	cancel context.CancelFunc,
	signals chan os.Signal,
	client *adapters.Client,
) {
	<-signals
	fmt.Println("<<< Client Shutdowned! >>>")

	cancel()

	if err := client.Disconnect(); err != nil {
		log.Fatal("Error occurred while disconnecting client",
			zap.Error(err),
		)
	}
}
