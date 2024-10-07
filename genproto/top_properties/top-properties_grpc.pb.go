// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: top-properties.proto

package top_properties

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
	TopPropertiesService_Get_FullMethodName    = "/top_properties.TopPropertiesService/Get"
	TopPropertiesService_GetAll_FullMethodName = "/top_properties.TopPropertiesService/GetAll"
	TopPropertiesService_Create_FullMethodName = "/top_properties.TopPropertiesService/Create"
	TopPropertiesService_Update_FullMethodName = "/top_properties.TopPropertiesService/Update"
	TopPropertiesService_Delete_FullMethodName = "/top_properties.TopPropertiesService/Delete"
)

// TopPropertiesServiceClient is the client API for TopPropertiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopPropertiesServiceClient interface {
	Get(ctx context.Context, in *GetTopPropertyReq, opts ...grpc.CallOption) (*GetTopPropertyRes, error)
	GetAll(ctx context.Context, in *GetAllTopPropertyReq, opts ...grpc.CallOption) (*GetAllTopPropertyRes, error)
	Create(ctx context.Context, in *CreateTopPropertyReq, opts ...grpc.CallOption) (*CreateTopPropertyRes, error)
	Update(ctx context.Context, in *UpdateTopPropertyReq, opts ...grpc.CallOption) (*UpdateTopPropertyRes, error)
	Delete(ctx context.Context, in *DeleteTopPropertyReq, opts ...grpc.CallOption) (*DeleteTopPropertyRes, error)
}

type topPropertiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopPropertiesServiceClient(cc grpc.ClientConnInterface) TopPropertiesServiceClient {
	return &topPropertiesServiceClient{cc}
}

func (c *topPropertiesServiceClient) Get(ctx context.Context, in *GetTopPropertyReq, opts ...grpc.CallOption) (*GetTopPropertyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopPropertyRes)
	err := c.cc.Invoke(ctx, TopPropertiesService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topPropertiesServiceClient) GetAll(ctx context.Context, in *GetAllTopPropertyReq, opts ...grpc.CallOption) (*GetAllTopPropertyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllTopPropertyRes)
	err := c.cc.Invoke(ctx, TopPropertiesService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topPropertiesServiceClient) Create(ctx context.Context, in *CreateTopPropertyReq, opts ...grpc.CallOption) (*CreateTopPropertyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTopPropertyRes)
	err := c.cc.Invoke(ctx, TopPropertiesService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topPropertiesServiceClient) Update(ctx context.Context, in *UpdateTopPropertyReq, opts ...grpc.CallOption) (*UpdateTopPropertyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTopPropertyRes)
	err := c.cc.Invoke(ctx, TopPropertiesService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topPropertiesServiceClient) Delete(ctx context.Context, in *DeleteTopPropertyReq, opts ...grpc.CallOption) (*DeleteTopPropertyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTopPropertyRes)
	err := c.cc.Invoke(ctx, TopPropertiesService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopPropertiesServiceServer is the server API for TopPropertiesService service.
// All implementations must embed UnimplementedTopPropertiesServiceServer
// for forward compatibility.
type TopPropertiesServiceServer interface {
	Get(context.Context, *GetTopPropertyReq) (*GetTopPropertyRes, error)
	GetAll(context.Context, *GetAllTopPropertyReq) (*GetAllTopPropertyRes, error)
	Create(context.Context, *CreateTopPropertyReq) (*CreateTopPropertyRes, error)
	Update(context.Context, *UpdateTopPropertyReq) (*UpdateTopPropertyRes, error)
	Delete(context.Context, *DeleteTopPropertyReq) (*DeleteTopPropertyRes, error)
	mustEmbedUnimplementedTopPropertiesServiceServer()
}

// UnimplementedTopPropertiesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTopPropertiesServiceServer struct{}

func (UnimplementedTopPropertiesServiceServer) Get(context.Context, *GetTopPropertyReq) (*GetTopPropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTopPropertiesServiceServer) GetAll(context.Context, *GetAllTopPropertyReq) (*GetAllTopPropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTopPropertiesServiceServer) Create(context.Context, *CreateTopPropertyReq) (*CreateTopPropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTopPropertiesServiceServer) Update(context.Context, *UpdateTopPropertyReq) (*UpdateTopPropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTopPropertiesServiceServer) Delete(context.Context, *DeleteTopPropertyReq) (*DeleteTopPropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTopPropertiesServiceServer) mustEmbedUnimplementedTopPropertiesServiceServer() {}
func (UnimplementedTopPropertiesServiceServer) testEmbeddedByValue()                              {}

// UnsafeTopPropertiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopPropertiesServiceServer will
// result in compilation errors.
type UnsafeTopPropertiesServiceServer interface {
	mustEmbedUnimplementedTopPropertiesServiceServer()
}

func RegisterTopPropertiesServiceServer(s grpc.ServiceRegistrar, srv TopPropertiesServiceServer) {
	// If the following call pancis, it indicates UnimplementedTopPropertiesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TopPropertiesService_ServiceDesc, srv)
}

func _TopPropertiesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopPropertiesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopPropertiesService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopPropertiesServiceServer).Get(ctx, req.(*GetTopPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopPropertiesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTopPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopPropertiesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopPropertiesService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopPropertiesServiceServer).GetAll(ctx, req.(*GetAllTopPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopPropertiesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopPropertiesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopPropertiesService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopPropertiesServiceServer).Create(ctx, req.(*CreateTopPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopPropertiesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopPropertiesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopPropertiesService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopPropertiesServiceServer).Update(ctx, req.(*UpdateTopPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopPropertiesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopPropertiesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopPropertiesService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopPropertiesServiceServer).Delete(ctx, req.(*DeleteTopPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TopPropertiesService_ServiceDesc is the grpc.ServiceDesc for TopPropertiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopPropertiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "top_properties.TopPropertiesService",
	HandlerType: (*TopPropertiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TopPropertiesService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TopPropertiesService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TopPropertiesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TopPropertiesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TopPropertiesService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "top-properties.proto",
}
