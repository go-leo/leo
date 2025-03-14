package internalsd

import (
	"context"
	"github.com/go-kit/kit/sd"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/gox/netx/addrx"
	"github.com/go-leo/leo/v3/sdx"
	"net"
	"net/url"
	"runtime"
)

func NewRegistrar(ctx context.Context, lis net.Listener, builder sdx.Builder, instance string, color string, logger kitlog.Logger) (sd.Registrar, error) {
	if builder == nil {
		return nil, nil
	}
	instanceUrl, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	ip, port, err := addrx.GlobalUnicastAddr(lis.Addr())
	if err != nil {
		return nil, err
	}
	return builder.BuildRegistrar(ctx, instanceUrl, ip, port, color, logger)
}

func Register(registrar sd.Registrar) {
	if registrar == nil {
		return
	}
	go func() {
		// make service register after server serve
		runtime.Gosched()
		registrar.Register()
	}()
	// make service register after server serve
	runtime.Gosched()
}

func Deregister(registrar sd.Registrar) {
	if registrar == nil {
		return
	}
	registrar.Deregister()
}
