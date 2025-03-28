// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package library

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	metadatax "github.com/go-leo/leo/v3/metadatax"
	stainx "github.com/go-leo/leo/v3/stainx"
	grpctransportx "github.com/go-leo/leo/v3/transportx/grpctransportx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

func NewLibraryServiceGrpcServer(svc LibraryServiceService, opts ...grpctransportx.ServerOption) LibraryServiceServer {
	options := grpctransportx.NewServerOptions(opts...)
	endpoints := &libraryServiceServerEndpoints{
		svc:         svc,
		middlewares: options.Middlewares(),
	}
	transports := &libraryServiceGrpcServerTransports{
		endpoints: endpoints,
	}
	return &libraryServiceGrpcServer{
		createShelf:  transports.CreateShelf(),
		getShelf:     transports.GetShelf(),
		listShelves:  transports.ListShelves(),
		deleteShelf:  transports.DeleteShelf(),
		mergeShelves: transports.MergeShelves(),
		createBook:   transports.CreateBook(),
		getBook:      transports.GetBook(),
		listBooks:    transports.ListBooks(),
		deleteBook:   transports.DeleteBook(),
		updateBook:   transports.UpdateBook(),
		moveBook:     transports.MoveBook(),
	}
}

func NewLibraryServiceGrpcClient(target string, opts ...grpctransportx.ClientOption) LibraryServiceService {
	options := grpctransportx.NewClientOptions(opts...)
	transports := &libraryServiceGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &libraryServiceFactories{
		transports: transports,
	}
	endpointer := &libraryServiceEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &libraryServiceBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &libraryServiceClientEndpoints{
		balancers: balancers,
	}
	return &libraryServiceClientService{
		endpoints:     endpoints,
		transportName: grpctransportx.GrpcClient,
	}
}

type libraryServiceGrpcServerTransports struct {
	endpoints LibraryServiceServerEndpoints
}

func (t *libraryServiceGrpcServerTransports) CreateShelf() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.CreateShelf(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/CreateShelf")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) GetShelf() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.GetShelf(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/GetShelf")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) ListShelves() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.ListShelves(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/ListShelves")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) DeleteShelf() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.DeleteShelf(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/DeleteShelf")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) MergeShelves() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.MergeShelves(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/MergeShelves")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) CreateBook() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.CreateBook(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/CreateBook")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) GetBook() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.GetBook(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/GetBook")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) ListBooks() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.ListBooks(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/ListBooks")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) DeleteBook() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.DeleteBook(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/DeleteBook")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) UpdateBook() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.UpdateBook(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/UpdateBook")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

func (t *libraryServiceGrpcServerTransports) MoveBook() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.MoveBook(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpctransportx.ServerEndpointInjector("/google.example.library.v1.LibraryService/MoveBook")),
		grpc.ServerBefore(grpctransportx.ServerTransportInjector),
		grpc.ServerBefore(metadatax.GrpcIncomingInjector),
		grpc.ServerBefore(stainx.GrpcIncomingInjector),
	)
}

type libraryServiceGrpcServer struct {
	createShelf  grpc.Handler
	getShelf     grpc.Handler
	listShelves  grpc.Handler
	deleteShelf  grpc.Handler
	mergeShelves grpc.Handler
	createBook   grpc.Handler
	getBook      grpc.Handler
	listBooks    grpc.Handler
	deleteBook   grpc.Handler
	updateBook   grpc.Handler
	moveBook     grpc.Handler
}

func (s *libraryServiceGrpcServer) CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error) {
	ctx, rep, err := s.createShelf.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Shelf), nil
}

func (s *libraryServiceGrpcServer) GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error) {
	ctx, rep, err := s.getShelf.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Shelf), nil
}

func (s *libraryServiceGrpcServer) ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error) {
	ctx, rep, err := s.listShelves.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*ListShelvesResponse), nil
}

func (s *libraryServiceGrpcServer) DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteShelf.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *libraryServiceGrpcServer) MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error) {
	ctx, rep, err := s.mergeShelves.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Shelf), nil
}

func (s *libraryServiceGrpcServer) CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error) {
	ctx, rep, err := s.createBook.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Book), nil
}

func (s *libraryServiceGrpcServer) GetBook(ctx context.Context, request *GetBookRequest) (*Book, error) {
	ctx, rep, err := s.getBook.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Book), nil
}

func (s *libraryServiceGrpcServer) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	ctx, rep, err := s.listBooks.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*ListBooksResponse), nil
}

func (s *libraryServiceGrpcServer) DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteBook.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *libraryServiceGrpcServer) UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error) {
	ctx, rep, err := s.updateBook.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Book), nil
}

func (s *libraryServiceGrpcServer) MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error) {
	ctx, rep, err := s.moveBook.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Book), nil
}

func (s *libraryServiceGrpcServer) mustEmbedUnimplementedLibraryServiceServer() {}

type libraryServiceGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *libraryServiceGrpcClientTransports) CreateShelf(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"CreateShelf",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Shelf{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) GetShelf(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"GetShelf",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Shelf{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) ListShelves(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"ListShelves",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		ListShelvesResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) DeleteShelf(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"DeleteShelf",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) MergeShelves(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"MergeShelves",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Shelf{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) CreateBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"CreateBook",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Book{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) GetBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"GetBook",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Book{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) ListBooks(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"ListBooks",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		ListBooksResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) DeleteBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"DeleteBook",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) UpdateBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"UpdateBook",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Book{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *libraryServiceGrpcClientTransports) MoveBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(metadatax.GrpcOutgoingInjector),
		grpc.ClientBefore(stainx.GrpcOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.library.v1.LibraryService",
		"MoveBook",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Book{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}
