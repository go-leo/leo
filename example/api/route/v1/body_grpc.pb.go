// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: route/v1/body.proto

package route

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "google.golang.org/genproto/googleapis/rpc/http"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Body_StarBody_FullMethodName          = "/leo.example.route.body.Body/StarBody"
	Body_NamedBody_FullMethodName         = "/leo.example.route.body.Body/NamedBody"
	Body_NonBody_FullMethodName           = "/leo.example.route.body.Body/NonBody"
	Body_HttpBodyStarBody_FullMethodName  = "/leo.example.route.body.Body/HttpBodyStarBody"
	Body_HttpBodyNamedBody_FullMethodName = "/leo.example.route.body.Body/HttpBodyNamedBody"
	Body_HttpRequest_FullMethodName       = "/leo.example.route.body.Body/HttpRequest"
)

// BodyClient is the client API for Body service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BodyClient interface {
	StarBody(ctx context.Context, in *BodyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, in *BodyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NonBody(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, in *HttpBodyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HttpRequest(ctx context.Context, in *http.HttpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bodyClient struct {
	cc grpc.ClientConnInterface
}

func NewBodyClient(cc grpc.ClientConnInterface) BodyClient {
	return &bodyClient{cc}
}

func (c *bodyClient) StarBody(ctx context.Context, in *BodyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Body_StarBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bodyClient) NamedBody(ctx context.Context, in *BodyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Body_NamedBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bodyClient) NonBody(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Body_NonBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bodyClient) HttpBodyStarBody(ctx context.Context, in *httpbody.HttpBody, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Body_HttpBodyStarBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bodyClient) HttpBodyNamedBody(ctx context.Context, in *HttpBodyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Body_HttpBodyNamedBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bodyClient) HttpRequest(ctx context.Context, in *http.HttpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Body_HttpRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BodyServer is the server API for Body service.
// All implementations must embed UnimplementedBodyServer
// for forward compatibility
type BodyServer interface {
	StarBody(context.Context, *BodyRequest) (*emptypb.Empty, error)
	NamedBody(context.Context, *BodyRequest) (*emptypb.Empty, error)
	NonBody(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(context.Context, *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(context.Context, *HttpBodyRequest) (*emptypb.Empty, error)
	HttpRequest(context.Context, *http.HttpRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBodyServer()
}

// UnimplementedBodyServer must be embedded to have forward compatible implementations.
type UnimplementedBodyServer struct {
}

func (UnimplementedBodyServer) StarBody(context.Context, *BodyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StarBody not implemented")
}
func (UnimplementedBodyServer) NamedBody(context.Context, *BodyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamedBody not implemented")
}
func (UnimplementedBodyServer) NonBody(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NonBody not implemented")
}
func (UnimplementedBodyServer) HttpBodyStarBody(context.Context, *httpbody.HttpBody) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HttpBodyStarBody not implemented")
}
func (UnimplementedBodyServer) HttpBodyNamedBody(context.Context, *HttpBodyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HttpBodyNamedBody not implemented")
}
func (UnimplementedBodyServer) HttpRequest(context.Context, *http.HttpRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HttpRequest not implemented")
}
func (UnimplementedBodyServer) mustEmbedUnimplementedBodyServer() {}

// UnsafeBodyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BodyServer will
// result in compilation errors.
type UnsafeBodyServer interface {
	mustEmbedUnimplementedBodyServer()
}

func RegisterBodyServer(s grpc.ServiceRegistrar, srv BodyServer) {
	s.RegisterService(&Body_ServiceDesc, srv)
}

func _Body_StarBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyServer).StarBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Body_StarBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyServer).StarBody(ctx, req.(*BodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Body_NamedBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyServer).NamedBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Body_NamedBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyServer).NamedBody(ctx, req.(*BodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Body_NonBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyServer).NonBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Body_NonBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyServer).NonBody(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Body_HttpBodyStarBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(httpbody.HttpBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyServer).HttpBodyStarBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Body_HttpBodyStarBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyServer).HttpBodyStarBody(ctx, req.(*httpbody.HttpBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Body_HttpBodyNamedBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HttpBodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyServer).HttpBodyNamedBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Body_HttpBodyNamedBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyServer).HttpBodyNamedBody(ctx, req.(*HttpBodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Body_HttpRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(http.HttpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BodyServer).HttpRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Body_HttpRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BodyServer).HttpRequest(ctx, req.(*http.HttpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Body_ServiceDesc is the grpc.ServiceDesc for Body service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Body_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "leo.example.route.body.Body",
	HandlerType: (*BodyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StarBody",
			Handler:    _Body_StarBody_Handler,
		},
		{
			MethodName: "NamedBody",
			Handler:    _Body_NamedBody_Handler,
		},
		{
			MethodName: "NonBody",
			Handler:    _Body_NonBody_Handler,
		},
		{
			MethodName: "HttpBodyStarBody",
			Handler:    _Body_HttpBodyStarBody_Handler,
		},
		{
			MethodName: "HttpBodyNamedBody",
			Handler:    _Body_HttpBodyNamedBody_Handler,
		},
		{
			MethodName: "HttpRequest",
			Handler:    _Body_HttpRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route/v1/body.proto",
}
