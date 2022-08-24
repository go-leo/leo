package client

import (
	"context"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	"github.com/go-leo/leo/global"
	"github.com/go-leo/leo/registry"
)

var _ resolver.Builder = new(ResolverBuilder)

type ResolverBuilder struct {
	discovery registry.Discovery
}

func NewResolverBuilder(discovery registry.Discovery) resolver.Builder {
	return &ResolverBuilder{discovery: discovery}
}

/**
target格式：
Scheme://Authority/Endpoint
consul://username:password@ip:port/service.name?scheme=http&datacenter=dev&token=12345&wait_time=1s&tls=true
nacos://ip1:port1,ip2:port2/service.name?contentType=yaml&namespace=ns&group=g&dataId=d
*/

// Build 创建Resolver并启动服务名解析
func (rb *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// 比较target的Scheme是否与discovery的Scheme一直
	if target.URL.Scheme != rb.discovery.Scheme() {
		return nil, fmt.Errorf("target schema is %s, but discovery schema is %s", target.URL.Scheme, rb.discovery.Scheme())
	}
	// 从url上获取serviceInfo信息
	uri := target.URL
	serviceInfo := registry.ServiceInfoFromURL(uri, registry.TransportGRPC)
	// 创建Resolver
	ctx, cancel := context.WithCancel(context.Background())
	r := &Resolver{
		ctx:         ctx,
		cancelFunc:  cancel,
		serviceInfo: serviceInfo,
		cc:          cc,
		opts:        opts,
		discovery:   rb.discovery,
	}

	// 开始解析
	if err := r.Start(); err != nil {
		return nil, err
	}
	return r, nil
}

func (rb *ResolverBuilder) Scheme() string {
	return rb.discovery.Scheme()
}

var _ resolver.Resolver = new(Resolver)

type Resolver struct {
	// ctx 和 cancelFunc，可以监听解析器Close信号，以便goroutine可以正常退出
	ctx        context.Context
	cancelFunc context.CancelFunc
	// cc 和 opts 原生gRPC框架传入的，可以控制客户端连接的状态和行为
	cc   resolver.ClientConn
	opts resolver.BuildOptions
	// discovery Leo服务发现组件
	discovery registry.Discovery
	// serviceInfo 当前需要被解析的服务
	serviceInfo *registry.ServiceInfo
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (r *Resolver) Close() {
	r.cancelFunc()
	err := r.discovery.StopWatch(context.Background(), r.serviceInfo)
	if err != nil {
		global.Logger().Errorf("failed to stop watch, %v", err)
	}
}

func (r *Resolver) Start() error {
	// 一开始先获取一次服务信息
	service, err := r.discovery.GetService(r.ctx, r.serviceInfo)
	if err != nil {
		return err
	}
	// 更新grpc的连接状态和ip地址
	if err := r.update(service); err != nil {
		return err
	}
	// 监听服务变化
	eventC, err := r.discovery.Watch(r.ctx, r.serviceInfo)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-r.ctx.Done():
				// 退出
				return
			case service := <-eventC:
				// 有事件产生，更新grpc的连接状态和ip地址
				if err := r.update(service); err != nil {
					global.Logger().Errorf("failed to resolve, %v", err)
				}
			}
		}
	}()
	return nil
}

func (r *Resolver) update(service []*registry.ServiceInfo) error {
	if len(service) <= 0 {
		return nil
	}
	// registry.ServiceInfo转成resolver.Address
	addresses := r.convertService(service)
	// 更新grpc连接状态
	return r.cc.UpdateState(resolver.State{Addresses: addresses})
}

func (r *Resolver) convertService(service []*registry.ServiceInfo) []resolver.Address {
	addresses := make([]resolver.Address, 0, len(service))
	for _, service := range service {
		attr := &attributes.Attributes{}
		address := resolver.Address{
			Addr:       net.JoinHostPort(service.Host, strconv.Itoa(service.Port)),
			ServerName: r.serviceInfo.Name,
			Attributes: attr,
		}
		addresses = append(addresses, address)
	}
	sort.Slice(addresses, func(i, j int) bool {
		return strings.Compare(addresses[i].Addr, addresses[j].Addr) > 0
	})
	return addresses
}
