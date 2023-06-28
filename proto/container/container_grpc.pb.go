// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/container/container.proto

package container

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

// ContainerServiceClient is the client API for ContainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerServiceClient interface {
	// Sends a greeting
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerReply, error)
}

type containerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerServiceClient(cc grpc.ClientConnInterface) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerReply, error) {
	out := new(CreateContainerReply)
	err := c.cc.Invoke(ctx, "/container.ContainerService/CreateContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServiceServer is the server API for ContainerService service.
// All implementations must embed UnimplementedContainerServiceServer
// for forward compatibility
type ContainerServiceServer interface {
	// Sends a greeting
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerReply, error)
	mustEmbedUnimplementedContainerServiceServer()
}

// UnimplementedContainerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContainerServiceServer struct {
}

func (UnimplementedContainerServiceServer) CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (UnimplementedContainerServiceServer) mustEmbedUnimplementedContainerServiceServer() {}

// UnsafeContainerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerServiceServer will
// result in compilation errors.
type UnsafeContainerServiceServer interface {
	mustEmbedUnimplementedContainerServiceServer()
}

func RegisterContainerServiceServer(s grpc.ServiceRegistrar, srv ContainerServiceServer) {
	s.RegisterService(&ContainerService_ServiceDesc, srv)
}

func _ContainerService_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.ContainerService/CreateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).CreateContainer(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContainerService_ServiceDesc is the grpc.ServiceDesc for ContainerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "container.ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContainer",
			Handler:    _ContainerService_CreateContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/container/container.proto",
}
