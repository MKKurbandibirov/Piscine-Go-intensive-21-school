package grpc_server

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
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

func (s *StatisticServer) Greeting(ctx context.Context, in *pb.Empty) (*pb.Hello, error) {
	_ = ctx
	_ = in

	var (
		mean = generateMean()
		svd  = generateSVD()
		id   = generateRandomUUID()
	)

	s.log.Debug("Get new connection",
		zap.String("Session_ID", id),
	)

	s.clientsInfo[id] = StatisticData{
		Mean: mean,
		SVD:  svd,
	}

	return &pb.Hello{SessionId: &pb.SessionID{ID: id}}, nil
}

func (s *StatisticServer) GetStatistics(id *pb.SessionID, stream pb.Transmitter_GetStatisticsServer) error {
	data, ok := s.clientsInfo[id.GetID()]
	if !ok {
		s.log.Error("Couldn't find client with",
			zap.String("SessionID", id.GetID()),
		)
	}

	s.log.Debug("Randomly generated statistic values",
		zap.String("SessionID", id.GetID()),
		zap.Float64("Mean", data.Mean),
		zap.Float64("SVD", data.SVD),
	)

	dist := distuv.Normal{
		Mu:    data.Mean,
		Sigma: data.SVD,
	}

	for {
		msg := &pb.StatisticValue{
			SessionId:   id,
			Frequency:   dist.Rand(),
			CurrentTime: timestamppb.Now(),
		}

		if err := stream.Send(msg); err != nil {
			s.log.Error("Error occurred while trying to end message",
				zap.String("Session_ID", id.ID),
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
