package lbx

// 最快响应时间

//
//type _Balancer struct {
//	mutex    sync.RWMutex
//	conns    []*conn
//	lastSync time.Time
//	endpoint string
//}
//
//func (b *_Balancer) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
//	b.mutex.RLock()
//	if len(b.conns) == 0 {
//		b.mutex.RUnlock()
//		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
//	}
//	var res *conn
//	for _, c := range b.conns {
//		if res == nil || res.response > c.response {
//			res = c
//		}
//	}
//	b.mutex.RUnlock()
//
//	return balancer.PickResult{
//		SubConn: res.SubConn,
//		Done: func(info balancer.DoneInfo) {
//		},
//	}, nil
//}
//
//func (b *_Builder) Build(info base.PickerBuildInfo) balancer.Picker {
//	conns := make([]*conn, 0, len(info.ReadySCs))
//	for con, val := range info.ReadySCs {
//		conns = append(conns, &conn{
//			SubConn: con,
//			address: val.Address,
//			// 随便设置一个默认值。当然这个默认值会对初始的负载均衡有影响
//			// 不过一段时间之后就没什么影响了
//			response: time.Millisecond * 100,
//		})
//	}
//	res := &_Balancer{
//		conns: conns,
//	}
//
//	// 基本的思路是启动一个 goroutine 异步地去拉 prometheus 上的响应时间的数据，即调用 updateResp
//	// 但是有一个很大的问题：我们这里不好怎么退出，因为没有 gRPC 不会调用 Close 方法
//	// 可以考虑使用 runtime.SetFinalizer 来在 res 被回收的时候得到通知
//	panic("implement me")
//	return res
//}
//
//func (b *_Balancer) updateRespTime(endpoint, query string) {
//
//}
//
//type _Builder struct {
//	// prometheus 的地址
//	Endpoint string
//	Query    string
//	// 刷新响应时间的间隔
//	Interval time.Duration
//}
//
//type conn struct {
//	balancer.SubConn
//	address resolver.Address
//	// 响应时间
//	response time.Duration
//}
