// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: blame.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlameServiceClient is the client API for BlameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlameServiceClient interface {
	Blame(ctx context.Context, in *BlameRequest, opts ...grpc.CallOption) (BlameService_BlameClient, error)
}

type blameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlameServiceClient(cc grpc.ClientConnInterface) BlameServiceClient {
	return &blameServiceClient{cc}
}

func (c *blameServiceClient) Blame(ctx context.Context, in *BlameRequest, opts ...grpc.CallOption) (BlameService_BlameClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlameService_ServiceDesc.Streams[0], "/rpc.BlameService/Blame", opts...)
	if err != nil {
		return nil, err
	}
	x := &blameServiceBlameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlameService_BlameClient interface {
	Recv() (*BlamePart, error)
	grpc.ClientStream
}

type blameServiceBlameClient struct {
	grpc.ClientStream
}

func (x *blameServiceBlameClient) Recv() (*BlamePart, error) {
	m := new(BlamePart)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlameServiceServer is the server API for BlameService service.
// All implementations must embed UnimplementedBlameServiceServer
// for forward compatibility
type BlameServiceServer interface {
	Blame(*BlameRequest, BlameService_BlameServer) error
	mustEmbedUnimplementedBlameServiceServer()
}

// UnimplementedBlameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlameServiceServer struct {
}

func (UnimplementedBlameServiceServer) Blame(*BlameRequest, BlameService_BlameServer) error {
	return status.Errorf(codes.Unimplemented, "method Blame not implemented")
}
func (UnimplementedBlameServiceServer) mustEmbedUnimplementedBlameServiceServer() {}

// UnsafeBlameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlameServiceServer will
// result in compilation errors.
type UnsafeBlameServiceServer interface {
	mustEmbedUnimplementedBlameServiceServer()
}

func RegisterBlameServiceServer(s grpc.ServiceRegistrar, srv BlameServiceServer) {
	s.RegisterService(&BlameService_ServiceDesc, srv)
}

func _BlameService_Blame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlameServiceServer).Blame(m, &blameServiceBlameServer{stream})
}

type BlameService_BlameServer interface {
	Send(*BlamePart) error
	grpc.ServerStream
}

type blameServiceBlameServer struct {
	grpc.ServerStream
}

func (x *blameServiceBlameServer) Send(m *BlamePart) error {
	return x.ServerStream.SendMsg(m)
}

// BlameService_ServiceDesc is the grpc.ServiceDesc for BlameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.BlameService",
	HandlerType: (*BlameServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Blame",
			Handler:       _BlameService_Blame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blame.proto",
}
