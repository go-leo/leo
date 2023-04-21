package hystrixbreaker

import (
	"context"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	metricCollector "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/afex/hystrix-go/plugins"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(commandName string, opts ...Option) grpc.UnaryClientInterceptor {
	o := defaultOptions()
	o.apply(opts...)
	if o.statsD != nil {
		c, err := plugins.InitializeStatsdCollector(o.statsD)
		if err != nil {
			panic(err)
		}
		metricCollector.Registry.Register(c.NewStatsdCollector)
	}
	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
		Timeout:                durationToInt(o.timeout, time.Millisecond),
		MaxConcurrentRequests:  o.maxConcurrentRequests,
		RequestVolumeThreshold: o.requestVolumeThreshold,
		SleepWindow:            durationToInt(o.sleepWindow, time.Millisecond),
		ErrorPercentThreshold:  o.errorPercentThreshold,
	})
	return unaryClientInterceptor(commandName, o)
}

func unaryClientInterceptor(commandName string, o *options) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) (err error) {
		err = hystrix.DoC(ctx, commandName, func(ctx context.Context) error {
			err = invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return err
			}
			return nil
		}, o.fallbackFunc)
		return err
	}
}

func durationToInt(duration, unit time.Duration) int {
	durationAsNumber := duration / unit
	if int64(durationAsNumber) > int64(maxInt) {
		return maxInt
	}
	return int(durationAsNumber)
}
