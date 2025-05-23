// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: operator.proto

package proto

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
	Operator_Add_FullMethodName  = "/Operator/Add"
	Operator_Edit_FullMethodName = "/Operator/Edit"
	Operator_List_FullMethodName = "/Operator/List"
)

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorClient interface {
	Add(ctx context.Context, in *OperatorAddReq, opts ...grpc.CallOption) (*OperatorAddResp, error)
	Edit(ctx context.Context, in *OperatorEditReq, opts ...grpc.CallOption) (*OperatorEditResp, error)
	List(ctx context.Context, in *OperatorListReq, opts ...grpc.CallOption) (*OperatorListResp, error)
}

type operatorClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorClient(cc grpc.ClientConnInterface) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) Add(ctx context.Context, in *OperatorAddReq, opts ...grpc.CallOption) (*OperatorAddResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperatorAddResp)
	err := c.cc.Invoke(ctx, Operator_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) Edit(ctx context.Context, in *OperatorEditReq, opts ...grpc.CallOption) (*OperatorEditResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperatorEditResp)
	err := c.cc.Invoke(ctx, Operator_Edit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) List(ctx context.Context, in *OperatorListReq, opts ...grpc.CallOption) (*OperatorListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperatorListResp)
	err := c.cc.Invoke(ctx, Operator_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
// All implementations must embed UnimplementedOperatorServer
// for forward compatibility.
type OperatorServer interface {
	Add(context.Context, *OperatorAddReq) (*OperatorAddResp, error)
	Edit(context.Context, *OperatorEditReq) (*OperatorEditResp, error)
	List(context.Context, *OperatorListReq) (*OperatorListResp, error)
	mustEmbedUnimplementedOperatorServer()
}

// UnimplementedOperatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOperatorServer struct{}

func (UnimplementedOperatorServer) Add(context.Context, *OperatorAddReq) (*OperatorAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedOperatorServer) Edit(context.Context, *OperatorEditReq) (*OperatorEditResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedOperatorServer) List(context.Context, *OperatorListReq) (*OperatorListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOperatorServer) mustEmbedUnimplementedOperatorServer() {}
func (UnimplementedOperatorServer) testEmbeddedByValue()                  {}

// UnsafeOperatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServer will
// result in compilation errors.
type UnsafeOperatorServer interface {
	mustEmbedUnimplementedOperatorServer()
}

func RegisterOperatorServer(s grpc.ServiceRegistrar, srv OperatorServer) {
	// If the following call pancis, it indicates UnimplementedOperatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Operator_ServiceDesc, srv)
}

func _Operator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).Add(ctx, req.(*OperatorAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorEditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_Edit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).Edit(ctx, req.(*OperatorEditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperatorListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operator_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).List(ctx, req.(*OperatorListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Operator_ServiceDesc is the grpc.ServiceDesc for Operator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Operator_Add_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _Operator_Edit_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Operator_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operator.proto",
}
