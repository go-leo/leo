package chainx

// Middleware interface is anything which implements a MiddlewareFunc named Middleware.
type Middleware interface {
	Middleware(cmd Command) Command
}

// MiddlewareFunc is a function which receives a Command and returns another Command.
type MiddlewareFunc func(cmd Command) Command

// Middleware allows MiddlewareFunc to implement the Middleware interface.
func (mw MiddlewareFunc) Middleware(cmd Command) Command {
	return mw(cmd)
}
