// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bookshop.proto

package inventory

import (HashTableClient
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HashTable_GetURL_FullMethodName   = "/HashTable/GetURL"
	HashTable_GetValue_FullMethodName = "/HashTable/GetValue"
	HashTable_InsertKv_FullMethodName = "/HashTable/InsertKv"
)

// HashTableClient is the client API for HashTable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashTableClient interface {
	GetURL(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
	GetValue(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*ValueResponse, error)
	InsertKv(ctx context.Context, in *InsertValue, opts ...grpc.CallOption) (*UrlResponse, error)
}

type hashTableClient struct {
	cc grpc.ClientConnInterface
}

func NewHashTableClient(cc grpc.ClientConnInterface) HashTableClient {
	return &hashTableClient{cc}
}

func (c *hashTableClient) GetURL(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, HashTable_GetURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashTableClient) GetValue(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*ValueResponse, error) {
	out := new(ValueResponse)
	err := c.cc.Invoke(ctx, HashTable_GetValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashTableClient) InsertKv(ctx context.Context, in *InsertValue, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, HashTable_InsertKv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashTableServer is the server API for HashTable service.
// All implementations must embed UnimplementedHashTableServer
// for forward compatibility
type HashTableServer interface {
	GetURL(context.Context, *UrlRequest) (*UrlResponse, error)
	GetValue(context.Context, *UrlRequest) (*ValueResponse, error)
	InsertKv(context.Context, *InsertValue) (*UrlResponse, error)
	mustEmbedUnimplementedHashTableServer()
}

// UnimplementedHashTableServer must be embedded to have forward compatible implementations.
type UnimplementedHashTableServer struct {
}

func (UnimplementedHashTableServer) GetURL(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURL not implemented")
}
func (UnimplementedHashTableServer) GetValue(context.Context, *UrlRequest) (*ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedHashTableServer) InsertKv(context.Context, *InsertValue) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertKv not implemented")
}
func (UnimplementedHashTableServer) mustEmbedUnimplementedHashTableServer() {}

// UnsafeHashTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashTableServer will
// result in compilation errors.
type UnsafeHashTableServer interface {
	mustEmbedUnimplementedHashTableServer()
}

func RegisterHashTableServer(s grpc.ServiceRegistrar, srv HashTableServer) {
	s.RegisterService(&HashTable_ServiceDesc, srv)
}

func _HashTable_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashTableServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HashTable_GetURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashTableServer).GetURL(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashTable_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashTableServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HashTable_GetValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashTableServer).GetValue(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HashTable_InsertKv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashTableServer).InsertKv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HashTable_InsertKv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashTableServer).InsertKv(ctx, req.(*InsertValue))
	}
	return interceptor(ctx, in, info, handler)
}

// HashTable_ServiceDesc is the grpc.ServiceDesc for HashTable service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashTable_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HashTable",
	HandlerType: (*HashTableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetURL",
			Handler:    _HashTable_GetURL_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _HashTable_GetValue_Handler,
		},
		{
			MethodName: "InsertKv",
			Handler:    _HashTable_InsertKv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookshop.proto",
}
