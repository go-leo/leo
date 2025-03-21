// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package demo

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
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// DemoService is a service
type DemoService interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
	GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
	UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*emptypb.Empty, error)
	GetUserAvatar(ctx context.Context, request *GetUserAvatarRequest) (*httpbody.HttpBody, error)
}

// DemoServerEndpoints is server endpoints
type DemoServerEndpoints interface {
	CreateUser(ctx context.Context) endpoint.Endpoint
	DeleteUser(ctx context.Context) endpoint.Endpoint
	UpdateUser(ctx context.Context) endpoint.Endpoint
	GetUser(ctx context.Context) endpoint.Endpoint
	GetUsers(ctx context.Context) endpoint.Endpoint
	UploadUserAvatar(ctx context.Context) endpoint.Endpoint
	GetUserAvatar(ctx context.Context) endpoint.Endpoint
}

// DemoClientEndpoints is client endpoints
type DemoClientEndpoints interface {
	CreateUser(ctx context.Context) (endpoint.Endpoint, error)
	DeleteUser(ctx context.Context) (endpoint.Endpoint, error)
	UpdateUser(ctx context.Context) (endpoint.Endpoint, error)
	GetUser(ctx context.Context) (endpoint.Endpoint, error)
	GetUsers(ctx context.Context) (endpoint.Endpoint, error)
	UploadUserAvatar(ctx context.Context) (endpoint.Endpoint, error)
	GetUserAvatar(ctx context.Context) (endpoint.Endpoint, error)
}

// DemoClientTransports is client transports
type DemoClientTransports interface {
	CreateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	DeleteUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	UpdateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	GetUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	GetUsers(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	UploadUserAvatar(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	GetUserAvatar(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

// DemoFactories is client factories
type DemoFactories interface {
	CreateUser(ctx context.Context) sd.Factory
	DeleteUser(ctx context.Context) sd.Factory
	UpdateUser(ctx context.Context) sd.Factory
	GetUser(ctx context.Context) sd.Factory
	GetUsers(ctx context.Context) sd.Factory
	UploadUserAvatar(ctx context.Context) sd.Factory
	GetUserAvatar(ctx context.Context) sd.Factory
}

// DemoEndpointers is client endpointers
type DemoEndpointers interface {
	CreateUser(ctx context.Context, color string) (sd.Endpointer, error)
	DeleteUser(ctx context.Context, color string) (sd.Endpointer, error)
	UpdateUser(ctx context.Context, color string) (sd.Endpointer, error)
	GetUser(ctx context.Context, color string) (sd.Endpointer, error)
	GetUsers(ctx context.Context, color string) (sd.Endpointer, error)
	UploadUserAvatar(ctx context.Context, color string) (sd.Endpointer, error)
	GetUserAvatar(ctx context.Context, color string) (sd.Endpointer, error)
}

// DemoBalancers is client balancers
type DemoBalancers interface {
	CreateUser(ctx context.Context) (lb.Balancer, error)
	DeleteUser(ctx context.Context) (lb.Balancer, error)
	UpdateUser(ctx context.Context) (lb.Balancer, error)
	GetUser(ctx context.Context) (lb.Balancer, error)
	GetUsers(ctx context.Context) (lb.Balancer, error)
	UploadUserAvatar(ctx context.Context) (lb.Balancer, error)
	GetUserAvatar(ctx context.Context) (lb.Balancer, error)
}

// demoServerEndpoints implements DemoServerEndpoints
type demoServerEndpoints struct {
	svc         DemoService
	middlewares []endpoint.Middleware
}

func (e *demoServerEndpoints) CreateUser(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateUser(ctx, request.(*CreateUserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *demoServerEndpoints) DeleteUser(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteUser(ctx, request.(*DeleteUsersRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *demoServerEndpoints) UpdateUser(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateUser(ctx, request.(*UpdateUserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *demoServerEndpoints) GetUser(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetUser(ctx, request.(*GetUserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *demoServerEndpoints) GetUsers(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetUsers(ctx, request.(*GetUsersRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *demoServerEndpoints) UploadUserAvatar(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.UploadUserAvatar(ctx, request.(*UploadUserAvatarRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *demoServerEndpoints) GetUserAvatar(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetUserAvatar(ctx, request.(*GetUserAvatarRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

// demoFactories implements DemoFactories
type demoFactories struct {
	transports DemoClientTransports
}

func (f *demoFactories) CreateUser(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.CreateUser(ctx, instance)
	}
}

func (f *demoFactories) DeleteUser(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.DeleteUser(ctx, instance)
	}
}

func (f *demoFactories) UpdateUser(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.UpdateUser(ctx, instance)
	}
}

func (f *demoFactories) GetUser(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.GetUser(ctx, instance)
	}
}

func (f *demoFactories) GetUsers(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.GetUsers(ctx, instance)
	}
}

func (f *demoFactories) UploadUserAvatar(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.UploadUserAvatar(ctx, instance)
	}
}

func (f *demoFactories) GetUserAvatar(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.GetUserAvatar(ctx, instance)
	}
}

// demoEndpointers implements DemoEndpointers
type demoEndpointers struct {
	target    string
	builder   sdx.Builder
	factories DemoFactories
	logger    log.Logger
	options   []sd.EndpointerOption
}

func (e *demoEndpointers) CreateUser(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.CreateUser(ctx), e.logger, e.options...)
}

func (e *demoEndpointers) DeleteUser(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.DeleteUser(ctx), e.logger, e.options...)
}

func (e *demoEndpointers) UpdateUser(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.UpdateUser(ctx), e.logger, e.options...)
}

func (e *demoEndpointers) GetUser(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.GetUser(ctx), e.logger, e.options...)
}

func (e *demoEndpointers) GetUsers(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.GetUsers(ctx), e.logger, e.options...)
}

func (e *demoEndpointers) UploadUserAvatar(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.UploadUserAvatar(ctx), e.logger, e.options...)
}

func (e *demoEndpointers) GetUserAvatar(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.GetUserAvatar(ctx), e.logger, e.options...)
}

// demoBalancers implements DemoBalancers
type demoBalancers struct {
	factory          lbx.BalancerFactory
	endpointer       DemoEndpointers
	createUser       lazyloadx.Group[lb.Balancer]
	deleteUser       lazyloadx.Group[lb.Balancer]
	updateUser       lazyloadx.Group[lb.Balancer]
	getUser          lazyloadx.Group[lb.Balancer]
	getUsers         lazyloadx.Group[lb.Balancer]
	uploadUserAvatar lazyloadx.Group[lb.Balancer]
	getUserAvatar    lazyloadx.Group[lb.Balancer]
}

func (b *demoBalancers) CreateUser(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.createUser.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.CreateUser))
	return balancer, err
}
func (b *demoBalancers) DeleteUser(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.deleteUser.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.DeleteUser))
	return balancer, err
}
func (b *demoBalancers) UpdateUser(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.updateUser.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.UpdateUser))
	return balancer, err
}
func (b *demoBalancers) GetUser(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.getUser.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.GetUser))
	return balancer, err
}
func (b *demoBalancers) GetUsers(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.getUsers.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.GetUsers))
	return balancer, err
}
func (b *demoBalancers) UploadUserAvatar(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.uploadUserAvatar.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.UploadUserAvatar))
	return balancer, err
}
func (b *demoBalancers) GetUserAvatar(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
	balancer, err, _ := b.getUserAvatar.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.GetUserAvatar))
	return balancer, err
}
func newDemoBalancers(factory lbx.BalancerFactory, endpointer DemoEndpointers) DemoBalancers {
	return &demoBalancers{
		factory:          factory,
		endpointer:       endpointer,
		createUser:       lazyloadx.Group[lb.Balancer]{},
		deleteUser:       lazyloadx.Group[lb.Balancer]{},
		updateUser:       lazyloadx.Group[lb.Balancer]{},
		getUser:          lazyloadx.Group[lb.Balancer]{},
		getUsers:         lazyloadx.Group[lb.Balancer]{},
		uploadUserAvatar: lazyloadx.Group[lb.Balancer]{},
		getUserAvatar:    lazyloadx.Group[lb.Balancer]{},
	}
}

// demoClientEndpoints implements DemoClientEndpoints
type demoClientEndpoints struct {
	balancers DemoBalancers
}

func (e *demoClientEndpoints) CreateUser(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *demoClientEndpoints) DeleteUser(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.DeleteUser(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *demoClientEndpoints) UpdateUser(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.UpdateUser(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *demoClientEndpoints) GetUser(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *demoClientEndpoints) GetUsers(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *demoClientEndpoints) UploadUserAvatar(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.UploadUserAvatar(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

func (e *demoClientEndpoints) GetUserAvatar(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.GetUserAvatar(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}

// demoClientService implements DemoClientService
type demoClientService struct {
	endpoints     DemoClientEndpoints
	transportName string
}

func (c *demoClientService) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/CreateUser")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*CreateUserResponse), nil
}

func (c *demoClientService) DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/DeleteUser")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.DeleteUser(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *demoClientService) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/UpdateUser")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.UpdateUser(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *demoClientService) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/GetUser")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*GetUserResponse), nil
}

func (c *demoClientService) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/GetUsers")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*GetUsersResponse), nil
}

func (c *demoClientService) UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*emptypb.Empty, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/UploadUserAvatar")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.UploadUserAvatar(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *demoClientService) GetUserAvatar(ctx context.Context, request *GetUserAvatarRequest) (*httpbody.HttpBody, error) {
	ctx = endpointx.NameInjector(ctx, "/leo.example.demo.v1.Demo/GetUserAvatar")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.GetUserAvatar(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*httpbody.HttpBody), nil
}
