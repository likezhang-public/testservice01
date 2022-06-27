// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: Testservice01.proto

package grpc_go

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

// Testservice01Client is the client API for Testservice01 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Testservice01Client interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloworldRequest, opts ...grpc.CallOption) (*HelloworldResponse, error)
}

type testservice01Client struct {
	cc grpc.ClientConnInterface
}

func NewTestservice01Client(cc grpc.ClientConnInterface) Testservice01Client {
	return &testservice01Client{cc}
}

func (c *testservice01Client) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/com.test1.Testservice01/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testservice01Client) Helloworld(ctx context.Context, in *HelloworldRequest, opts ...grpc.CallOption) (*HelloworldResponse, error) {
	out := new(HelloworldResponse)
	err := c.cc.Invoke(ctx, "/com.test1.Testservice01/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Testservice01Server is the server API for Testservice01 service.
// All implementations must embed UnimplementedTestservice01Server
// for forward compatibility
type Testservice01Server interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloworldRequest) (*HelloworldResponse, error)
	mustEmbedUnimplementedTestservice01Server()
}

// UnimplementedTestservice01Server must be embedded to have forward compatible implementations.
type UnimplementedTestservice01Server struct {
}

func (UnimplementedTestservice01Server) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedTestservice01Server) Helloworld(context.Context, *HelloworldRequest) (*HelloworldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedTestservice01Server) mustEmbedUnimplementedTestservice01Server() {}

// UnsafeTestservice01Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Testservice01Server will
// result in compilation errors.
type UnsafeTestservice01Server interface {
	mustEmbedUnimplementedTestservice01Server()
}

func RegisterTestservice01Server(s grpc.ServiceRegistrar, srv Testservice01Server) {
	s.RegisterService(&Testservice01_ServiceDesc, srv)
}

func _Testservice01_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Testservice01Server).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.test1.Testservice01/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Testservice01Server).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Testservice01_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloworldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Testservice01Server).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.test1.Testservice01/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Testservice01Server).Helloworld(ctx, req.(*HelloworldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Testservice01_ServiceDesc is the grpc.ServiceDesc for Testservice01 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Testservice01_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.test1.Testservice01",
	HandlerType: (*Testservice01Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _Testservice01_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _Testservice01_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Testservice01.proto",
}
