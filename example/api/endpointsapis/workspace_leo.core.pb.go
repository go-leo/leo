// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package endpointsapis

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
	stain "github.com/go-leo/leo/v3/sdx/stain"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// WorkspacesService is a service
type WorkspacesService interface {
	ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error)
	CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error)
}

// WorkspacesServerEndpoints is server endpoints
type WorkspacesServerEndpoints interface {
	ListWorkspaces(ctx context.Context) endpoint.Endpoint
	GetWorkspace(ctx context.Context) endpoint.Endpoint
	CreateWorkspace(ctx context.Context) endpoint.Endpoint
	UpdateWorkspace(ctx context.Context) endpoint.Endpoint
	DeleteWorkspace(ctx context.Context) endpoint.Endpoint
}

// WorkspacesClientEndpoints is client endpoints
type WorkspacesClientEndpoints interface {
	ListWorkspaces(ctx context.Context) (endpoint.Endpoint, error)
	GetWorkspace(ctx context.Context) (endpoint.Endpoint, error)
	CreateWorkspace(ctx context.Context) (endpoint.Endpoint, error)
	UpdateWorkspace(ctx context.Context) (endpoint.Endpoint, error)
	DeleteWorkspace(ctx context.Context) (endpoint.Endpoint, error)
}

// WorkspacesClientTransports is client transports
type WorkspacesClientTransports interface {
	ListWorkspaces(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	GetWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	CreateWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	UpdateWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	DeleteWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

// WorkspacesFactories is client factories
type WorkspacesFactories interface {
	ListWorkspaces(ctx context.Context) sd.Factory
	GetWorkspace(ctx context.Context) sd.Factory
	CreateWorkspace(ctx context.Context) sd.Factory
	UpdateWorkspace(ctx context.Context) sd.Factory
	DeleteWorkspace(ctx context.Context) sd.Factory
}

// WorkspacesEndpointers is client endpointers
type WorkspacesEndpointers interface {
	ListWorkspaces(ctx context.Context, color string) (sd.Endpointer, error)
	GetWorkspace(ctx context.Context, color string) (sd.Endpointer, error)
	CreateWorkspace(ctx context.Context, color string) (sd.Endpointer, error)
	UpdateWorkspace(ctx context.Context, color string) (sd.Endpointer, error)
	DeleteWorkspace(ctx context.Context, color string) (sd.Endpointer, error)
}

// WorkspacesBalancers is client balancers
type WorkspacesBalancers interface {
	ListWorkspaces(ctx context.Context) (lb.Balancer, error)
	GetWorkspace(ctx context.Context) (lb.Balancer, error)
	CreateWorkspace(ctx context.Context) (lb.Balancer, error)
	UpdateWorkspace(ctx context.Context) (lb.Balancer, error)
	DeleteWorkspace(ctx context.Context) (lb.Balancer, error)
}

// workspacesServerEndpoints implements WorkspacesServerEndpoints
type workspacesServerEndpoints struct {
	svc         WorkspacesService
	middlewares []endpoint.Middleware
}

func (e *workspacesServerEndpoints) ListWorkspaces(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.ListWorkspaces(ctx, request.(*ListWorkspacesRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *workspacesServerEndpoints) GetWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetWorkspace(ctx, request.(*GetWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *workspacesServerEndpoints) CreateWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateWorkspace(ctx, request.(*CreateWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *workspacesServerEndpoints) UpdateWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateWorkspace(ctx, request.(*UpdateWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *workspacesServerEndpoints) DeleteWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteWorkspace(ctx, request.(*DeleteWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newWorkspacesServerEndpoints(svc WorkspacesService, middlewares ...endpoint.Middleware) WorkspacesServerEndpoints {
	return &workspacesServerEndpoints{svc: svc, middlewares: middlewares}
}

// workspacesClientEndpoints implements WorkspacesClientEndpoints
type workspacesClientEndpoints struct {
	balancers WorkspacesBalancers
}

func (e *workspacesClientEndpoints) ListWorkspaces(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.ListWorkspaces(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *workspacesClientEndpoints) GetWorkspace(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.GetWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *workspacesClientEndpoints) CreateWorkspace(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.CreateWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *workspacesClientEndpoints) UpdateWorkspace(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.UpdateWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *workspacesClientEndpoints) DeleteWorkspace(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.DeleteWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func newWorkspacesClientEndpoints(
	target string,
	transports WorkspacesClientTransports,
	builder sdx.Builder,
	endpointerOptions []sd.EndpointerOption,
	balancerFactory lbx.BalancerFactory,
	logger log.Logger,
) WorkspacesClientEndpoints {
	factories := newWorkspacesFactories(transports)
	endpointers := newWorkspacesEndpointers(target, builder, factories, logger, endpointerOptions...)
	balancers := newWorkspacesBalancers(balancerFactory, endpointers)
	return &workspacesClientEndpoints{balancers: balancers}
}

// workspacesFactories implements WorkspacesFactories
type workspacesFactories struct {
	transports WorkspacesClientTransports
}

func (f *workspacesFactories) ListWorkspaces(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.ListWorkspaces(ctx, instance)
	}
}
func (f *workspacesFactories) GetWorkspace(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.GetWorkspace(ctx, instance)
	}
}
func (f *workspacesFactories) CreateWorkspace(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.CreateWorkspace(ctx, instance)
	}
}
func (f *workspacesFactories) UpdateWorkspace(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.UpdateWorkspace(ctx, instance)
	}
}
func (f *workspacesFactories) DeleteWorkspace(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.DeleteWorkspace(ctx, instance)
	}
}
func newWorkspacesFactories(transports WorkspacesClientTransports) WorkspacesFactories {
	return &workspacesFactories{transports: transports}
}

// workspacesEndpointers implements WorkspacesEndpointers
type workspacesEndpointers struct {
	target    string
	builder   sdx.Builder
	factories WorkspacesFactories
	logger    log.Logger
	options   []sd.EndpointerOption
}

func (e *workspacesEndpointers) ListWorkspaces(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.ListWorkspaces(ctx), e.logger, e.options...)
}
func (e *workspacesEndpointers) GetWorkspace(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.GetWorkspace(ctx), e.logger, e.options...)
}
func (e *workspacesEndpointers) CreateWorkspace(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.CreateWorkspace(ctx), e.logger, e.options...)
}
func (e *workspacesEndpointers) UpdateWorkspace(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.UpdateWorkspace(ctx), e.logger, e.options...)
}
func (e *workspacesEndpointers) DeleteWorkspace(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.DeleteWorkspace(ctx), e.logger, e.options...)
}
func newWorkspacesEndpointers(
	target string,
	builder sdx.Builder,
	factories WorkspacesFactories,
	logger log.Logger,
	options ...sd.EndpointerOption,
) WorkspacesEndpointers {
	return &workspacesEndpointers{
		target:    target,
		builder:   builder,
		factories: factories,
		logger:    logger,
		options:   options,
	}
}

// workspacesBalancers implements WorkspacesBalancers
type workspacesBalancers struct {
	factory         lbx.BalancerFactory
	endpointer      WorkspacesEndpointers
	listWorkspaces  lazyloadx.Group[lb.Balancer]
	getWorkspace    lazyloadx.Group[lb.Balancer]
	createWorkspace lazyloadx.Group[lb.Balancer]
	updateWorkspace lazyloadx.Group[lb.Balancer]
	deleteWorkspace lazyloadx.Group[lb.Balancer]
}

func (b *workspacesBalancers) ListWorkspaces(ctx context.Context) (lb.Balancer, error) {
	color, _ := stain.ExtractColor(ctx)
	balancer, err, _ := b.listWorkspaces.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.ListWorkspaces))
	return balancer, err
}
func (b *workspacesBalancers) GetWorkspace(ctx context.Context) (lb.Balancer, error) {
	color, _ := stain.ExtractColor(ctx)
	balancer, err, _ := b.getWorkspace.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.GetWorkspace))
	return balancer, err
}
func (b *workspacesBalancers) CreateWorkspace(ctx context.Context) (lb.Balancer, error) {
	color, _ := stain.ExtractColor(ctx)
	balancer, err, _ := b.createWorkspace.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.CreateWorkspace))
	return balancer, err
}
func (b *workspacesBalancers) UpdateWorkspace(ctx context.Context) (lb.Balancer, error) {
	color, _ := stain.ExtractColor(ctx)
	balancer, err, _ := b.updateWorkspace.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.UpdateWorkspace))
	return balancer, err
}
func (b *workspacesBalancers) DeleteWorkspace(ctx context.Context) (lb.Balancer, error) {
	color, _ := stain.ExtractColor(ctx)
	balancer, err, _ := b.deleteWorkspace.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.DeleteWorkspace))
	return balancer, err
}
func newWorkspacesBalancers(factory lbx.BalancerFactory, endpointer WorkspacesEndpointers) WorkspacesBalancers {
	return &workspacesBalancers{
		factory:         factory,
		endpointer:      endpointer,
		listWorkspaces:  lazyloadx.Group[lb.Balancer]{},
		getWorkspace:    lazyloadx.Group[lb.Balancer]{},
		createWorkspace: lazyloadx.Group[lb.Balancer]{},
		updateWorkspace: lazyloadx.Group[lb.Balancer]{},
		deleteWorkspace: lazyloadx.Group[lb.Balancer]{},
	}
}

// workspacesClientService implements WorkspacesClientService
type workspacesClientService struct {
	endpoints     WorkspacesClientEndpoints
	transportName string
}

func (c *workspacesClientService) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.ListWorkspaces(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*ListWorkspacesResponse), nil
}
func (c *workspacesClientService) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.GetWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*Workspace), nil
}
func (c *workspacesClientService) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.CreateWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*Workspace), nil
}
func (c *workspacesClientService) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.UpdateWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*Workspace), nil
}
func (c *workspacesClientService) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.DeleteWorkspace(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}
func newWorkspacesClientService(endpoints WorkspacesClientEndpoints, transportName string) WorkspacesService {
	return &workspacesClientService{endpoints: endpoints, transportName: transportName}
}
