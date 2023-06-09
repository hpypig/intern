// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pb/NewsPush.proto

package pb

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

// NewsPusherClient is the client API for NewsPusher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsPusherClient interface {
	ImportantNewsPush(ctx context.Context, in *ImportantNewsRequest, opts ...grpc.CallOption) (NewsPusher_ImportantNewsPushClient, error)
}

type newsPusherClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsPusherClient(cc grpc.ClientConnInterface) NewsPusherClient {
	return &newsPusherClient{cc}
}

func (c *newsPusherClient) ImportantNewsPush(ctx context.Context, in *ImportantNewsRequest, opts ...grpc.CallOption) (NewsPusher_ImportantNewsPushClient, error) {
	stream, err := c.cc.NewStream(ctx, &NewsPusher_ServiceDesc.Streams[0], "/pb.NewsPusher/ImportantNewsPush", opts...)
	if err != nil {
		return nil, err
	}
	x := &newsPusherImportantNewsPushClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NewsPusher_ImportantNewsPushClient interface {
	Recv() (*UpdatedNewsResponse, error)
	grpc.ClientStream
}

type newsPusherImportantNewsPushClient struct {
	grpc.ClientStream
}

func (x *newsPusherImportantNewsPushClient) Recv() (*UpdatedNewsResponse, error) {
	m := new(UpdatedNewsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NewsPusherServer is the server API for NewsPusher service.
// All implementations must embed UnimplementedNewsPusherServer
// for forward compatibility
type NewsPusherServer interface {
	ImportantNewsPush(*ImportantNewsRequest, NewsPusher_ImportantNewsPushServer) error
	mustEmbedUnimplementedNewsPusherServer()
}

// UnimplementedNewsPusherServer must be embedded to have forward compatible implementations.
type UnimplementedNewsPusherServer struct {
}

func (UnimplementedNewsPusherServer) ImportantNewsPush(*ImportantNewsRequest, NewsPusher_ImportantNewsPushServer) error {
	return status.Errorf(codes.Unimplemented, "method ImportantNewsPush not implemented")
}
func (UnimplementedNewsPusherServer) mustEmbedUnimplementedNewsPusherServer() {}

// UnsafeNewsPusherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsPusherServer will
// result in compilation errors.
type UnsafeNewsPusherServer interface {
	mustEmbedUnimplementedNewsPusherServer()
}

func RegisterNewsPusherServer(s grpc.ServiceRegistrar, srv NewsPusherServer) {
	s.RegisterService(&NewsPusher_ServiceDesc, srv)
}

func _NewsPusher_ImportantNewsPush_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImportantNewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NewsPusherServer).ImportantNewsPush(m, &newsPusherImportantNewsPushServer{stream})
}

type NewsPusher_ImportantNewsPushServer interface {
	Send(*UpdatedNewsResponse) error
	grpc.ServerStream
}

type newsPusherImportantNewsPushServer struct {
	grpc.ServerStream
}

func (x *newsPusherImportantNewsPushServer) Send(m *UpdatedNewsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// NewsPusher_ServiceDesc is the grpc.ServiceDesc for NewsPusher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsPusher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NewsPusher",
	HandlerType: (*NewsPusherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImportantNewsPush",
			Handler:       _NewsPusher_ImportantNewsPush_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/NewsPush.proto",
}
