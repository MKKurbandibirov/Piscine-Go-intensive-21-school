// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/transmitter.proto

package generated

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransmitterClient is the client API for Transmitter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransmitterClient interface {
	GetStatistics(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Transmitter_GetStatisticsClient, error)
}

type transmitterClient struct {
	cc grpc.ClientConnInterface
}

func NewTransmitterClient(cc grpc.ClientConnInterface) TransmitterClient {
	return &transmitterClient{cc}
}

func (c *transmitterClient) GetStatistics(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Transmitter_GetStatisticsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transmitter_ServiceDesc.Streams[0], "/transmitter.Transmitter/GetStatistics", opts...)
	if err != nil {
		return nil, err
	}
	x := &transmitterGetStatisticsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transmitter_GetStatisticsClient interface {
	Recv() (*StatisticValue, error)
	grpc.ClientStream
}

type transmitterGetStatisticsClient struct {
	grpc.ClientStream
}

func (x *transmitterGetStatisticsClient) Recv() (*StatisticValue, error) {
	m := new(StatisticValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransmitterServer is the server API for Transmitter service.
// All implementations must embed UnimplementedTransmitterServer
// for forward compatibility
type TransmitterServer interface {
	GetStatistics(*empty.Empty, Transmitter_GetStatisticsServer) error
	mustEmbedUnimplementedTransmitterServer()
}

// UnimplementedTransmitterServer must be embedded to have forward compatible implementations.
type UnimplementedTransmitterServer struct {
}

func (UnimplementedTransmitterServer) GetStatistics(*empty.Empty, Transmitter_GetStatisticsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedTransmitterServer) mustEmbedUnimplementedTransmitterServer() {}

// UnsafeTransmitterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransmitterServer will
// result in compilation errors.
type UnsafeTransmitterServer interface {
	mustEmbedUnimplementedTransmitterServer()
}

func RegisterTransmitterServer(s grpc.ServiceRegistrar, srv TransmitterServer) {
	s.RegisterService(&Transmitter_ServiceDesc, srv)
}

func _Transmitter_GetStatistics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransmitterServer).GetStatistics(m, &transmitterGetStatisticsServer{stream})
}

type Transmitter_GetStatisticsServer interface {
	Send(*StatisticValue) error
	grpc.ServerStream
}

type transmitterGetStatisticsServer struct {
	grpc.ServerStream
}

func (x *transmitterGetStatisticsServer) Send(m *StatisticValue) error {
	return x.ServerStream.SendMsg(m)
}

// Transmitter_ServiceDesc is the grpc.ServiceDesc for Transmitter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transmitter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transmitter.Transmitter",
	HandlerType: (*TransmitterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatistics",
			Handler:       _Transmitter_GetStatistics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/transmitter.proto",
}
