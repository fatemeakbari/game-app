// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: contract/presence/presence.proto

package presenceadapter

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

// PresenceClient is the client API for Presence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresenceClient interface {
	GetPresence(ctx context.Context, in *GetPresenceRequest, opts ...grpc.CallOption) (*GetPresenceResponse, error)
}

type presenceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresenceClient(cc grpc.ClientConnInterface) PresenceClient {
	return &presenceClient{cc}
}

func (c *presenceClient) GetPresence(ctx context.Context, in *GetPresenceRequest, opts ...grpc.CallOption) (*GetPresenceResponse, error) {
	out := new(GetPresenceResponse)
	err := c.cc.Invoke(ctx, "/presence.Presence/GetPresence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresenceServer is the server API for Presence service.
// All implementations must embed UnimplementedPresenceServer
// for forward compatibility
type PresenceServer interface {
	GetPresence(context.Context, *GetPresenceRequest) (*GetPresenceResponse, error)
	mustEmbedUnimplementedPresenceServer()
}

// UnimplementedPresenceServer must be embedded to have forward compatible implementations.
type UnimplementedPresenceServer struct {
}

func (UnimplementedPresenceServer) GetPresence(context.Context, *GetPresenceRequest) (*GetPresenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPresence not implemented")
}
func (UnimplementedPresenceServer) mustEmbedUnimplementedPresenceServer() {}

// UnsafePresenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresenceServer will
// result in compilation errors.
type UnsafePresenceServer interface {
	mustEmbedUnimplementedPresenceServer()
}

func RegisterPresenceServer(s grpc.ServiceRegistrar, srv PresenceServer) {
	s.RegisterService(&Presence_ServiceDesc, srv)
}

func _Presence_GetPresence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPresenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServer).GetPresence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/presence.Presence/GetPresence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServer).GetPresence(ctx, req.(*GetPresenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Presence_ServiceDesc is the grpc.ServiceDesc for Presence service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Presence_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "presence.Presence",
	HandlerType: (*PresenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPresence",
			Handler:    _Presence_GetPresence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract/presence/presence.proto",
}