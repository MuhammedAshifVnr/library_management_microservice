// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.23.4
// source: proto/book_management.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BookManagementService_CheckOutBook_FullMethodName = "/book_management.BookManagementService/CheckOutBook"
	BookManagementService_CheckInBook_FullMethodName  = "/book_management.BookManagementService/CheckInBook"
)

// BookManagementServiceClient is the client API for BookManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookManagementServiceClient interface {
	CheckOutBook(ctx context.Context, in *CheckOutBookRequest, opts ...grpc.CallOption) (*CheckOutBookResponse, error)
	CheckInBook(ctx context.Context, in *CheckInBookRequest, opts ...grpc.CallOption) (*CheckInBookResponse, error)
}

type bookManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookManagementServiceClient(cc grpc.ClientConnInterface) BookManagementServiceClient {
	return &bookManagementServiceClient{cc}
}

func (c *bookManagementServiceClient) CheckOutBook(ctx context.Context, in *CheckOutBookRequest, opts ...grpc.CallOption) (*CheckOutBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckOutBookResponse)
	err := c.cc.Invoke(ctx, BookManagementService_CheckOutBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookManagementServiceClient) CheckInBook(ctx context.Context, in *CheckInBookRequest, opts ...grpc.CallOption) (*CheckInBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckInBookResponse)
	err := c.cc.Invoke(ctx, BookManagementService_CheckInBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookManagementServiceServer is the server API for BookManagementService service.
// All implementations must embed UnimplementedBookManagementServiceServer
// for forward compatibility
type BookManagementServiceServer interface {
	CheckOutBook(context.Context, *CheckOutBookRequest) (*CheckOutBookResponse, error)
	CheckInBook(context.Context, *CheckInBookRequest) (*CheckInBookResponse, error)
	mustEmbedUnimplementedBookManagementServiceServer()
}

// UnimplementedBookManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookManagementServiceServer struct {
}

func (UnimplementedBookManagementServiceServer) CheckOutBook(context.Context, *CheckOutBookRequest) (*CheckOutBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOutBook not implemented")
}
func (UnimplementedBookManagementServiceServer) CheckInBook(context.Context, *CheckInBookRequest) (*CheckInBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInBook not implemented")
}
func (UnimplementedBookManagementServiceServer) mustEmbedUnimplementedBookManagementServiceServer() {}

// UnsafeBookManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookManagementServiceServer will
// result in compilation errors.
type UnsafeBookManagementServiceServer interface {
	mustEmbedUnimplementedBookManagementServiceServer()
}

func RegisterBookManagementServiceServer(s grpc.ServiceRegistrar, srv BookManagementServiceServer) {
	s.RegisterService(&BookManagementService_ServiceDesc, srv)
}

func _BookManagementService_CheckOutBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServiceServer).CheckOutBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookManagementService_CheckOutBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServiceServer).CheckOutBook(ctx, req.(*CheckOutBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookManagementService_CheckInBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServiceServer).CheckInBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookManagementService_CheckInBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServiceServer).CheckInBook(ctx, req.(*CheckInBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookManagementService_ServiceDesc is the grpc.ServiceDesc for BookManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_management.BookManagementService",
	HandlerType: (*BookManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckOutBook",
			Handler:    _BookManagementService_CheckOutBook_Handler,
		},
		{
			MethodName: "CheckInBook",
			Handler:    _BookManagementService_CheckInBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book_management.proto",
}
