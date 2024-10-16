// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: request.proto

package request

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RequestService_CreateRequest_FullMethodName = "/request.RequestService/CreateRequest"
	RequestService_GetRequest_FullMethodName    = "/request.RequestService/GetRequest"
	RequestService_DeleteRequest_FullMethodName = "/request.RequestService/DeleteRequest"
)

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	CreateRequest(ctx context.Context, in *CreateRequestRequest, opts ...grpc.CallOption) (*CreateRequestResponse, error)
	GetRequest(ctx context.Context, in *GetRequestRequest, opts ...grpc.CallOption) (*GetRequestResponse, error)
	DeleteRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Void, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) CreateRequest(ctx context.Context, in *CreateRequestRequest, opts ...grpc.CallOption) (*CreateRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRequestResponse)
	err := c.cc.Invoke(ctx, RequestService_CreateRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetRequest(ctx context.Context, in *GetRequestRequest, opts ...grpc.CallOption) (*GetRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRequestResponse)
	err := c.cc.Invoke(ctx, RequestService_GetRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) DeleteRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, RequestService_DeleteRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations must embed UnimplementedRequestServiceServer
// for forward compatibility.
type RequestServiceServer interface {
	CreateRequest(context.Context, *CreateRequestRequest) (*CreateRequestResponse, error)
	GetRequest(context.Context, *GetRequestRequest) (*GetRequestResponse, error)
	DeleteRequest(context.Context, *Request) (*Void, error)
	mustEmbedUnimplementedRequestServiceServer()
}

// UnimplementedRequestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRequestServiceServer struct{}

func (UnimplementedRequestServiceServer) CreateRequest(context.Context, *CreateRequestRequest) (*CreateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequest not implemented")
}
func (UnimplementedRequestServiceServer) GetRequest(context.Context, *GetRequestRequest) (*GetRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequest not implemented")
}
func (UnimplementedRequestServiceServer) DeleteRequest(context.Context, *Request) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequest not implemented")
}
func (UnimplementedRequestServiceServer) mustEmbedUnimplementedRequestServiceServer() {}
func (UnimplementedRequestServiceServer) testEmbeddedByValue()                        {}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	// If the following call pancis, it indicates UnimplementedRequestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_CreateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).CreateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_CreateRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).CreateRequest(ctx, req.(*CreateRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_GetRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequest(ctx, req.(*GetRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_DeleteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).DeleteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RequestService_DeleteRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).DeleteRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "request.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRequest",
			Handler:    _RequestService_CreateRequest_Handler,
		},
		{
			MethodName: "GetRequest",
			Handler:    _RequestService_GetRequest_Handler,
		},
		{
			MethodName: "DeleteRequest",
			Handler:    _RequestService_DeleteRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "request.proto",
}
