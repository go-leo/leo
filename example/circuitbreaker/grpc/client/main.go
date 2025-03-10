package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-leo/leo/v3/circuitbreakerx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	//GoBreaker()
	GoogleSRE()
	//Hystrix()
}

func GoBreaker() {
	// sony gobreaker
	mdw := circuitbreakerx.GoBreaker(func(endpointName string) (*gobreaker.CircuitBreaker, error) {
		return gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name: endpointName,
			//	MaxRequests 的定义:
			//		* MaxRequests 是断路器（CircuitBreaker）在半开启状态（half-open state）下允许通过的最大请求数量。
			//	行为逻辑:
			//		* 如果 MaxRequests 设置为 0：
			//			* 断路器在半开启状态下仅允许 1 个请求通过。
			//		* 如果 MaxRequests 设置为大于 0 的值：
			//			* 断路器在半开启状态下允许指定数量的请求通过。
			//	为什么要限制请求数量:
			//		* 在半开启状态下，断路器允许少量请求通过以测试下游服务是否恢复正常。
			//		* 限制请求数量可以防止过多的请求同时通过，从而避免对下游服务造成过大的压力或再次触发故障。
			MaxRequests: 10,
			//	Interval 的定义:
			//		* Interval 是断路器（CircuitBreaker）在关闭状态（closed state）下的一个周期性时间间隔。
			//		* 在这个时间间隔内，断路器会定期清除内部的计数器（Counts）。
			//	行为逻辑:
			//		* 如果 Interval 设置为小于或等于 0：
			//			* 断路器在关闭状态下不会清除内部的计数器。
			//		* 如果 Interval 设置为大于 0：
			//			* 断路器会在每个 Interval 周期内清除内部的计数器。
			//	为什么要清除计数器:
			//		* 计数器通常用于记录请求的成功、失败或其他状态信息。
			//		* 定期清除计数器可以防止历史数据对当前状态评估的影响，确保断路器的决策基于最近的请求状态。
			Interval: 10 * time.Second,
			//	Timeout 的定义:
			//		* Timeout 是断路器（CircuitBreaker）在开启状态（open state）下的一个时间间隔。
			//		* 在这个时间间隔后，断路器的状态会从开启状态（open）转变为半开启状态（half-open）。
			//	行为逻辑:
			//		* 如果 Timeout 设置为小于或等于 0：
			//			* 断路器的超时值会被默认设置为 60 秒。
			//		* 如果 Timeout 设置为大于 0：
			//			* 断路器会在指定的时间间隔后进入半开启状态。
			//	为什么要从开启状态变为半开启状态:
			//		* 当断路器处于开启状态时，它会阻止所有请求以保护系统免受过载或故障的影响。
			//		* 转变为半开启状态后，断路器允许少量请求通过，以测试下游服务是否恢复正常。
			//		* 如果请求成功，断路器会切换到关闭状态；如果请求失败，则重新进入开启状态。
			Timeout: 5 * time.Second,
			// 	ReadyToTrip 的定义:
			//		* ReadyToTrip 是一个回调函数，当断路器处于关闭状态（closed state）且请求失败时，会调用该函数。
			//		* 该函数接收一个 Counts 的副本作为参数，Counts 通常包含请求的成功、失败等统计信息。
			//	行为逻辑:
			//		* 如果 ReadyToTrip 返回 true：
			//			* 断路器将被切换到开启状态（open state）。
			//		* 如果 ReadyToTrip 返回 false：
			//			* 断路器保持在关闭状态。
			// 	为什么要使用 ReadyToTrip:
			//		* ReadyToTrip 提供了一种灵活的方式来定义断路器何时应该从关闭状态切换到开启状态。
			//		* 通过自定义 ReadyToTrip 函数，可以根据具体的业务需求来决定断路器的行为。
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return failureRatio > 0.1
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("name:%s, from:%s, to:%s\n", name, from.String(), to.String())
			},
		}), nil
	})
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	for {
		reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
		if err != nil {
			//log.Printf("could not greet: %v\n", err)
		}
		_ = reply
		//log.Printf("greet: %v\n", reply)
		time.Sleep(100 * time.Millisecond)
	}
}

func GoogleSRE() {
	mdw := circuitbreakerx.GoogleSreBreaker(func(endpointName string) (circuitbreaker.CircuitBreaker, error) {
		return sre.NewBreaker(
			// 	WithSuccess 的定义:
			//		* WithSuccess 设置断路器的成功阈值。
			//		* 成功阈值 K 通过 1 / Success 计算得出，其中 Success 是一个介于 0 和 1 之间的浮点数
			// 	为什么要使用 WithSuccess:
			//		* 通过设置不同的成功阈值，可以控制断路器对请求成功率的敏感度。
			//		* 较低的成功阈值（如 0.2）会使断路器更容易进入开启状态，因为允许的失败比例更高。
			//		* 较高的成功阈值（如 0.8）会使断路器更严格，只有在请求成功率非常低的情况下才会进入开启状态。
			sre.WithSuccess(0.1),
			// 	WithRequest 的定义:
			//		* WithRequest 设置断路器的最小请求数量。
			//	如果设置为 100：
			//		* 断路器在关闭状态下至少需要处理 100 个请求，才会根据成功率来决定是否切换到开启状态。
			//	如果设置为 0：
			//		* 断路器在处理任何数量的请求后都会根据成功率来决定是否切换到开启状态。
			// 	为什么要使用 WithRequest:
			//		* 通过设置最小请求数量，可以确保断路器在有足够的请求样本后才进行成功率计算，从而避免因样本不足而导致的误判。
			//		* 较高的最小请求数量可以提高断路器决策的准确性，特别是在请求量较小的情况下。
			sre.WithRequest(100),
			// 	WithWindow 的定义:
			//		* WithWindow设置断路器的统计窗口持续时间。
			//	如果 d 设置为 3 * time.Second：
			//		* 断路器会在 3 秒的窗口内收集请求的成功和失败数据。
			//	如果 d 设置为 5 * time.Second：
			//		* 断路器会在 5 秒的窗口内收集请求的成功和失败数据。
			// 	为什么要使用 WithWindow:
			//		* 通过设置统计窗口的持续时间，可以控制断路器在多长时间内收集请求数据，从而影响其决策的粒度和响应速度。
			//		* 较短的窗口时间可以更快地响应系统状态的变化，但可能会导致统计样本较少。
			//		* 较长的窗口时间可以收集更多的样本数据，提高统计的准确性，但可能会延迟对系统状态变化的响应。
			sre.WithWindow(3*time.Second),
			// 	WithBucket 的定义:
			//		* WithBucket设置断路器的统计窗口内的时间桶数量。
			//	如果 b 设置为 10：
			//		* 断路器会在统计窗口内划分 10 个时间桶。
			//	如果 b 设置为 20：
			//		* 断路器会在统计窗口内划分 20 个时间桶。
			//	为什么要使用 WithBucket:
			//		* 通过设置时间桶的数量，可以控制断路器在统计窗口内的时间粒度。
			//		* 较多的时间桶可以提供更细粒度的统计信息，有助于更准确地评估请求的成功率和失败率。
			//		* 较少的时间桶可以减少内存占用，但可能会降低统计的准确性。
			sre.WithBucket(100),
		), nil
	})
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	for {
		time.Sleep(100 * time.Millisecond)
		reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
		if err != nil {
			fmt.Println("----------fail----------")
		} else {
			fmt.Println("OK")
		}
		_ = reply
	}
}

func Hystrix() {
	hystrix.Configure(map[string]hystrix.CommandConfig{
		"/helloworld.Greeter/SayHello": {
			// 	Timeout定义：
			//		命令执行的最大超时时间（以毫秒为单位）。
			//	行为逻辑：
			//		如果命令在指定时间内未完成，则会被中断并视为失败。
			//	示例：
			//		Timeout: 1000 表示命令的超时时间为 1 秒。
			Timeout: 1000,
			//	MaxConcurrentRequests定义：
			//		允许同时执行的最大请求数量。
			//	行为逻辑：
			//		当并发请求数量达到此限制时，后续请求将被拒绝。
			//	示例：
			//		MaxConcurrentRequests: 100 表示最多允许 100 个并发请求。
			MaxConcurrentRequests: 100,
			// 	RequestVolumeThreshold定义
			//		触发断路器打开所需的最小请求数量。
			//	行为逻辑
			//		只有当在滚动窗口内接收到的请求数量超过此阈值时，才会评估错误率。
			//	示例：
			//		RequestVolumeThreshold: 10 表示至少需要 10 个请求来评估错误率。
			RequestVolumeThreshold: 10,
			// 	SleepWindow定义
			//		断路器从打开状态切换到半开状态的时间间隔（以毫秒为单位）。
			//	行为逻辑：
			//		在此期间，所有请求都会立即失败；之后，断路器会尝试少量请求以检测服务是否恢复正常。
			//	示例：
			//		SleepWindow: 1000 表示断路器在打开状态下等待 1 秒后进入半开状态。
			SleepWindow: 1000,
			// 	ErrorPercentThreshold定义：
			//		触发断路器打开的错误百分比阈值。
			//	行为逻辑：
			//		如果在滚动窗口内的错误率超过此阈值，断路器将打开。
			//	示例：
			//		ErrorPercentThreshold: 10 表示如果错误率超过 10%，断路器将打开。
			ErrorPercentThreshold: 10,
		},
	})
	mdw := circuitbreakerx.Hystrix()
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	for {
		time.Sleep(100 * time.Millisecond)
		reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
		if err != nil {
			fmt.Println("----------fail----------")
		} else {
			fmt.Println("OK")
		}
		_ = reply
	}
}
