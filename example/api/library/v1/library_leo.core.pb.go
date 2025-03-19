// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package library

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	lb "github.com/go-kit/kit/sd/lb"
	log "github.com/go-kit/log"
	lazyloadx "github.com/go-leo/gox/syncx/lazyloadx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	sdx "github.com/go-leo/leo/v3/sdx"
	lbx "github.com/go-leo/leo/v3/sdx/lbx"
	stainx "github.com/go-leo/leo/v3/stainx"
	transportx "github.com/go-leo/leo/v3/transportx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// LibraryServiceService is a service
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

type (
	LibraryServiceService_CreateShelf_Handler interface {
		Handle(ctx context.Context, request *CreateShelfRequest) (*Shelf, error)
	}
	LibraryServiceService_GetShelf_Handler interface {
		Handle(ctx context.Context, request *GetShelfRequest) (*Shelf, error)
	}
	LibraryServiceService_ListShelves_Handler interface {
		Handle(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error)
	}
	LibraryServiceService_DeleteShelf_Handler interface {
		Handle(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error)
	}
	LibraryServiceService_MergeShelves_Handler interface {
		Handle(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error)
	}
	LibraryServiceService_CreateBook_Handler interface {
		Handle(ctx context.Context, request *CreateBookRequest) (*Book, error)
	}
	LibraryServiceService_GetBook_Handler interface {
		Handle(ctx context.Context, request *GetBookRequest) (*Book, error)
	}
	LibraryServiceService_ListBooks_Handler interface {
		Handle(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error)
	}
	LibraryServiceService_DeleteBook_Handler interface {
		Handle(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error)
	}
	LibraryServiceService_UpdateBook_Handler interface {
		Handle(ctx context.Context, request *UpdateBookRequest) (*Book, error)
	}
	LibraryServiceService_MoveBook_Handler interface {
		Handle(ctx context.Context, request *MoveBookRequest) (*Book, error)
	}
)

// LibraryServiceServerEndpoints is server endpoints
type LibraryServiceServerEndpoints interface {
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

// LibraryServiceClientEndpoints is client endpoints
type LibraryServiceClientEndpoints interface {
	CreateShelf(ctx context.Context) (endpoint.Endpoint, error)
	GetShelf(ctx context.Context) (endpoint.Endpoint, error)
	ListShelves(ctx context.Context) (endpoint.Endpoint, error)
	DeleteShelf(ctx context.Context) (endpoint.Endpoint, error)
	MergeShelves(ctx context.Context) (endpoint.Endpoint, error)
	CreateBook(ctx context.Context) (endpoint.Endpoint, error)
	GetBook(ctx context.Context) (endpoint.Endpoint, error)
	ListBooks(ctx context.Context) (endpoint.Endpoint, error)
	DeleteBook(ctx context.Context) (endpoint.Endpoint, error)
	UpdateBook(ctx context.Context) (endpoint.Endpoint, error)
	MoveBook(ctx context.Context) (endpoint.Endpoint, error)
}

// LibraryServiceClientTransports is client transports
type LibraryServiceClientTransports interface {
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

// LibraryServiceFactories is client factories
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

// LibraryServiceEndpointers is client endpointers
type LibraryServiceEndpointers interface {
	CreateShelf(ctx context.Context, color string) (sd.Endpointer, error)
	GetShelf(ctx context.Context, color string) (sd.Endpointer, error)
	ListShelves(ctx context.Context, color string) (sd.Endpointer, error)
	DeleteShelf(ctx context.Context, color string) (sd.Endpointer, error)
	MergeShelves(ctx context.Context, color string) (sd.Endpointer, error)
	CreateBook(ctx context.Context, color string) (sd.Endpointer, error)
	GetBook(ctx context.Context, color string) (sd.Endpointer, error)
	ListBooks(ctx context.Context, color string) (sd.Endpointer, error)
	DeleteBook(ctx context.Context, color string) (sd.Endpointer, error)
	UpdateBook(ctx context.Context, color string) (sd.Endpointer, error)
	MoveBook(ctx context.Context, color string) (sd.Endpointer, error)
}

// LibraryServiceBalancers is client balancers
type LibraryServiceBalancers interface {
	CreateShelf(ctx context.Context) (lb.Balancer, error)
	GetShelf(ctx context.Context) (lb.Balancer, error)
	ListShelves(ctx context.Context) (lb.Balancer, error)
	DeleteShelf(ctx context.Context) (lb.Balancer, error)
	MergeShelves(ctx context.Context) (lb.Balancer, error)
	CreateBook(ctx context.Context) (lb.Balancer, error)
	GetBook(ctx context.Context) (lb.Balancer, error)
	ListBooks(ctx context.Context) (lb.Balancer, error)
	DeleteBook(ctx context.Context) (lb.Balancer, error)
	UpdateBook(ctx context.Context) (lb.Balancer, error)
	MoveBook(ctx context.Context) (lb.Balancer, error)
}

// libraryServiceServerEndpoints implements LibraryServiceServerEndpoints
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

// libraryServiceFactories implements LibraryServiceFactories
type libraryServiceFactories struct {
	transports LibraryServiceClientTransports
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

// libraryServiceEndpointers implements LibraryServiceEndpointers
type libraryServiceEndpointers struct {
	target    string
	builder   sdx.Builder
	factories LibraryServiceFactories
	logger    log.Logger
	options   []sd.EndpointerOption
}

func (e *libraryServiceEndpointers) CreateShelf(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.CreateShelf(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) GetShelf(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.GetShelf(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) ListShelves(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.ListShelves(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) DeleteShelf(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.DeleteShelf(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) MergeShelves(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.MergeShelves(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) CreateBook(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.CreateBook(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) GetBook(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.GetBook(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) ListBooks(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.ListBooks(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) DeleteBook(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.DeleteBook(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) UpdateBook(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.UpdateBook(ctx), e.logger, e.options...)
}

func (e *libraryServiceEndpointers) MoveBook(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.MoveBook(ctx), e.logger, e.options...)
}

// libraryServiceBalancers implements LibraryServiceBalancers
type libraryServiceBalancers struct {
	factory      lbx.BalancerFactory
	endpointer   LibraryServiceEndpointers
	createShelf  lazyloadx.Group[lb.Balancer]
	getShelf     lazyloadx.Group[lb.Balancer]
	listShelves  lazyloadx.Group[lb.Balancer]
	deleteShelf  lazyloadx.Group[lb.Balancer]
	mergeShelves lazyloadx.Group[lb.Balancer]
	createBook   lazyloadx.Group[lb.Balancer]
	getBook      lazyloadx.Group[lb.Balancer]
	listBooks    lazyloadx.Group[lb.Balancer]
	deleteBook   lazyloadx.Group[lb.Balancer]
	updateBook   lazyloadx.Group[lb.Balancer]
	moveBook     lazyloadx.Group[lb.Balancer]
}

func (b *libraryServiceBalancers) CreateShelf(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.createShelf.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.CreateShelf))
	return balancer, err
}
func (b *libraryServiceBalancers) GetShelf(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.getShelf.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.GetShelf))
	return balancer, err
}
func (b *libraryServiceBalancers) ListShelves(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.listShelves.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.ListShelves))
	return balancer, err
}
func (b *libraryServiceBalancers) DeleteShelf(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.deleteShelf.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.DeleteShelf))
	return balancer, err
}
func (b *libraryServiceBalancers) MergeShelves(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.mergeShelves.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.MergeShelves))
	return balancer, err
}
func (b *libraryServiceBalancers) CreateBook(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.createBook.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.CreateBook))
	return balancer, err
}
func (b *libraryServiceBalancers) GetBook(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.getBook.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.GetBook))
	return balancer, err
}
func (b *libraryServiceBalancers) ListBooks(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.listBooks.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.ListBooks))
	return balancer, err
}
func (b *libraryServiceBalancers) DeleteBook(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.deleteBook.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.DeleteBook))
	return balancer, err
}
func (b *libraryServiceBalancers) UpdateBook(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.updateBook.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.UpdateBook))
	return balancer, err
}
func (b *libraryServiceBalancers) MoveBook(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.moveBook.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.MoveBook))
	return balancer, err
}
func newLibraryServiceBalancers(factory lbx.BalancerFactory, endpointer LibraryServiceEndpointers) LibraryServiceBalancers {
	return &libraryServiceBalancers{
		factory:      factory,
		endpointer:   endpointer,
		createShelf:  lazyloadx.Group[lb.Balancer]{},
		getShelf:     lazyloadx.Group[lb.Balancer]{},
		listShelves:  lazyloadx.Group[lb.Balancer]{},
		deleteShelf:  lazyloadx.Group[lb.Balancer]{},
		mergeShelves: lazyloadx.Group[lb.Balancer]{},
		createBook:   lazyloadx.Group[lb.Balancer]{},
		getBook:      lazyloadx.Group[lb.Balancer]{},
		listBooks:    lazyloadx.Group[lb.Balancer]{},
		deleteBook:   lazyloadx.Group[lb.Balancer]{},
		updateBook:   lazyloadx.Group[lb.Balancer]{},
		moveBook:     lazyloadx.Group[lb.Balancer]{},
	}
}

// libraryServiceClientEndpoints implements LibraryServiceClientEndpoints
type libraryServiceClientEndpoints struct {
	balancers LibraryServiceBalancers
}

func (e *libraryServiceClientEndpoints) CreateShelf(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.CreateShelf(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) GetShelf(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.GetShelf(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) ListShelves(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.ListShelves(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) DeleteShelf(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.DeleteShelf(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) MergeShelves(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.MergeShelves(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) CreateBook(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.CreateBook(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) GetBook(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.GetBook(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) ListBooks(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.ListBooks(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) DeleteBook(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.DeleteBook(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) UpdateBook(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.UpdateBook(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *libraryServiceClientEndpoints) MoveBook(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.MoveBook(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

// libraryServiceClientService implements LibraryServiceClientService
type libraryServiceClientService struct {
	endpoints     LibraryServiceClientEndpoints
	transportName string
}

func (c *libraryServiceClientService) CreateShelf(ctx context.Context, request *CreateShelfRequest) (*Shelf, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/CreateShelf")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.CreateShelf(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *libraryServiceClientService) GetShelf(ctx context.Context, request *GetShelfRequest) (*Shelf, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/GetShelf")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.GetShelf(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *libraryServiceClientService) ListShelves(ctx context.Context, request *ListShelvesRequest) (*ListShelvesResponse, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/ListShelves")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.ListShelves(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListShelvesResponse), nil
}

func (c *libraryServiceClientService) DeleteShelf(ctx context.Context, request *DeleteShelfRequest) (*emptypb.Empty, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/DeleteShelf")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.DeleteShelf(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *libraryServiceClientService) MergeShelves(ctx context.Context, request *MergeShelvesRequest) (*Shelf, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/MergeShelves")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.MergeShelves(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Shelf), nil
}

func (c *libraryServiceClientService) CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/CreateBook")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.CreateBook(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *libraryServiceClientService) GetBook(ctx context.Context, request *GetBookRequest) (*Book, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/GetBook")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.GetBook(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *libraryServiceClientService) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/ListBooks")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.ListBooks(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListBooksResponse), nil
}

func (c *libraryServiceClientService) DeleteBook(ctx context.Context, request *DeleteBookRequest) (*emptypb.Empty, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/DeleteBook")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.DeleteBook(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *libraryServiceClientService) UpdateBook(ctx context.Context, request *UpdateBookRequest) (*Book, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/UpdateBook")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.UpdateBook(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}

func (c *libraryServiceClientService) MoveBook(ctx context.Context, request *MoveBookRequest) (*Book, error) {
	ctx = endpointx.NameInjector(ctx, "/google.example.library.v1.LibraryService/MoveBook")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.MoveBook(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Book), nil
}
