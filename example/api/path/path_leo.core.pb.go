// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package path

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
	stainx "github.com/go-leo/leo/v3/sdx/stainx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

type PathService interface {
	BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
}

type PathServerEndpoints interface {
	BoolPath(ctx context.Context) endpoint.Endpoint
	Int32Path(ctx context.Context) endpoint.Endpoint
	Int64Path(ctx context.Context) endpoint.Endpoint
	Uint32Path(ctx context.Context) endpoint.Endpoint
	Uint64Path(ctx context.Context) endpoint.Endpoint
	FloatPath(ctx context.Context) endpoint.Endpoint
	DoublePath(ctx context.Context) endpoint.Endpoint
	StringPath(ctx context.Context) endpoint.Endpoint
	EnumPath(ctx context.Context) endpoint.Endpoint
}

type PathClientEndpoints interface {
	BoolPath(ctx context.Context) (endpoint.Endpoint, error)
	Int32Path(ctx context.Context) (endpoint.Endpoint, error)
	Int64Path(ctx context.Context) (endpoint.Endpoint, error)
	Uint32Path(ctx context.Context) (endpoint.Endpoint, error)
	Uint64Path(ctx context.Context) (endpoint.Endpoint, error)
	FloatPath(ctx context.Context) (endpoint.Endpoint, error)
	DoublePath(ctx context.Context) (endpoint.Endpoint, error)
	StringPath(ctx context.Context) (endpoint.Endpoint, error)
	EnumPath(ctx context.Context) (endpoint.Endpoint, error)
}

type PathClientTransports interface {
	BoolPath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	Int32Path(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	Int64Path(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	Uint32Path(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	Uint64Path(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	FloatPath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	DoublePath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	StringPath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	EnumPath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type PathFactories interface {
	BoolPath(ctx context.Context) sd.Factory
	Int32Path(ctx context.Context) sd.Factory
	Int64Path(ctx context.Context) sd.Factory
	Uint32Path(ctx context.Context) sd.Factory
	Uint64Path(ctx context.Context) sd.Factory
	FloatPath(ctx context.Context) sd.Factory
	DoublePath(ctx context.Context) sd.Factory
	StringPath(ctx context.Context) sd.Factory
	EnumPath(ctx context.Context) sd.Factory
}

type PathEndpointers interface {
	BoolPath(ctx context.Context, color string) (sd.Endpointer, error)
	Int32Path(ctx context.Context, color string) (sd.Endpointer, error)
	Int64Path(ctx context.Context, color string) (sd.Endpointer, error)
	Uint32Path(ctx context.Context, color string) (sd.Endpointer, error)
	Uint64Path(ctx context.Context, color string) (sd.Endpointer, error)
	FloatPath(ctx context.Context, color string) (sd.Endpointer, error)
	DoublePath(ctx context.Context, color string) (sd.Endpointer, error)
	StringPath(ctx context.Context, color string) (sd.Endpointer, error)
	EnumPath(ctx context.Context, color string) (sd.Endpointer, error)
}

type PathBalancers interface {
	BoolPath(ctx context.Context) (lb.Balancer, error)
	Int32Path(ctx context.Context) (lb.Balancer, error)
	Int64Path(ctx context.Context) (lb.Balancer, error)
	Uint32Path(ctx context.Context) (lb.Balancer, error)
	Uint64Path(ctx context.Context) (lb.Balancer, error)
	FloatPath(ctx context.Context) (lb.Balancer, error)
	DoublePath(ctx context.Context) (lb.Balancer, error)
	StringPath(ctx context.Context) (lb.Balancer, error)
	EnumPath(ctx context.Context) (lb.Balancer, error)
}

type pathServerEndpoints struct {
	svc         PathService
	middlewares []endpoint.Middleware
}

func (e *pathServerEndpoints) BoolPath(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.BoolPath(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) Int32Path(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.Int32Path(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) Int64Path(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.Int64Path(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) Uint32Path(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.Uint32Path(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) Uint64Path(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.Uint64Path(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) FloatPath(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.FloatPath(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) DoublePath(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DoublePath(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) StringPath(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StringPath(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *pathServerEndpoints) EnumPath(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.EnumPath(ctx, request.(*PathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newPathServerEndpoints(svc PathService, middlewares ...endpoint.Middleware) PathServerEndpoints {
	return &pathServerEndpoints{svc: svc, middlewares: middlewares}
}

type pathClientEndpoints struct {
	balancers PathBalancers
}

func (e *pathClientEndpoints) BoolPath(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.BoolPath(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) Int32Path(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.Int32Path(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) Int64Path(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.Int64Path(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) Uint32Path(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.Uint32Path(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) Uint64Path(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.Uint64Path(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) FloatPath(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.FloatPath(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) DoublePath(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.DoublePath(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) StringPath(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.StringPath(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *pathClientEndpoints) EnumPath(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.EnumPath(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func newPathClientEndpoints(
	target string,
	transports PathClientTransports,
	instancerFactory sdx.InstancerFactory,
	endpointerOptions []sd.EndpointerOption,
	balancerFactory lbx.BalancerFactory,
	logger log.Logger,
) PathClientEndpoints {
	factories := newPathFactories(transports)
	endpointers := newPathEndpointers(target, instancerFactory, factories, logger, endpointerOptions...)
	balancers := newPathBalancers(balancerFactory, endpointers)
	return &pathClientEndpoints{balancers: balancers}
}

type pathFactories struct {
	transports PathClientTransports
}

func (f *pathFactories) BoolPath(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.BoolPath(ctx, instance)
	}
}
func (f *pathFactories) Int32Path(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.Int32Path(ctx, instance)
	}
}
func (f *pathFactories) Int64Path(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.Int64Path(ctx, instance)
	}
}
func (f *pathFactories) Uint32Path(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.Uint32Path(ctx, instance)
	}
}
func (f *pathFactories) Uint64Path(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.Uint64Path(ctx, instance)
	}
}
func (f *pathFactories) FloatPath(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.FloatPath(ctx, instance)
	}
}
func (f *pathFactories) DoublePath(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.DoublePath(ctx, instance)
	}
}
func (f *pathFactories) StringPath(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.StringPath(ctx, instance)
	}
}
func (f *pathFactories) EnumPath(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.EnumPath(ctx, instance)
	}
}
func newPathFactories(transports PathClientTransports) PathFactories {
	return &pathFactories{transports: transports}
}

type pathEndpointers struct {
	target           string
	instancerFactory sdx.InstancerFactory
	factories        PathFactories
	logger           log.Logger
	options          []sd.EndpointerOption
}

func (e *pathEndpointers) BoolPath(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.BoolPath(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) Int32Path(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.Int32Path(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) Int64Path(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.Int64Path(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) Uint32Path(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.Uint32Path(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) Uint64Path(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.Uint64Path(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) FloatPath(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.FloatPath(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) DoublePath(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.DoublePath(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) StringPath(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.StringPath(ctx), e.logger, e.options...)
}
func (e *pathEndpointers) EnumPath(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.EnumPath(ctx), e.logger, e.options...)
}
func newPathEndpointers(
	target string,
	instancerFactory sdx.InstancerFactory,
	factories PathFactories,
	logger log.Logger,
	options ...sd.EndpointerOption,
) PathEndpointers {
	return &pathEndpointers{
		target:           target,
		instancerFactory: instancerFactory,
		factories:        factories,
		logger:           logger,
		options:          options,
	}
}

type pathBalancers struct {
	factory    lbx.BalancerFactory
	endpointer PathEndpointers
	boolPath   lazyloadx.Group[lb.Balancer]
	int32Path  lazyloadx.Group[lb.Balancer]
	int64Path  lazyloadx.Group[lb.Balancer]
	uint32Path lazyloadx.Group[lb.Balancer]
	uint64Path lazyloadx.Group[lb.Balancer]
	floatPath  lazyloadx.Group[lb.Balancer]
	doublePath lazyloadx.Group[lb.Balancer]
	stringPath lazyloadx.Group[lb.Balancer]
	enumPath   lazyloadx.Group[lb.Balancer]
}

func (b *pathBalancers) BoolPath(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.boolPath.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.BoolPath))
	return balancer, err
}
func (b *pathBalancers) Int32Path(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.int32Path.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.Int32Path))
	return balancer, err
}
func (b *pathBalancers) Int64Path(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.int64Path.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.Int64Path))
	return balancer, err
}
func (b *pathBalancers) Uint32Path(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.uint32Path.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.Uint32Path))
	return balancer, err
}
func (b *pathBalancers) Uint64Path(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.uint64Path.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.Uint64Path))
	return balancer, err
}
func (b *pathBalancers) FloatPath(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.floatPath.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.FloatPath))
	return balancer, err
}
func (b *pathBalancers) DoublePath(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.doublePath.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.DoublePath))
	return balancer, err
}
func (b *pathBalancers) StringPath(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.stringPath.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.StringPath))
	return balancer, err
}
func (b *pathBalancers) EnumPath(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.enumPath.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.EnumPath))
	return balancer, err
}
func newPathBalancers(factory lbx.BalancerFactory, endpointer PathEndpointers) PathBalancers {
	return &pathBalancers{
		factory:    factory,
		endpointer: endpointer,
		boolPath:   lazyloadx.Group[lb.Balancer]{},
		int32Path:  lazyloadx.Group[lb.Balancer]{},
		int64Path:  lazyloadx.Group[lb.Balancer]{},
		uint32Path: lazyloadx.Group[lb.Balancer]{},
		uint64Path: lazyloadx.Group[lb.Balancer]{},
		floatPath:  lazyloadx.Group[lb.Balancer]{},
		doublePath: lazyloadx.Group[lb.Balancer]{},
		stringPath: lazyloadx.Group[lb.Balancer]{},
		enumPath:   lazyloadx.Group[lb.Balancer]{},
	}
}

type pathClientService struct {
	endpoints     PathClientEndpoints
	transportName string
}

func (c *pathClientService) BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/BoolPath")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.BoolPath(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Int32Path")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.Int32Path(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Int64Path")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.Int64Path(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Uint32Path")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.Uint32Path(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Uint64Path")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.Uint64Path(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/FloatPath")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.FloatPath(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/DoublePath")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.DoublePath(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/StringPath")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.StringPath(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func (c *pathClientService) EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/EnumPath")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.EnumPath(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func newPathClientService(endpoints PathClientEndpoints, transportName string) PathService {
	return &pathClientService{endpoints: endpoints, transportName: transportName}
}
