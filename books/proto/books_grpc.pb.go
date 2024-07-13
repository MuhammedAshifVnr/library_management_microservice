// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.23.4
// source: proto/books.proto

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
	BookService_CreateBook_FullMethodName  = "/books.BookService/CreateBook"
	BookService_DeleteBook_FullMethodName  = "/books.BookService/DeleteBook"
	BookService_GetAllBooks_FullMethodName = "/books.BookService/GetAllBooks"
	BookService_GetBook_FullMethodName     = "/books.BookService/GetBook"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
	GetAllBooks(ctx context.Context, in *GetAllBooksRequest, opts ...grpc.CallOption) (*GetAllBooksResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, BookService_CreateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, BookService_DeleteBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetAllBooks(ctx context.Context, in *GetAllBooksRequest, opts ...grpc.CallOption) (*GetAllBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllBooksResponse)
	err := c.cc.Invoke(ctx, BookService_GetAllBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, BookService_GetBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
	GetAllBooks(context.Context, *GetAllBooksRequest) (*GetAllBooksResponse, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) GetAllBooks(context.Context, *GetAllBooksRequest) (*GetAllBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedBookServiceServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetAllBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetAllBooks(ctx, req.(*GetAllBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
		{
			MethodName: "GetAllBooks",
			Handler:    _BookService_GetAllBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/books.proto",
}