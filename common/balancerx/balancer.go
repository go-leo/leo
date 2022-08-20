package balancerx

type ServerGroup int

const (
	All ServerGroup = iota
	StatusUp
	StatusNotUp
)

type Balancer interface {
	// AddServers adds servers to balancer
	AddServers(newServers []*Server)
	// ChooseServer chooses a server from balancer.
	// key An object that the balancer may use to determine which server to return. nil if
	// the balancer does not use this parameter.
	ChooseServer(key any) *Server
	// MarkServerDown to be called by the clients of the load balancer to notify that a Server is down
	// else, the balancer will think its still Alive until the next Ping cycle - potentially
	// (assuming that the balancer Impl does a ping)
	MarkServerDown(server *Server)
	// GetReachableServers returns Only the servers that are up and reachable.
	GetReachableServers() []*Server
	// GetAllServers returns all known servers, both reachable and unreachable.
	GetAllServers() []*Server
	// GetServerList returns list of servers that this balancer knows about
	GetServerList(serverGroup ServerGroup) []*Server
	// GetBalancerStats obtain LoadBalancer related Statistics
	GetBalancerStats() *BalancerStats
}
