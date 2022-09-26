// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/service.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AntiBruteForceClient is the client API for AntiBruteForce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AntiBruteForceClient interface {
	Try(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Status, error)
	ClearBucket(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddWhiteNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error)
	AddBlackNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error)
	RemoveWhiteNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error)
	RemoveBlackNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error)
}

type antiBruteForceClient struct {
	cc grpc.ClientConnInterface
}

func NewAntiBruteForceClient(cc grpc.ClientConnInterface) AntiBruteForceClient {
	return &antiBruteForceClient{cc}
}

func (c *antiBruteForceClient) Try(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.AntiBruteForce/Try", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteForceClient) ClearBucket(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/service.AntiBruteForce/ClearBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteForceClient) AddWhiteNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.AntiBruteForce/AddWhiteNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteForceClient) AddBlackNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.AntiBruteForce/AddBlackNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteForceClient) RemoveWhiteNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.AntiBruteForce/RemoveWhiteNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *antiBruteForceClient) RemoveBlackNet(ctx context.Context, in *IpRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service.AntiBruteForce/RemoveBlackNet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AntiBruteForceServer is the server API for AntiBruteForce service.
// All implementations must embed UnimplementedAntiBruteForceServer
// for forward compatibility
type AntiBruteForceServer interface {
	Try(context.Context, *CheckRequest) (*Status, error)
	ClearBucket(context.Context, *CheckRequest) (*empty.Empty, error)
	AddWhiteNet(context.Context, *IpRequest) (*Status, error)
	AddBlackNet(context.Context, *IpRequest) (*Status, error)
	RemoveWhiteNet(context.Context, *IpRequest) (*Status, error)
	RemoveBlackNet(context.Context, *IpRequest) (*Status, error)
	mustEmbedUnimplementedAntiBruteForceServer()
}

// UnimplementedAntiBruteForceServer must be embedded to have forward compatible implementations.
type UnimplementedAntiBruteForceServer struct {
}

func (UnimplementedAntiBruteForceServer) Try(context.Context, *CheckRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Try not implemented")
}
func (UnimplementedAntiBruteForceServer) ClearBucket(context.Context, *CheckRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearBucket not implemented")
}
func (UnimplementedAntiBruteForceServer) AddWhiteNet(context.Context, *IpRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWhiteNet not implemented")
}
func (UnimplementedAntiBruteForceServer) AddBlackNet(context.Context, *IpRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlackNet not implemented")
}
func (UnimplementedAntiBruteForceServer) RemoveWhiteNet(context.Context, *IpRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWhiteNet not implemented")
}
func (UnimplementedAntiBruteForceServer) RemoveBlackNet(context.Context, *IpRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlackNet not implemented")
}
func (UnimplementedAntiBruteForceServer) mustEmbedUnimplementedAntiBruteForceServer() {}

// UnsafeAntiBruteForceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AntiBruteForceServer will
// result in compilation errors.
type UnsafeAntiBruteForceServer interface {
	mustEmbedUnimplementedAntiBruteForceServer()
}

func RegisterAntiBruteForceServer(s grpc.ServiceRegistrar, srv AntiBruteForceServer) {
	s.RegisterService(&AntiBruteForce_ServiceDesc, srv)
}

func _AntiBruteForce_Try_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteForceServer).Try(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AntiBruteForce/Try",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteForceServer).Try(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteForce_ClearBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteForceServer).ClearBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AntiBruteForce/ClearBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteForceServer).ClearBucket(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteForce_AddWhiteNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteForceServer).AddWhiteNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AntiBruteForce/AddWhiteNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteForceServer).AddWhiteNet(ctx, req.(*IpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteForce_AddBlackNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteForceServer).AddBlackNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AntiBruteForce/AddBlackNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteForceServer).AddBlackNet(ctx, req.(*IpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteForce_RemoveWhiteNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteForceServer).RemoveWhiteNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AntiBruteForce/RemoveWhiteNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteForceServer).RemoveWhiteNet(ctx, req.(*IpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AntiBruteForce_RemoveBlackNet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AntiBruteForceServer).RemoveBlackNet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AntiBruteForce/RemoveBlackNet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AntiBruteForceServer).RemoveBlackNet(ctx, req.(*IpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AntiBruteForce_ServiceDesc is the grpc.ServiceDesc for AntiBruteForce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AntiBruteForce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AntiBruteForce",
	HandlerType: (*AntiBruteForceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Try",
			Handler:    _AntiBruteForce_Try_Handler,
		},
		{
			MethodName: "ClearBucket",
			Handler:    _AntiBruteForce_ClearBucket_Handler,
		},
		{
			MethodName: "AddWhiteNet",
			Handler:    _AntiBruteForce_AddWhiteNet_Handler,
		},
		{
			MethodName: "AddBlackNet",
			Handler:    _AntiBruteForce_AddBlackNet_Handler,
		},
		{
			MethodName: "RemoveWhiteNet",
			Handler:    _AntiBruteForce_RemoveWhiteNet_Handler,
		},
		{
			MethodName: "RemoveBlackNet",
			Handler:    _AntiBruteForce_RemoveBlackNet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}
