package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"client/config"
	adapters "client/internal/adapters/grpc"
	"client/internal/adapters/postgres"
	"client/internal/usecase"
	"client/internal/usecase/anomaly_detect"
	"client/internal/usecase/predict"
	"server/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	log := logger.NewLogger()

	cfg, err := config.DefaultConfigParser()
	if err != nil {
		log.Fatal("Error occurred while reading config",
			zap.Error(err),
		)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	dbComm, err := postgres.NewCommunicator(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
	))
	if err != nil {
		log.Fatal("Error occurred while connecting to DB",
			zap.Error(err),
		)
	}

	if err := dbComm.Ping(); err != nil {
		log.Fatal("Error occurred while ping DB",
			zap.Error(err),
		)
	}

	client := adapters.NewClient(log)
	if err := client.Connect(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
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

	uc := GetUseCase(log, cfg, client, dbComm)
	go uc.Detect(uc.Predict())

	GracefulShutdown(log, cancel, signals, client)
}



func GetUseCase(log *zap.Logger, cfg *config.Config, client *adapters.Client, db *postgres.Communicator) *usecase.UseCase {
	predictor := predict.NewPredictor(log, client, cfg.Max)
	detector := anomaly_detect.NewDetector(log, client, db, float64(cfg.K))

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
