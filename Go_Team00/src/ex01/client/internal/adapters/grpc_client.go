package adapters

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(addr, credentials)
	if err != nil {
		//TODO wrapping
		return err
	}

	c.conn = conn
	c.client = pb.NewTransmitterClient(c.conn)

	return nil
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) GetStatistics(ctx context.Context) error {
	stream, err := c.client.GetStatistics(ctx, &empty.Empty{})
	if err != nil {
		//TODO wrapping
		return err
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				//TODO logging
				return
			}

			c.values <- msg.GetFrequency()
		}
	}()

	return nil
}

func (c *Client) GetValues() <-chan float64 {
	return c.values
}
