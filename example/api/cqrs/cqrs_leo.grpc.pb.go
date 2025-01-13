// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package cqrs

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

func NewCQRSGrpcServer(svc CQRSService, middlewares ...endpoint.Middleware) CQRSService {
	endpoints := newCQRSServerEndpoints(svc, middlewares...)
	transports := &cQRSGrpcServerTransports{endpoints: endpoints}
	return &cQRSGrpcServer{
		createUser: transports.CreateUser(),
		deleteUser: transports.DeleteUser(),
		updateUser: transports.UpdateUser(),
		findUser:   transports.FindUser(),
	}
}

func NewCQRSGrpcClient(target string, opts ...grpcx.ClientOption) CQRSService {
	options := grpcx.NewClientOptions(opts...)
	transports := newCQRSGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newCQRSClientEndpoints(target, transports, options.Builder(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newCQRSClientService(endpoints, grpcx.GrpcClient)
}

type cQRSGrpcServerTransports struct {
	endpoints CQRSServerEndpoints
}

func (t *cQRSGrpcServerTransports) CreateUser() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.CreateUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/pb.CQRS/CreateUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *cQRSGrpcServerTransports) DeleteUser() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.DeleteUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/pb.CQRS/DeleteUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *cQRSGrpcServerTransports) UpdateUser() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.UpdateUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/pb.CQRS/UpdateUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *cQRSGrpcServerTransports) FindUser() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.FindUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/pb.CQRS/FindUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type cQRSGrpcServer struct {
	createUser grpc.Handler
	deleteUser grpc.Handler
	updateUser grpc.Handler
	findUser   grpc.Handler
}

func (s *cQRSGrpcServer) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.createUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *cQRSGrpcServer) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	ctx, rep, err := s.deleteUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*DeleteUserResponse), nil
}

func (s *cQRSGrpcServer) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	ctx, rep, err := s.updateUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UpdateUserResponse), nil
}

func (s *cQRSGrpcServer) FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error) {
	ctx, rep, err := s.findUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*GetUserResponse), nil
}

type cQRSGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *cQRSGrpcClientTransports) CreateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"pb.CQRS",
		"CreateUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *cQRSGrpcClientTransports) DeleteUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"pb.CQRS",
		"DeleteUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		DeleteUserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *cQRSGrpcClientTransports) UpdateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"pb.CQRS",
		"UpdateUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		UpdateUserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *cQRSGrpcClientTransports) FindUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"pb.CQRS",
		"FindUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		GetUserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newCQRSGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) CQRSClientTransports {
	return &cQRSGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}
