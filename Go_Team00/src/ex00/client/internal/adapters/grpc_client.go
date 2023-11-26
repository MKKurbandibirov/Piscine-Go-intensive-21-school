package adapters

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	pb "messages/generated"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.TransmitterClient
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

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println(msg.String())
	}

	return nil
}
