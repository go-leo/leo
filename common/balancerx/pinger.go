package balancerx

type Pinger interface {
	IsAlive(server *Server)
}
