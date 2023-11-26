package main

import (
	"client/internal/adapters"
	"client/internal/usecase"
	"client/internal/usecase/anomaly_detect"
	"client/internal/usecase/predict"
	"context"
	"go.uber.org/zap"
	"server/logger"
)

func GetUseCase(log *zap.Logger, client *adapters.Client) *usecase.UseCase {
	predictor := predict.NewPredictor(log, client, 100)
	detector := anomaly_detect.NewDetector(log, client, 3)

	return usecase.NewUseCase(predictor, detector)
}

func main() {
	log := logger.NewLogger()
	client := adapters.NewClient(log)
	err := client.Connect("localhost:4000")
	if err != nil {
		log.Fatal("Error occurred while connection",
			zap.Error(err),
		)
	}
	defer client.Disconnect()

	if err := client.GetStatistics(context.Background()); err != nil {
		log.Fatal("Error occurred while getting statistics",
			zap.Error(err),
		)
	}

	uc := GetUseCase(log, client)
	uc.Detect(uc.Predict())
}
