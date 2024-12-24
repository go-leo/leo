// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package demo

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// =========================== grpc server ===========================

type DemoGrpcServerTransports interface {
	CreateUser() *grpc.Server
	DeleteUser() *grpc.Server
	UpdateUser() *grpc.Server
	GetUser() *grpc.Server
	GetUsers() *grpc.Server
	UploadUserAvatar() *grpc.Server
	GetUserAvatar() *grpc.Server
}

type demoGrpcServerTransports struct {
	createUser       *grpc.Server
	deleteUser       *grpc.Server
	updateUser       *grpc.Server
	getUser          *grpc.Server
	getUsers         *grpc.Server
	uploadUserAvatar *grpc.Server
	getUserAvatar    *grpc.Server
}

func (t *demoGrpcServerTransports) CreateUser() *grpc.Server {
	return t.createUser
}

func (t *demoGrpcServerTransports) DeleteUser() *grpc.Server {
	return t.deleteUser
}

func (t *demoGrpcServerTransports) UpdateUser() *grpc.Server {
	return t.updateUser
}

func (t *demoGrpcServerTransports) GetUser() *grpc.Server {
	return t.getUser
}

func (t *demoGrpcServerTransports) GetUsers() *grpc.Server {
	return t.getUsers
}

func (t *demoGrpcServerTransports) UploadUserAvatar() *grpc.Server {
	return t.uploadUserAvatar
}

func (t *demoGrpcServerTransports) GetUserAvatar() *grpc.Server {
	return t.getUserAvatar
}

func newDemoGrpcServerTransports(endpoints DemoEndpoints) DemoGrpcServerTransports {
	return &demoGrpcServerTransports{
		createUser:       _Demo_CreateUser_GrpcServer_Transport(endpoints),
		deleteUser:       _Demo_DeleteUser_GrpcServer_Transport(endpoints),
		updateUser:       _Demo_UpdateUser_GrpcServer_Transport(endpoints),
		getUser:          _Demo_GetUser_GrpcServer_Transport(endpoints),
		getUsers:         _Demo_GetUsers_GrpcServer_Transport(endpoints),
		uploadUserAvatar: _Demo_UploadUserAvatar_GrpcServer_Transport(endpoints),
		getUserAvatar:    _Demo_GetUserAvatar_GrpcServer_Transport(endpoints),
	}
}

type demoGrpcServer struct {
	createUser       *grpc.Server
	deleteUser       *grpc.Server
	updateUser       *grpc.Server
	getUser          *grpc.Server
	getUsers         *grpc.Server
	uploadUserAvatar *grpc.Server
	getUserAvatar    *grpc.Server
}

func (s *demoGrpcServer) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, rep, err := s.createUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*CreateUserResponse), nil
}

func (s *demoGrpcServer) DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *demoGrpcServer) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.updateUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *demoGrpcServer) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	ctx, rep, err := s.getUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*GetUserResponse), nil
}

func (s *demoGrpcServer) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	ctx, rep, err := s.getUsers.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*GetUsersResponse), nil
}

func (s *demoGrpcServer) UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uploadUserAvatar.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *demoGrpcServer) GetUserAvatar(ctx context.Context, request *GetUserAvatarRequest) (*httpbody.HttpBody, error) {
	ctx, rep, err := s.getUserAvatar.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*httpbody.HttpBody), nil
}

func NewDemoGrpcServer(svc DemoService, middlewares ...endpoint.Middleware) DemoService {
	endpoints := newDemoServerEndpoints(svc, middlewares...)
	transports := newDemoGrpcServerTransports(endpoints)
	return &demoGrpcServer{
		createUser:       transports.CreateUser(),
		deleteUser:       transports.DeleteUser(),
		updateUser:       transports.UpdateUser(),
		getUser:          transports.GetUser(),
		getUsers:         transports.GetUsers(),
		uploadUserAvatar: transports.UploadUserAvatar(),
		getUserAvatar:    transports.GetUserAvatar(),
	}
}

// =========================== grpc client ===========================

type demoGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *demoGrpcClientTransports) CreateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"CreateUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		CreateUserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *demoGrpcClientTransports) DeleteUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"DeleteUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *demoGrpcClientTransports) UpdateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"UpdateUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *demoGrpcClientTransports) GetUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"GetUser",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		GetUserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *demoGrpcClientTransports) GetUsers(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"GetUsers",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		GetUsersResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *demoGrpcClientTransports) UploadUserAvatar(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"UploadUserAvatar",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *demoGrpcClientTransports) GetUserAvatar(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.demo.v1.Demo",
		"GetUserAvatar",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		httpbody.HttpBody{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newDemoGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) DemoClientTransports {
	return &demoGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

type demoGrpcClient struct {
	balancers DemoBalancers
}

func (c *demoGrpcClient) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/CreateUser")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*CreateUserResponse), nil
}

func (c *demoGrpcClient) DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/DeleteUser")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.DeleteUser(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *demoGrpcClient) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/UpdateUser")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.UpdateUser(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *demoGrpcClient) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/GetUser")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*GetUserResponse), nil
}

func (c *demoGrpcClient) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/GetUsers")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*GetUsersResponse), nil
}

func (c *demoGrpcClient) UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/UploadUserAvatar")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.UploadUserAvatar(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *demoGrpcClient) GetUserAvatar(ctx context.Context, request *GetUserAvatarRequest) (*httpbody.HttpBody, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.demo.v1.Demo/GetUserAvatar")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	balancer, err := c.balancers.GetUserAvatar(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*httpbody.HttpBody), nil
}

func NewDemoGrpcClient(target string, opts ...grpcx.ClientOption) DemoService {
	options := grpcx.NewClientOptions(opts...)
	transports := newDemoGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	factories := newDemoFactories(transports)
	endpointers := newDemoEndpointers(target, options.InstancerFactory(), factories, options.Logger(), options.EndpointerOptions()...)
	balancers := newDemoBalancers(options.BalancerFactory(), endpointers)
	return &demoHttpClient{balancers: balancers}
}

// =========================== grpc transport ===========================

func _Demo_CreateUser_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.CreateUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/CreateUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Demo_DeleteUser_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.DeleteUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/DeleteUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Demo_UpdateUser_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.UpdateUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/UpdateUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Demo_GetUser_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.GetUser(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/GetUser")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Demo_GetUsers_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.GetUsers(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/GetUsers")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Demo_UploadUserAvatar_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.UploadUserAvatar(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/UploadUserAvatar")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Demo_GetUserAvatar_GrpcServer_Transport(endpoints DemoEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.GetUserAvatar(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.demo.v1.Demo/GetUserAvatar")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}
