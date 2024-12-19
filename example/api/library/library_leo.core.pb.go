// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package library

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	log "github.com/go-kit/log"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	transportx "github.com/go-leo/leo/v3/transportx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

type LibraryServiceService interface {
	CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error)
	GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error)
	ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error)
	DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error)
	MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error)
	CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error)
	GetBook(ctx context.Context, request *GetBookRequest) (*Book, error)
	ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error)
	DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error)
	UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error)
	MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error)
}

type LibraryServiceEndpoints interface {
	CreateShelf(ctx context.Context) endpoint.Endpoint
	GetShelf(ctx context.Context) endpoint.Endpoint
	ListShelves(ctx context.Context) endpoint.Endpoint
	DeleteShelf(ctx context.Context) endpoint.Endpoint
	MergeShelves(ctx context.Context) endpoint.Endpoint
	CreateBook(ctx context.Context) endpoint.Endpoint
	GetBook(ctx context.Context) endpoint.Endpoint
	ListBooks(ctx context.Context) endpoint.Endpoint
	DeleteBook(ctx context.Context) endpoint.Endpoint
	UpdateBook(ctx context.Context) endpoint.Endpoint
	MoveBook(ctx context.Context) endpoint.Endpoint
}

type LibraryServiceClientTransports interface {
	CreateShelf() transportx.ClientTransport
	GetShelf() transportx.ClientTransport
	ListShelves() transportx.ClientTransport
	DeleteShelf() transportx.ClientTransport
	MergeShelves() transportx.ClientTransport
	CreateBook() transportx.ClientTransport
	GetBook() transportx.ClientTransport
	ListBooks() transportx.ClientTransport
	DeleteBook() transportx.ClientTransport
	UpdateBook() transportx.ClientTransport
	MoveBook() transportx.ClientTransport
}
type LibraryServiceClientTransportsV2 interface {
	CreateShelf(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	GetShelf(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	ListShelves(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	DeleteShelf(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	MergeShelves(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	CreateBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	GetBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	ListBooks(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	DeleteBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	UpdateBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	MoveBook(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type LibraryServiceFactories interface {
	CreateShelf(ctx context.Context) sd.Factory
	GetShelf(ctx context.Context) sd.Factory
	ListShelves(ctx context.Context) sd.Factory
	DeleteShelf(ctx context.Context) sd.Factory
	MergeShelves(ctx context.Context) sd.Factory
	CreateBook(ctx context.Context) sd.Factory
	GetBook(ctx context.Context) sd.Factory
	ListBooks(ctx context.Context) sd.Factory
	DeleteBook(ctx context.Context) sd.Factory
	UpdateBook(ctx context.Context) sd.Factory
	MoveBook(ctx context.Context) sd.Factory
}

type LibraryServiceEndpointers interface {
	CreateShelf(ctx context.Context) sd.Endpointer
	GetShelf(ctx context.Context) sd.Endpointer
	ListShelves(ctx context.Context) sd.Endpointer
	DeleteShelf(ctx context.Context) sd.Endpointer
	MergeShelves(ctx context.Context) sd.Endpointer
	CreateBook(ctx context.Context) sd.Endpointer
	GetBook(ctx context.Context) sd.Endpointer
	ListBooks(ctx context.Context) sd.Endpointer
	DeleteBook(ctx context.Context) sd.Endpointer
	UpdateBook(ctx context.Context) sd.Endpointer
	MoveBook(ctx context.Context) sd.Endpointer
}

type libraryServiceServerEndpoints struct {
	svc         LibraryServiceService
	middlewares []endpoint.Middleware
}

func (e *libraryServiceServerEndpoints) CreateShelf(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateShelf(ctx, request.(*CreateShelfRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) GetShelf(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetShelf(ctx, request.(*GetShelfRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) ListShelves(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.ListShelves(ctx, request.(*ListShelvesRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) DeleteShelf(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteShelf(ctx, request.(*DeleteShelfRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) MergeShelves(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.MergeShelves(ctx, request.(*MergeShelvesRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) CreateBook(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateBook(ctx, request.(*CreateBookRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) GetBook(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetBook(ctx, request.(*GetBookRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) ListBooks(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.ListBooks(ctx, request.(*ListBooksRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) DeleteBook(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteBook(ctx, request.(*DeleteBookRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) UpdateBook(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateBook(ctx, request.(*UpdateBookRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *libraryServiceServerEndpoints) MoveBook(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.MoveBook(ctx, request.(*MoveBookRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newLibraryServiceServerEndpoints(svc LibraryServiceService, middlewares ...endpoint.Middleware) LibraryServiceEndpoints {
	return &libraryServiceServerEndpoints{svc: svc, middlewares: middlewares}
}

type libraryServiceClientEndpoints struct {
	transports  LibraryServiceClientTransports
	middlewares []endpoint.Middleware
}

func (e *libraryServiceClientEndpoints) CreateShelf(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.CreateShelf().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) GetShelf(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.GetShelf().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) ListShelves(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.ListShelves().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) DeleteShelf(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.DeleteShelf().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) MergeShelves(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.MergeShelves().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) CreateBook(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.CreateBook().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) GetBook(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.GetBook().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) ListBooks(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.ListBooks().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) DeleteBook(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.DeleteBook().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) UpdateBook(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.UpdateBook().Endpoint(ctx), e.middlewares...)
}
func (e *libraryServiceClientEndpoints) MoveBook(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.MoveBook().Endpoint(ctx), e.middlewares...)
}
func newLibraryServiceClientEndpoints(transports LibraryServiceClientTransports, middlewares ...endpoint.Middleware) LibraryServiceEndpoints {
	return &libraryServiceClientEndpoints{transports: transports, middlewares: middlewares}
}

type libraryServiceFactories struct {
	transports LibraryServiceClientTransportsV2
}

func (f *libraryServiceFactories) CreateShelf(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.CreateShelf(ctx, instance)
	}
}
func (f *libraryServiceFactories) GetShelf(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.GetShelf(ctx, instance)
	}
}
func (f *libraryServiceFactories) ListShelves(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.ListShelves(ctx, instance)
	}
}
func (f *libraryServiceFactories) DeleteShelf(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.DeleteShelf(ctx, instance)
	}
}
func (f *libraryServiceFactories) MergeShelves(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.MergeShelves(ctx, instance)
	}
}
func (f *libraryServiceFactories) CreateBook(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.CreateBook(ctx, instance)
	}
}
func (f *libraryServiceFactories) GetBook(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.GetBook(ctx, instance)
	}
}
func (f *libraryServiceFactories) ListBooks(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.ListBooks(ctx, instance)
	}
}
func (f *libraryServiceFactories) DeleteBook(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.DeleteBook(ctx, instance)
	}
}
func (f *libraryServiceFactories) UpdateBook(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.UpdateBook(ctx, instance)
	}
}
func (f *libraryServiceFactories) MoveBook(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.MoveBook(ctx, instance)
	}
}
func newLibraryServiceFactories(transports LibraryServiceClientTransportsV2) LibraryServiceFactories {
	return &libraryServiceFactories{transports: transports}
}

type libraryServiceEndpointers struct {
	instancer sd.Instancer
	factories LibraryServiceFactories
	logger    log.Logger
	options   []sd.EndpointerOption
}

func (e *libraryServiceEndpointers) CreateShelf(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.CreateShelf(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) GetShelf(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.GetShelf(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) ListShelves(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.ListShelves(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) DeleteShelf(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.DeleteShelf(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) MergeShelves(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.MergeShelves(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) CreateBook(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.CreateBook(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) GetBook(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.GetBook(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) ListBooks(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.ListBooks(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) DeleteBook(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.DeleteBook(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) UpdateBook(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.UpdateBook(ctx), e.logger, e.options...)
}
func (e *libraryServiceEndpointers) MoveBook(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.MoveBook(ctx), e.logger, e.options...)
}
func newLibraryServiceEndpointers(instancer sd.Instancer, factories LibraryServiceFactories, logger log.Logger, options ...sd.EndpointerOption) LibraryServiceEndpointers {
	return &libraryServiceEndpointers{instancer: instancer, factories: factories, logger: logger, options: options}
}
