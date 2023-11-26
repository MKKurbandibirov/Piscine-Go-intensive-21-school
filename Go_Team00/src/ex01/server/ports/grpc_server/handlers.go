package grpc_server

import (
	"go.uber.org/zap"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"gonum.org/v1/gonum/stat/distuv"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "messages/generated"
)

var (
	sendInterval = 50 * time.Millisecond
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (s *StatisticServer) GetStatistics(in *empty.Empty, stream pb.Transmitter_GetStatisticsServer) error {
	_ = in

	var (
		mean = generateMean()
		svd  = generateSVD()
		id   = generateRandomUUID()
	)

	s.log.Debug("Randomly generated statistic values on a new connection",
		zap.String("Session_ID", id),
		zap.Float64("Mean", mean),
		zap.Float64("SVD", svd),
	)

	dist := distuv.Normal{
		Mu:    mean,
		Sigma: svd,
	}

	for {
		msg := &pb.StatisticValue{
			SessionId:   id,
			Frequency:   dist.Rand(),
			CurrentTime: timestamppb.Now(),
		}

		if err := stream.Send(msg); err != nil {
			s.log.Error("Error occurred while trying to end message",
				zap.String("Session_ID", id),
			)

			return err
		}

		time.Sleep(sendInterval)
	}
}

func generateRandomUUID() string {
	id := uuid.New()

	return id.String()
}

func generateMean() float64 {
	return float64((rand.Int() % 20) - 10)
}

func generateSVD() float64 {
	return rand.Float64()*(1.5-0.3) + 0.3
}
