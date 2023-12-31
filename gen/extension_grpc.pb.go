// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: extension.proto

package extension

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

// ExtensionServiceClient is the client API for ExtensionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionServiceClient interface {
	// Name is used to get the name of the extension.
	Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error)
	// ListMethods is used to list the methods which the extension provides.
	ListMethods(ctx context.Context, in *ListMethodsRequest, opts ...grpc.CallOption) (*ListMethodsResponse, error)
	// Execute is used to execute a method provided by the extension.
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	// Initialize is used to create a new extension instance
	Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*InitializeResponse, error)
}

type extensionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionServiceClient(cc grpc.ClientConnInterface) ExtensionServiceClient {
	return &extensionServiceClient{cc}
}

func (c *extensionServiceClient) Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/extension.ExtensionService/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionServiceClient) ListMethods(ctx context.Context, in *ListMethodsRequest, opts ...grpc.CallOption) (*ListMethodsResponse, error) {
	out := new(ListMethodsResponse)
	err := c.cc.Invoke(ctx, "/extension.ExtensionService/ListMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/extension.ExtensionService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extensionServiceClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*InitializeResponse, error) {
	out := new(InitializeResponse)
	err := c.cc.Invoke(ctx, "/extension.ExtensionService/Initialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionServiceServer is the server API for ExtensionService service.
// All implementations must embed UnimplementedExtensionServiceServer
// for forward compatibility
type ExtensionServiceServer interface {
	// Name is used to get the name of the extension.
	Name(context.Context, *NameRequest) (*NameResponse, error)
	// ListMethods is used to list the methods which the extension provides.
	ListMethods(context.Context, *ListMethodsRequest) (*ListMethodsResponse, error)
	// Execute is used to execute a method provided by the extension.
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	// Initialize is used to create a new extension instance
	Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error)
	mustEmbedUnimplementedExtensionServiceServer()
}

// UnimplementedExtensionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionServiceServer struct {
}

func (UnimplementedExtensionServiceServer) Name(context.Context, *NameRequest) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (UnimplementedExtensionServiceServer) ListMethods(context.Context, *ListMethodsRequest) (*ListMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMethods not implemented")
}
func (UnimplementedExtensionServiceServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedExtensionServiceServer) Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialize not implemented")
}
func (UnimplementedExtensionServiceServer) mustEmbedUnimplementedExtensionServiceServer() {}

// UnsafeExtensionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionServiceServer will
// result in compilation errors.
type UnsafeExtensionServiceServer interface {
	mustEmbedUnimplementedExtensionServiceServer()
}

func RegisterExtensionServiceServer(s grpc.ServiceRegistrar, srv ExtensionServiceServer) {
	s.RegisterService(&ExtensionService_ServiceDesc, srv)
}

func _ExtensionService_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServiceServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extension.ExtensionService/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServiceServer).Name(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionService_ListMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServiceServer).ListMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extension.ExtensionService/ListMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServiceServer).ListMethods(ctx, req.(*ListMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extension.ExtensionService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServiceServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtensionService_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionServiceServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extension.ExtensionService/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionServiceServer).Initialize(ctx, req.(*InitializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtensionService_ServiceDesc is the grpc.ServiceDesc for ExtensionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "extension.ExtensionService",
	HandlerType: (*ExtensionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _ExtensionService_Name_Handler,
		},
		{
			MethodName: "ListMethods",
			Handler:    _ExtensionService_ListMethods_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _ExtensionService_Execute_Handler,
		},
		{
			MethodName: "Initialize",
			Handler:    _ExtensionService_Initialize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extension.proto",
}
