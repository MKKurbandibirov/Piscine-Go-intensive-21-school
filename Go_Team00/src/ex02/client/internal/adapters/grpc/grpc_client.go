package adapters

import (
	"context"
	"io"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "messages/generated"
)

type Client struct {
	log    *zap.Logger
	values chan float64

	conn   *grpc.ClientConn
	client pb.TransmitterClient
}

func NewClient(log *zap.Logger) *Client {
	return &Client{
		log:    log,
		values: make(chan float64),
	}
}

func (c *Client) Connect(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.conn = conn
	c.client = pb.NewTransmitterClient(c.conn)

	c.log.Debug("Client created!")

	return nil
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) Hello(ctx context.Context) (*pb.SessionID, error) {
	msg, err := c.client.Greeting(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	c.log.Debug("Connected to server",
		zap.String("SessionID", msg.SessionId.GetID()),
	)

	return msg.SessionId, nil
}

func (c *Client) GetStatistics(ctx context.Context, id *pb.SessionID) error {
	stream, err := c.client.GetStatistics(ctx, id)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := stream.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					c.log.Error("Error while receiving message",
						zap.Error(err),
					)

					return
				}

				c.values <- msg.GetFrequency()
			}
		}
	}()

	return nil
}

func (c *Client) GetValues() <-chan float64 {
	return c.values
}
